package chunks

import (
	"testing"

	"github.com/dapperlabs/cadence/runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/dapperlabs/flow-go/engine/execution/computation/virtualmachine"
	"github.com/dapperlabs/flow-go/engine/verification"
	chModels "github.com/dapperlabs/flow-go/model/chunks"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/storage/ledger"
	"github.com/dapperlabs/flow-go/utils/unittest"
)

type ChunkVerifierTestSuite struct {
	suite.Suite
	verifier *ChunkVerifier
}

// Make sure variables are set properly
// SetupTest is executed prior to each individual test in this test suite
func (s *ChunkVerifierTestSuite) SetupTest() {
	s.verifier = NewChunkVerifier(newVirtualMachineMock())
}

// TestChunkVerifier invokes all the tests in this test suite
func TestChunkVerifier(t *testing.T) {
	suite.Run(t, new(ChunkVerifierTestSuite))
}

// TestHappyPath tests verification of the baseline verifiable chunk
func (s *ChunkVerifierTestSuite) TestHappyPath() {
	vch := GetBaselineVerifiableChunk(s.T(), []byte{})
	assert.NotNil(s.T(), vch)
	chFaults, err := s.verifier.Verify(vch)
	assert.Nil(s.T(), err)
	assert.Nil(s.T(), chFaults)
}

// TestMissingRegisterTouchForUpdate tests verification given a chunkdatapack missing a register touch (update)
func (s *ChunkVerifierTestSuite) TestMissingRegisterTouchForUpdate() {
	vch := GetBaselineVerifiableChunk(s.T(), []byte(""))
	assert.NotNil(s.T(), vch)
	// remove the second register touch
	vch.ChunkDataPack.RegisterTouches = vch.ChunkDataPack.RegisterTouches[:1]
	chFaults, err := s.verifier.Verify(vch)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), chFaults)
	_, ok := chFaults.(*chModels.CFMissingRegisterTouch)
	assert.True(s.T(), ok)
}

// TestMissingRegisterTouchForRead tests verification given a chunkdatapack missing a register touch (read)
func (s *ChunkVerifierTestSuite) TestMissingRegisterTouchForRead() {
	vch := GetBaselineVerifiableChunk(s.T(), []byte(""))
	assert.NotNil(s.T(), vch)
	// remove the second register touch
	vch.ChunkDataPack.RegisterTouches = vch.ChunkDataPack.RegisterTouches[1:]
	chFaults, err := s.verifier.Verify(vch)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), chFaults)
	_, ok := chFaults.(*chModels.CFMissingRegisterTouch)
	assert.True(s.T(), ok)
}

func (s *ChunkVerifierTestSuite) TestWrongEndState() {
	vch := GetBaselineVerifiableChunk(s.T(), []byte("wrongEndState"))
	assert.NotNil(s.T(), vch)
	chFaults, err := s.verifier.Verify(vch)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), chFaults)
	_, ok := chFaults.(*chModels.CFNonMatchingFinalState)
	assert.True(s.T(), ok)
}

func GetBaselineVerifiableChunk(t *testing.T, script []byte) *verification.VerifiableChunk {
	// Collection setup

	coll := unittest.CollectionFixture(5)
	coll.Transactions[3] = &flow.TransactionBody{Script: script}

	guarantee := coll.Guarantee()

	// Block setup
	payload := flow.Payload{
		Identities: unittest.IdentityListFixture(32),
		Guarantees: []*flow.CollectionGuarantee{&guarantee},
	}
	header := unittest.BlockHeaderFixture()
	header.PayloadHash = payload.Hash()
	block := flow.Block{
		Header:  header,
		Payload: payload,
	}

	// registerTouch and State setup
	id1 := make([]byte, 32)
	value1 := []byte{'a'}

	id2 := make([]byte, 32)
	id2[0] = byte(5)
	value2 := []byte{'b'}
	UpdatedValue2 := []byte{'B'}

	ids := make([][]byte, 0)
	values := make([][]byte, 0)
	ids = append(ids, id1, id2)
	values = append(values, value1, value2)

	var verifiableChunk verification.VerifiableChunk

	unittest.RunWithTempDBDir(t, func(dbDir string) {
		f, _ := ledger.NewTrieStorage(dbDir)
		startState, _ := f.UpdateRegisters(ids, values, f.EmptyStateCommitment())
		regTs, _ := f.GetRegisterTouches(ids, startState)

		ids = [][]byte{id2}
		values = [][]byte{UpdatedValue2}
		endState, _ := f.UpdateRegisters(ids, values, startState)

		// Chunk setup
		chunk := flow.Chunk{
			ChunkBody: flow.ChunkBody{
				CollectionIndex: 0,
				StartState:      startState,
			},
			Index: 0,
		}

		chunkDataPack := flow.ChunkDataPack{
			ChunkID:         chunk.ID(),
			StartState:      startState,
			RegisterTouches: regTs,
		}

		// ExecutionResult setup
		result := flow.ExecutionResult{
			ExecutionResultBody: flow.ExecutionResultBody{
				BlockID: block.ID(),
				Chunks:  flow.ChunkList{&chunk},
			},
		}

		receipt := flow.ExecutionReceipt{
			ExecutionResult: result,
		}

		verifiableChunk = verification.VerifiableChunk{
			ChunkIndex:    chunk.Index,
			EndState:      endState,
			Block:         &block,
			Receipt:       &receipt,
			Collection:    &coll,
			ChunkDataPack: &chunkDataPack,
		}
	})

	return &verifiableChunk

}

type blockContextMock struct {
	vm     *virtualMachineMock
	header *flow.Header
}

func (bc *blockContextMock) ExecuteTransaction(
	ledger virtualmachine.Ledger,
	tx *flow.TransactionBody,
	options ...virtualmachine.TransactionContextOption,
) (*virtualmachine.TransactionResult, error) {
	var txRes virtualmachine.TransactionResult
	switch string(tx.Script) {
	case "wrongEndState":
		id1 := make([]byte, 32)
		UpdatedValue1 := []byte{'F'}
		// add updates to the ledger
		ledger.Set(id1, UpdatedValue1)
		txRes = virtualmachine.TransactionResult{
			TransactionID: unittest.IdentifierFixture(),
			Events:        []runtime.Event{},
			Logs:          []string{"log1", "log2"}, // []string
			Error:         nil,                      // inside the runtime (e.g. div by zero, access account)
			GasUsed:       0,
		}
	default:
		id1 := make([]byte, 32)
		id2 := make([]byte, 32)
		id2[0] = byte(5)
		UpdatedValue2 := []byte{'B'}
		_, _ = ledger.Get(id1)
		ledger.Set(id2, UpdatedValue2)
		txRes = virtualmachine.TransactionResult{
			TransactionID: unittest.IdentifierFixture(),
			Events:        []runtime.Event{},
			Logs:          []string{"log1", "log2"}, // []string
			Error:         nil,                      // inside the runtime (e.g. div by zero, access account)
			GasUsed:       0,
		}
	}
	return &txRes, nil
}

func (bc *blockContextMock) ExecuteScript(
	ledger virtualmachine.Ledger,
	script []byte,
) (*virtualmachine.ScriptResult, error) {
	return nil, nil
}

// virtualMachineMock is a mocked virtualMachine
type virtualMachineMock struct {
}

func newVirtualMachineMock() *virtualMachineMock {
	// TODO set execution outcome
	return &virtualMachineMock{}
}

func (vm *virtualMachineMock) NewBlockContext(header *flow.Header) virtualmachine.BlockContext {
	return &blockContextMock{
		vm:     vm,
		header: header,
	}
}