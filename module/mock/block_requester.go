// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// BlockRequester is an autogenerated mock type for the BlockRequester type
type BlockRequester struct {
	mock.Mock
}

// Prune provides a mock function with given fields: final
func (_m *BlockRequester) Prune(final *flow.Header) {
	_m.Called(final)
}

// RequestBlock provides a mock function with given fields: blockID, height
func (_m *BlockRequester) RequestBlock(blockID flow.Identifier, height uint64) {
	_m.Called(blockID, height)
}

// RequestHeight provides a mock function with given fields: height
func (_m *BlockRequester) RequestHeight(height uint64) {
	_m.Called(height)
}

type mockConstructorTestingTNewBlockRequester interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlockRequester creates a new instance of BlockRequester. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlockRequester(t mockConstructorTestingTNewBlockRequester) *BlockRequester {
	mock := &BlockRequester{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
