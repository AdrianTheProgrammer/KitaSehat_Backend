// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt/v5"
	mock "github.com/stretchr/testify/mock"

	users "KitaSehat_Backend/internal/features/users"
)

// TokenUtilityInterface is an autogenerated mock type for the TokenUtilityInterface type
type TokenUtilityInterface struct {
	mock.Mock
}

// DecodeToken provides a mock function with given fields: _a0
func (_m *TokenUtilityInterface) DecodeToken(_a0 *jwt.Token) users.User {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DecodeToken")
	}

	var r0 users.User
	if rf, ok := ret.Get(0).(func(*jwt.Token) users.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	return r0
}

// GenerateToken provides a mock function with given fields: _a0
func (_m *TokenUtilityInterface) GenerateToken(_a0 users.User) (string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GenerateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(users.User) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(users.User) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(users.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenUtilityInterface creates a new instance of TokenUtilityInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenUtilityInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenUtilityInterface {
	mock := &TokenUtilityInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
