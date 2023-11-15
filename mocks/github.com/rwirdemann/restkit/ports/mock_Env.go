// Code generated by mockery v2.37.0. DO NOT EDIT.

package ports

import mock "github.com/stretchr/testify/mock"

// MockEnv is an autogenerated mock type for the Env type
type MockEnv struct {
	mock.Mock
}

type MockEnv_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEnv) EXPECT() *MockEnv_Expecter {
	return &MockEnv_Expecter{mock: &_m.Mock}
}

// RKRoot provides a mock function with given fields:
func (_m *MockEnv) RKRoot() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEnv_RKRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RKRoot'
type MockEnv_RKRoot_Call struct {
	*mock.Call
}

// RKRoot is a helper method to define mock.On call
func (_e *MockEnv_Expecter) RKRoot() *MockEnv_RKRoot_Call {
	return &MockEnv_RKRoot_Call{Call: _e.mock.On("RKRoot")}
}

func (_c *MockEnv_RKRoot_Call) Run(run func()) *MockEnv_RKRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEnv_RKRoot_Call) Return(_a0 string, _a1 error) *MockEnv_RKRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEnv_RKRoot_Call) RunAndReturn(run func() (string, error)) *MockEnv_RKRoot_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEnv creates a new instance of MockEnv. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEnv(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEnv {
	mock := &MockEnv{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}