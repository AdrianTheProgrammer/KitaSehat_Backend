// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	surveys "KitaSehat_Backend/internal/features/surveys"

	mock "github.com/stretchr/testify/mock"
)

// SService is an autogenerated mock type for the SService type
type SService struct {
	mock.Mock
}

// AddSurvey provides a mock function with given fields: _a0
func (_m *SService) AddSurvey(_a0 surveys.Survey) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AddSurvey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(surveys.Survey) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSurvey provides a mock function with given fields: _a0
func (_m *SService) DeleteSurvey(_a0 int) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSurvey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllSurveys provides a mock function with given fields:
func (_m *SService) GetAllSurveys() ([]surveys.Survey, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllSurveys")
	}

	var r0 []surveys.Survey
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]surveys.Survey, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []surveys.Survey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]surveys.Survey)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSurvey provides a mock function with given fields: _a0
func (_m *SService) GetSurvey(_a0 int) (surveys.Survey, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetSurvey")
	}

	var r0 surveys.Survey
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (surveys.Survey, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) surveys.Survey); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(surveys.Survey)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSurvey provides a mock function with given fields: _a0
func (_m *SService) UpdateSurvey(_a0 surveys.Survey) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSurvey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(surveys.Survey) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSService creates a new instance of SService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SService {
	mock := &SService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}