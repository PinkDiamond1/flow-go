// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import flow "github.com/dapperlabs/flow-go/model/flow"
import messages "github.com/dapperlabs/flow-go/model/messages"
import mock "github.com/stretchr/testify/mock"

// StateSynchronizer is an autogenerated mock type for the StateSynchronizer type
type StateSynchronizer struct {
	mock.Mock
}

// DeltaRange provides a mock function with given fields: startID, endID, onDelta
func (_m *StateSynchronizer) DeltaRange(startID flow.Identifier, endID flow.Identifier, onDelta func(*messages.ExecutionStateDelta) error) error {
	ret := _m.Called(startID, endID, onDelta)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.Identifier, func(*messages.ExecutionStateDelta) error) error); ok {
		r0 = rf(startID, endID, onDelta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}