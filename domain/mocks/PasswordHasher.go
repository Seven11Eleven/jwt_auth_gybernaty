// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PasswordHasher is an autogenerated mock type for the PasswordHasher type
type PasswordHasher struct {
	mock.Mock
}

// CompareHashAndPassword provides a mock function with given fields: hashedPassword, password, salt, localParam
func (_m *PasswordHasher) CompareHashAndPassword(hashedPassword string, password string, salt string, localParam string) error {
	ret := _m.Called(hashedPassword, password, salt, localParam)

	if len(ret) == 0 {
		panic("no return value specified for CompareHashAndPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(hashedPassword, password, salt, localParam)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateSalt provides a mock function with given fields:
func (_m *PasswordHasher) GenerateSalt() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GenerateSalt")
	}

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

// HashPassword provides a mock function with given fields: password, salt, localParam
func (_m *PasswordHasher) HashPassword(password string, salt string, localParam string) (string, error) {
	ret := _m.Called(password, salt, localParam)

	if len(ret) == 0 {
		panic("no return value specified for HashPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (string, error)); ok {
		return rf(password, salt, localParam)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(password, salt, localParam)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(password, salt, localParam)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPasswordHasher creates a new instance of PasswordHasher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPasswordHasher(t interface {
	mock.TestingT
	Cleanup(func())
}) *PasswordHasher {
	mock := &PasswordHasher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
