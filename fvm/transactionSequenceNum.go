package fvm

import (
	"fmt"

	"github.com/onflow/flow-go/fvm/errors"
	"github.com/onflow/flow-go/fvm/programs"
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/module/trace"
)

type TransactionSequenceNumberChecker struct{}

func NewTransactionSequenceNumberChecker() *TransactionSequenceNumberChecker {
	return &TransactionSequenceNumberChecker{}
}

func (c *TransactionSequenceNumberChecker) Process(
	_ *VirtualMachine,
	ctx *Context,
	proc *TransactionProcedure,
	sth *state.StateHolder,
	_ *programs.Programs,
) error {
	return c.checkAndIncrementSequenceNumber(proc, ctx, sth)
}

func (c *TransactionSequenceNumberChecker) checkAndIncrementSequenceNumber(
	proc *TransactionProcedure,
	ctx *Context,
	sth *state.StateHolder,
) error {

	if ctx.Tracer != nil && proc.TraceSpan != nil {
		span := ctx.Tracer.StartSpanFromParent(proc.TraceSpan, trace.FVMSeqNumCheckTransaction)
		defer span.End()
	}

	nestedTxnId, err := sth.BeginNestedTransaction()
	if err != nil {
		return err
	}

	defer func() {
		mergeError := sth.Commit(nestedTxnId)
		if mergeError != nil {
			panic(mergeError)
		}
	}()

	accounts := state.NewAccounts(sth)
	proposalKey := proc.Transaction.ProposalKey

	accountKey, err := accounts.GetPublicKey(proposalKey.Address, proposalKey.KeyIndex)
	if err != nil {
		err = errors.NewInvalidProposalSignatureError(proposalKey.Address, proposalKey.KeyIndex, err)
		return fmt.Errorf("checking sequence number failed: %w", err)
	}

	if accountKey.Revoked {
		err = fmt.Errorf("proposal key has been revoked")
		err = errors.NewInvalidProposalSignatureError(proposalKey.Address, proposalKey.KeyIndex, err)
		return fmt.Errorf("checking sequence number failed: %w", err)
	}

	// Note that proposal key verification happens at the txVerifier and not here.

	valid := accountKey.SeqNumber == proposalKey.SequenceNumber

	if !valid {
		return errors.NewInvalidProposalSeqNumberError(proposalKey.Address, proposalKey.KeyIndex, accountKey.SeqNumber, proposalKey.SequenceNumber)
	}

	accountKey.SeqNumber++

	_, err = accounts.SetPublicKey(proposalKey.Address, proposalKey.KeyIndex, accountKey)
	if err != nil {
		// NOTE: we need to disable limits during restart or else restart may
		// fail on merging.
		sth.DisableAllLimitEnforcements()
		_ = sth.RestartNestedTransaction(nestedTxnId)
		sth.EnableAllLimitEnforcements()
		return fmt.Errorf("checking sequence number failed: %w", err)
	}

	return nil
}
