// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	scholarship_api "github.com/Nusantara-Muda/scholarship-api"
	mock "github.com/stretchr/testify/mock"
)

// ScholarshipRepository is an autogenerated mock type for the ScholarshipRepository type
type ScholarshipRepository struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, userID, scholarshipID, applicant, essay, recommendationLetter
func (_m *ScholarshipRepository) Apply(ctx context.Context, userID int64, scholarshipID int64, applicant int, essay string, recommendationLetter scholarship_api.Image) error {
	ret := _m.Called(ctx, userID, scholarshipID, applicant, essay, recommendationLetter)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int, string, scholarship_api.Image) error); ok {
		r0 = rf(ctx, userID, scholarshipID, applicant, essay, recommendationLetter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, scholarship
func (_m *ScholarshipRepository) Create(ctx context.Context, scholarship scholarship_api.Scholarship) (scholarship_api.Scholarship, error) {
	ret := _m.Called(ctx, scholarship)

	var r0 scholarship_api.Scholarship
	if rf, ok := ret.Get(0).(func(context.Context, scholarship_api.Scholarship) scholarship_api.Scholarship); ok {
		r0 = rf(ctx, scholarship)
	} else {
		r0 = ret.Get(0).(scholarship_api.Scholarship)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, scholarship_api.Scholarship) error); ok {
		r1 = rf(ctx, scholarship)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, filter
func (_m *ScholarshipRepository) Fetch(ctx context.Context, filter scholarship_api.ScholarshipFilter) ([]scholarship_api.Scholarship, string, error) {
	ret := _m.Called(ctx, filter)

	var r0 []scholarship_api.Scholarship
	if rf, ok := ret.Get(0).(func(context.Context, scholarship_api.ScholarshipFilter) []scholarship_api.Scholarship); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]scholarship_api.Scholarship)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, scholarship_api.ScholarshipFilter) string); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, scholarship_api.ScholarshipFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *ScholarshipRepository) GetByID(ctx context.Context, ID int64) (scholarship_api.Scholarship, error) {
	ret := _m.Called(ctx, ID)

	var r0 scholarship_api.Scholarship
	if rf, ok := ret.Get(0).(func(context.Context, int64) scholarship_api.Scholarship); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(scholarship_api.Scholarship)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
