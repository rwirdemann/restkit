// Code generated by mockery v2.37.0. DO NOT EDIT.

package ports

import mock "github.com/stretchr/testify/mock"

// MockTemplate is an autogenerated mock type for the Template type
type MockTemplate struct {
	mock.Mock
}

type MockTemplate_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTemplate) EXPECT() *MockTemplate_Expecter {
	return &MockTemplate_Expecter{mock: &_m.Mock}
}

// Contains provides a mock function with given fields: filename, fragment
func (_m *MockTemplate) Contains(filename string, fragment string) (bool, error) {
	ret := _m.Called(filename, fragment)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(filename, fragment)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(filename, fragment)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(filename, fragment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTemplate_Contains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Contains'
type MockTemplate_Contains_Call struct {
	*mock.Call
}

// Contains is a helper method to define mock.On call
//   - filename string
//   - fragment string
func (_e *MockTemplate_Expecter) Contains(filename interface{}, fragment interface{}) *MockTemplate_Contains_Call {
	return &MockTemplate_Contains_Call{Call: _e.mock.On("Contains", filename, fragment)}
}

func (_c *MockTemplate_Contains_Call) Run(run func(filename string, fragment string)) *MockTemplate_Contains_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockTemplate_Contains_Call) Return(_a0 bool, _a1 error) *MockTemplate_Contains_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTemplate_Contains_Call) RunAndReturn(run func(string, string) (bool, error)) *MockTemplate_Contains_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: templ, out, path, data
func (_m *MockTemplate) Create(templ string, out string, path string, data interface{}) error {
	ret := _m.Called(templ, out, path, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, interface{}) error); ok {
		r0 = rf(templ, out, path, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTemplate_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockTemplate_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - templ string
//   - out string
//   - path string
//   - data interface{}
func (_e *MockTemplate_Expecter) Create(templ interface{}, out interface{}, path interface{}, data interface{}) *MockTemplate_Create_Call {
	return &MockTemplate_Create_Call{Call: _e.mock.On("Create", templ, out, path, data)}
}

func (_c *MockTemplate_Create_Call) Run(run func(templ string, out string, path string, data interface{})) *MockTemplate_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(interface{}))
	})
	return _c
}

func (_c *MockTemplate_Create_Call) Return(_a0 error) *MockTemplate_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTemplate_Create_Call) RunAndReturn(run func(string, string, string, interface{}) error) *MockTemplate_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with given fields: filename, before, fragment
func (_m *MockTemplate) Insert(filename string, before string, fragment string) error {
	ret := _m.Called(filename, before, fragment)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(filename, before, fragment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTemplate_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type MockTemplate_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - filename string
//   - before string
//   - fragment string
func (_e *MockTemplate_Expecter) Insert(filename interface{}, before interface{}, fragment interface{}) *MockTemplate_Insert_Call {
	return &MockTemplate_Insert_Call{Call: _e.mock.On("Insert", filename, before, fragment)}
}

func (_c *MockTemplate_Insert_Call) Run(run func(filename string, before string, fragment string)) *MockTemplate_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockTemplate_Insert_Call) Return(_a0 error) *MockTemplate_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTemplate_Insert_Call) RunAndReturn(run func(string, string, string) error) *MockTemplate_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTemplate creates a new instance of MockTemplate. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTemplate(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTemplate {
	mock := &MockTemplate{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
