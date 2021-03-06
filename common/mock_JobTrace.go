// Code generated by mockery v1.0.0

// This comment works around https://github.com/vektra/mockery/issues/155

package common

import context "context"
import mock "github.com/stretchr/testify/mock"

// MockJobTrace is an autogenerated mock type for the JobTrace type
type MockJobTrace struct {
	mock.Mock
}

// Fail provides a mock function with given fields: err
func (_m *MockJobTrace) Fail(err error, failureReason JobFailureReason) {
	_m.Called(err, failureReason)
}

// IsStdout provides a mock function with given fields:
func (_m *MockJobTrace) IsStdout() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetCancelFunc provides a mock function with given fields: cancelFunc
func (_m *MockJobTrace) SetCancelFunc(cancelFunc context.CancelFunc) {
	_m.Called(cancelFunc)
}

// Success provides a mock function with given fields:
func (_m *MockJobTrace) Success() {
	_m.Called()
}

// Write provides a mock function with given fields: p
func (_m *MockJobTrace) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
