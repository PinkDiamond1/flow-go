// Code generated by mockery v2.13.0. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// SealingConfigsGetter is an autogenerated mock type for the SealingConfigsGetter type
type SealingConfigsGetter struct {
	mock.Mock
}

// ApprovalRequestsThresholdConst provides a mock function with given fields:
func (_m *SealingConfigsGetter) ApprovalRequestsThresholdConst() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// ChunkAlphaConst provides a mock function with given fields:
func (_m *SealingConfigsGetter) ChunkAlphaConst() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// EmergencySealingActiveConst provides a mock function with given fields:
func (_m *SealingConfigsGetter) EmergencySealingActiveConst() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RequireApprovalsForSealConstructionDynamicValue provides a mock function with given fields:
func (_m *SealingConfigsGetter) RequireApprovalsForSealConstructionDynamicValue() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// RequireApprovalsForSealVerificationConst provides a mock function with given fields:
func (_m *SealingConfigsGetter) RequireApprovalsForSealVerificationConst() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

type NewSealingConfigsGetterT interface {
	mock.TestingT
	Cleanup(func())
}

// NewSealingConfigsGetter creates a new instance of SealingConfigsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSealingConfigsGetter(t NewSealingConfigsGetterT) *SealingConfigsGetter {
	mock := &SealingConfigsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}