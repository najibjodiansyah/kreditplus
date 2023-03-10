// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
	domain "github.com/najibjodiansyah/kreditplus/domain"
)

// MockUserUseCase is an autogenerated mock type for the MockUserUsecase type
type MockUserUseCase struct {
	mock.Mock
}

type MockUserUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserUseCase) EXPECT() *MockUserUsecase_Expecter {
	return &MockUserUsecase_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, us
func (_m *MockUserUseCase) Create(ctx echo.Context, us domain.User) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, domain.User) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserUsecase_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUserUsecase_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - ctx echo.Context
//  - us domain.User
func (_e *MockUserUsecase_Expecter) Create(ctx interface{}, us interface{}) *MockUserUsecase_Create_Call {
	return &MockUserUsecase_Create_Call{Call: _e.mock.On("Create", ctx, us)}
}

func (_c *MockUserUsecase_Create_Call) Run(run func(ctx echo.Context, us domain.User)) *MockUserUsecase_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(domain.User))
	})
	return _c
}

func (_c *MockUserUsecase_Create_Call) Return(_a0 error) *MockUserUsecase_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserUsecase_Create_Call) RunAndReturn(run func(echo.Context, domain.User) error) *MockUserUsecase_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: ctx, nik, pass
func (_m *MockUserUseCase) Login(ctx echo.Context, nik string, pass string) (string, error) {
	ret := _m.Called(ctx, nik, pass)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, string, string) (string, error)); ok {
		return rf(ctx, nik, pass)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, string, string) string); ok {
		r0 = rf(ctx, nik, pass)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, string, string) error); ok {
		r1 = rf(ctx, nik, pass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserUsecase_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type MockUserUsecase_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//  - ctx echo.Context
//  - nik string
//  - pass string
func (_e *MockUserUsecase_Expecter) Login(ctx interface{}, nik interface{}, pass interface{}) *MockUserUsecase_Login_Call {
	return &MockUserUsecase_Login_Call{Call: _e.mock.On("Login", ctx, nik, pass)}
}

func (_c *MockUserUsecase_Login_Call) Run(run func(ctx echo.Context, nik string, pass string)) *MockUserUsecase_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockUserUsecase_Login_Call) Return(_a0 string, _a1 error) *MockUserUsecase_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserUsecase_Login_Call) RunAndReturn(run func(echo.Context, string, string) (string, error)) *MockUserUsecase_Login_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, us
func (_m *MockUserUseCase) Update(ctx echo.Context, us domain.User) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, domain.User) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserUsecase_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUserUsecase_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - ctx echo.Context
//  - us domain.User
func (_e *MockUserUsecase_Expecter) Update(ctx interface{}, us interface{}) *MockUserUsecase_Update_Call {
	return &MockUserUsecase_Update_Call{Call: _e.mock.On("Update", ctx, us)}
}

func (_c *MockUserUsecase_Update_Call) Run(run func(ctx echo.Context, us domain.User)) *MockUserUsecase_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(domain.User))
	})
	return _c
}

func (_c *MockUserUsecase_Update_Call) Return(_a0 error) *MockUserUsecase_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserUsecase_Update_Call) RunAndReturn(run func(echo.Context, domain.User) error) *MockUserUsecase_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUserUsecase creates a new instance of MockUserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUserUsecase(t mockConstructorTestingTNewMockUserUsecase) *MockUserUseCase {
	mock := &MockUserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
