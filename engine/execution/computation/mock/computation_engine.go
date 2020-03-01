// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import flow "github.com/dapperlabs/flow-go/model/flow"
import mock "github.com/stretchr/testify/mock"
import state "github.com/dapperlabs/flow-go/engine/execution/state"

// ComputationEngine is an autogenerated mock type for the ComputationEngine type
type ComputationEngine struct {
	mock.Mock
}

// ExecuteScript provides a mock function with given fields: _a0, _a1, _a2
func (_m *ComputationEngine) ExecuteScript(_a0 []byte, _a1 *flow.Header, _a2 *state.View) ([]byte, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte, *flow.Header, *state.View) []byte); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, *flow.Header, *state.View) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Process provides a mock function with given fields: originID, event
func (_m *ComputationEngine) Process(originID flow.Identifier, event interface{}) error {
	ret := _m.Called(originID, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, interface{}) error); ok {
		r0 = rf(originID, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessLocal provides a mock function with given fields: event
func (_m *ComputationEngine) ProcessLocal(event interface{}) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Submit provides a mock function with given fields: originID, event
func (_m *ComputationEngine) Submit(originID flow.Identifier, event interface{}) {
	_m.Called(originID, event)
}

// SubmitLocal provides a mock function with given fields: event
func (_m *ComputationEngine) SubmitLocal(event interface{}) {
	_m.Called(event)
}