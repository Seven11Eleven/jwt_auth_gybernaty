// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	mock "github.com/stretchr/testify/mock"
)

// LoginService is an autogenerated mock type for the LoginService type
type LoginService struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: author, expired
func (_m *LoginService) CreateAccessToken(author *domain.Author, expired int) (string, error) {
	ret := _m.Called(author, expired)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccessToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Author, int) (string, error)); ok {
		return rf(author, expired)
	}
	if rf, ok := ret.Get(0).(func(*domain.Author, int) string); ok {
		r0 = rf(author, expired)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.Author, int) error); ok {
		r1 = rf(author, expired)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: author, expired
func (_m *LoginService) CreateRefreshToken(author *domain.Author, expired int) (string, error) {
	ret := _m.Called(author, expired)

	if len(ret) == 0 {
		panic("no return value specified for CreateRefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Author, int) (string, error)); ok {
		return rf(author, expired)
	}
	if rf, ok := ret.Get(0).(func(*domain.Author, int) string); ok {
		r0 = rf(author, expired)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.Author, int) error); ok {
		r1 = rf(author, expired)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: ctx, username
func (_m *LoginService) GetUserByUsername(ctx context.Context, username string) (*domain.Author, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsername")
	}

	var r0 *domain.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.Author, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Author); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLoginService creates a new instance of LoginService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoginService(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoginService {
	mock := &LoginService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
