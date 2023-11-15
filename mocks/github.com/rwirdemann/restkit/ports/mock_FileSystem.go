// Code generated by mockery v2.37.0. DO NOT EDIT.

package ports

import (
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// MockFileSystem is an autogenerated mock type for the FileSystem type
type MockFileSystem struct {
	mock.Mock
}

type MockFileSystem_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFileSystem) EXPECT() *MockFileSystem_Expecter {
	return &MockFileSystem_Expecter{mock: &_m.Mock}
}

// CreateDir provides a mock function with given fields: path
func (_m *MockFileSystem) CreateDir(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileSystem_CreateDir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDir'
type MockFileSystem_CreateDir_Call struct {
	*mock.Call
}

// CreateDir is a helper method to define mock.On call
//   - path string
func (_e *MockFileSystem_Expecter) CreateDir(path interface{}) *MockFileSystem_CreateDir_Call {
	return &MockFileSystem_CreateDir_Call{Call: _e.mock.On("CreateDir", path)}
}

func (_c *MockFileSystem_CreateDir_Call) Run(run func(path string)) *MockFileSystem_CreateDir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockFileSystem_CreateDir_Call) Return(_a0 error) *MockFileSystem_CreateDir_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileSystem_CreateDir_Call) RunAndReturn(run func(string) error) *MockFileSystem_CreateDir_Call {
	_c.Call.Return(run)
	return _c
}

// CreateFile provides a mock function with given fields: path
func (_m *MockFileSystem) CreateFile(path string) (*os.File, error) {
	ret := _m.Called(path)

	var r0 *os.File
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*os.File, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) *os.File); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSystem_CreateFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateFile'
type MockFileSystem_CreateFile_Call struct {
	*mock.Call
}

// CreateFile is a helper method to define mock.On call
//   - path string
func (_e *MockFileSystem_Expecter) CreateFile(path interface{}) *MockFileSystem_CreateFile_Call {
	return &MockFileSystem_CreateFile_Call{Call: _e.mock.On("CreateFile", path)}
}

func (_c *MockFileSystem_CreateFile_Call) Run(run func(path string)) *MockFileSystem_CreateFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockFileSystem_CreateFile_Call) Return(_a0 *os.File, _a1 error) *MockFileSystem_CreateFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFileSystem_CreateFile_Call) RunAndReturn(run func(string) (*os.File, error)) *MockFileSystem_CreateFile_Call {
	_c.Call.Return(run)
	return _c
}

// Exists provides a mock function with given fields: path
func (_m *MockFileSystem) Exists(path string) bool {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockFileSystem_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type MockFileSystem_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - path string
func (_e *MockFileSystem_Expecter) Exists(path interface{}) *MockFileSystem_Exists_Call {
	return &MockFileSystem_Exists_Call{Call: _e.mock.On("Exists", path)}
}

func (_c *MockFileSystem_Exists_Call) Run(run func(path string)) *MockFileSystem_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockFileSystem_Exists_Call) Return(_a0 bool) *MockFileSystem_Exists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileSystem_Exists_Call) RunAndReturn(run func(string) bool) *MockFileSystem_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFileSystem creates a new instance of MockFileSystem. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFileSystem(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFileSystem {
	mock := &MockFileSystem{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
