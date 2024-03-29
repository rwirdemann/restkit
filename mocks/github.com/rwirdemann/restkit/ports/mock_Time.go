// Code generated by mockery v2.37.0. DO NOT EDIT.

package ports

import mock "github.com/stretchr/testify/mock"

// MockTime is an autogenerated mock type for the Time type
type MockTime struct {
	mock.Mock
}

type MockTime_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTime) EXPECT() *MockTime_Expecter {
	return &MockTime_Expecter{mock: &_m.Mock}
}

// TS provides a mock function with given fields:
func (_m *MockTime) TS() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTime_TS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TS'
type MockTime_TS_Call struct {
	*mock.Call
}

// TS is a helper method to define mock.On call
func (_e *MockTime_Expecter) TS() *MockTime_TS_Call {
	return &MockTime_TS_Call{Call: _e.mock.On("TS")}
}

func (_c *MockTime_TS_Call) Run(run func()) *MockTime_TS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTime_TS_Call) Return(_a0 string) *MockTime_TS_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTime_TS_Call) RunAndReturn(run func() string) *MockTime_TS_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTime creates a new instance of MockTime. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTime(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTime {
	mock := &MockTime{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
