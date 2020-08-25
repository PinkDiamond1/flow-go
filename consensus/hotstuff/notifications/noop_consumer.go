package notifications

import (
	"github.com/dapperlabs/flow-go/consensus/hotstuff/model"
)

// NoopConsumer is an implementation of the notifications consumer that
// doesn't do anything.
type NoopConsumer struct{}

func NewNoopConsumer() *NoopConsumer {
	nc := &NoopConsumer{}
	return nc
}

func (*NoopConsumer) OnBlockIncorporated(*model.Block) {}

func (*NoopConsumer) OnFinalizedBlock(*model.Block) {}

func (*NoopConsumer) OnDoubleProposeDetected(*model.Block, *model.Block) {}

func (c *NoopConsumer) OnReceiveVote(uint64, *model.Vote) {}

func (c *NoopConsumer) OnReceiveProposal(uint64, *model.Proposal) {}

func (*NoopConsumer) OnEnteringView(uint64) {}

func (c *NoopConsumer) OnQcTriggeredViewChange(*model.QuorumCertificate, uint64) {}

func (c *NoopConsumer) OnProposingBlock(*model.Proposal) {}

func (c *NoopConsumer) OnVoting(*model.Vote) {}

func (c *NoopConsumer) OnQcConstructedFromVotes(*model.QuorumCertificate) {}

func (*NoopConsumer) OnStartingTimeout(*model.TimerInfo) {}

func (*NoopConsumer) OnReachedTimeout(*model.TimerInfo) {}

func (*NoopConsumer) OnQcIncorporated(*model.QuorumCertificate) {}

func (*NoopConsumer) OnForkChoiceGenerated(uint64, *model.QuorumCertificate) {}

func (*NoopConsumer) OnDoubleVotingDetected(*model.Vote, *model.Vote) {}

func (*NoopConsumer) OnInvalidVoteDetected(*model.Vote) {}
