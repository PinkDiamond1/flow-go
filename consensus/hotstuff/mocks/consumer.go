// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "github.com/dapperlabs/flow-go/consensus/hotstuff/model"
	mock "github.com/stretchr/testify/mock"
)

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

// OnBlockIncorporated provides a mock function with given fields: _a0
func (_m *Consumer) OnBlockIncorporated(_a0 *model.Block) {
	_m.Called(_a0)
}

// OnDoubleProposeDetected provides a mock function with given fields: _a0, _a1
func (_m *Consumer) OnDoubleProposeDetected(_a0 *model.Block, _a1 *model.Block) {
	_m.Called(_a0, _a1)
}

// OnDoubleVotingDetected provides a mock function with given fields: _a0, _a1
func (_m *Consumer) OnDoubleVotingDetected(_a0 *model.Vote, _a1 *model.Vote) {
	_m.Called(_a0, _a1)
}

// OnEnteringView provides a mock function with given fields: viewNumber
func (_m *Consumer) OnEnteringView(viewNumber uint64) {
	_m.Called(viewNumber)
}

// OnFinalizedBlock provides a mock function with given fields: _a0
func (_m *Consumer) OnFinalizedBlock(_a0 *model.Block) {
	_m.Called(_a0)
}

// OnForkChoiceGenerated provides a mock function with given fields: _a0, _a1
func (_m *Consumer) OnForkChoiceGenerated(_a0 uint64, _a1 *model.QuorumCertificate) {
	_m.Called(_a0, _a1)
}

// OnInvalidVoteDetected provides a mock function with given fields: _a0
func (_m *Consumer) OnInvalidVoteDetected(_a0 *model.Vote) {
	_m.Called(_a0)
}

// OnProposingBlock provides a mock function with given fields: proposal
func (_m *Consumer) OnProposingBlock(proposal *model.Proposal) {
	_m.Called(proposal)
}

// OnQcConstructedFromVotes provides a mock function with given fields: _a0
func (_m *Consumer) OnQcConstructedFromVotes(_a0 *model.QuorumCertificate) {
	_m.Called(_a0)
}

// OnQcIncorporated provides a mock function with given fields: _a0
func (_m *Consumer) OnQcIncorporated(_a0 *model.QuorumCertificate) {
	_m.Called(_a0)
}

// OnQcTriggeredViewChange provides a mock function with given fields: qc, newView
func (_m *Consumer) OnQcTriggeredViewChange(qc *model.QuorumCertificate, newView uint64) {
	_m.Called(qc, newView)
}

// OnReachedTimeout provides a mock function with given fields: timeout
func (_m *Consumer) OnReachedTimeout(timeout *model.TimerInfo) {
	_m.Called(timeout)
}

// OnReceiveProposal provides a mock function with given fields: currentView, proposal
func (_m *Consumer) OnReceiveProposal(currentView uint64, proposal *model.Proposal) {
	_m.Called(currentView, proposal)
}

// OnReceiveVote provides a mock function with given fields: currentView, vote
func (_m *Consumer) OnReceiveVote(currentView uint64, vote *model.Vote) {
	_m.Called(currentView, vote)
}

// OnStartingTimeout provides a mock function with given fields: _a0
func (_m *Consumer) OnStartingTimeout(_a0 *model.TimerInfo) {
	_m.Called(_a0)
}

// OnVoting provides a mock function with given fields: vote
func (_m *Consumer) OnVoting(vote *model.Vote) {
	_m.Called(vote)
}
