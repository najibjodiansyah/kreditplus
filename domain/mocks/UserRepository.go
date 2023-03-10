// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
	domain "github.com/najibjodiansyah/kreditplus/domain"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, us
func (_m *MockUserRepository) Create(ctx echo.Context, us domain.User) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, domain.User) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUserRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - ctx echo.Context
//  - us domain.User
func (_e *MockUserRepository_Expecter) Create(ctx interface{}, us interface{}) *MockUserRepository_Create_Call {
	return &MockUserRepository_Create_Call{Call: _e.mock.On("Create", ctx, us)}
}

func (_c *MockUserRepository_Create_Call) Run(run func(ctx echo.Context, us domain.User)) *MockUserRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(domain.User))
	})
	return _c
}

func (_c *MockUserRepository_Create_Call) Return(_a0 error) *MockUserRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Create_Call) RunAndReturn(run func(echo.Context, domain.User) error) *MockUserRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: ctx, nik
func (_m *MockUserRepository) Login(ctx echo.Context, nik string) (domain.User, error) {
	ret := _m.Called(ctx, nik)

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) (domain.User, error)); ok {
		return rf(ctx, nik)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, string) domain.User); ok {
		r0 = rf(ctx, nik)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, string) error); ok {
		r1 = rf(ctx, nik)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type MockUserRepository_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//  - ctx echo.Context
//  - nik string
func (_e *MockUserRepository_Expecter) Login(ctx interface{}, nik interface{}) *MockUserRepository_Login_Call {
	return &MockUserRepository_Login_Call{Call: _e.mock.On("Login", ctx, nik)}
}

func (_c *MockUserRepository_Login_Call) Run(run func(ctx echo.Context, nik string)) *MockUserRepository_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserRepository_Login_Call) Return(_a0 domain.User, _a1 error) *MockUserRepository_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_Login_Call) RunAndReturn(run func(echo.Context, string) (domain.User, error)) *MockUserRepository_Login_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, us
func (_m *MockUserRepository) Update(ctx echo.Context, us domain.User) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, domain.User) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUserRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - ctx echo.Context
//  - us domain.User
func (_e *MockUserRepository_Expecter) Update(ctx interface{}, us interface{}) *MockUserRepository_Update_Call {
	return &MockUserRepository_Update_Call{Call: _e.mock.On("Update", ctx, us)}
}

func (_c *MockUserRepository_Update_Call) Run(run func(ctx echo.Context, us domain.User)) *MockUserRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(domain.User))
	})
	return _c
}

func (_c *MockUserRepository_Update_Call) Return(_a0 error) *MockUserRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Update_Call) RunAndReturn(run func(echo.Context, domain.User) error) *MockUserRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUserRepository(t mockConstructorTestingTNewMockUserRepository) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
