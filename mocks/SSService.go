// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	savedsurveys "KitaSehat_Backend/internal/features/saved_surveys"

	mock "github.com/stretchr/testify/mock"
)

// SSService is an autogenerated mock type for the SSService type
type SSService struct {
	mock.Mock
}

// DeleteSavedSurvey provides a mock function with given fields: _a0
func (_m *SSService) DeleteSavedSurvey(_a0 int) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSavedSurvey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllSavedSurveys provides a mock function with given fields: _a0
func (_m *SSService) GetAllSavedSurveys(_a0 int) ([]savedsurveys.SavedSurvey, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAllSavedSurveys")
	}

	var r0 []savedsurveys.SavedSurvey
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]savedsurveys.SavedSurvey, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) []savedsurveys.SavedSurvey); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]savedsurveys.SavedSurvey)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSavedSurvey provides a mock function with given fields: _a0
func (_m *SSService) GetSavedSurvey(_a0 int) (savedsurveys.SavedSurvey, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetSavedSurvey")
	}

	var r0 savedsurveys.SavedSurvey
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (savedsurveys.SavedSurvey, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) savedsurveys.SavedSurvey); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(savedsurveys.SavedSurvey)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveSurvey provides a mock function with given fields: _a0
func (_m *SSService) SaveSurvey(_a0 savedsurveys.SavedSurvey) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SaveSurvey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(savedsurveys.SavedSurvey) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSSService creates a new instance of SSService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSSService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SSService {
	mock := &SSService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
