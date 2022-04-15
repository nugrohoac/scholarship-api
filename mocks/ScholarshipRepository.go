// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	mock "github.com/stretchr/testify/mock"
)

// ScholarshipRepository is an autogenerated mock type for the ScholarshipRepository type
type ScholarshipRepository struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, userID, scholarshipID, applicant, essay, recommendationLetter
func (_m *ScholarshipRepository) Apply(ctx context.Context, userID int64, scholarshipID int64, applicant int, essay string, recommendationLetter entity.Image) error {
	ret := _m.Called(ctx, userID, scholarshipID, applicant, essay, recommendationLetter)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int, string, entity.Image) error); ok {
		r0 = rf(ctx, userID, scholarshipID, applicant, essay, recommendationLetter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApprovedScholarship provides a mock function with given fields: ctx, status
func (_m *ScholarshipRepository) ApprovedScholarship(ctx context.Context, status int64) error {
	ret := _m.Called(ctx, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeStatus provides a mock function with given fields: ctx, ID, status
func (_m *ScholarshipRepository) ChangeStatus(ctx context.Context, ID int64, status int) error {
	ret := _m.Called(ctx, ID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int) error); ok {
		r0 = rf(ctx, ID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckApply provides a mock function with given fields: ctx, userID, scholarshipID
func (_m *ScholarshipRepository) CheckApply(ctx context.Context, userID int64, scholarshipID int64) (bool, int, error) {
	ret := _m.Called(ctx, userID, scholarshipID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) bool); ok {
		r0 = rf(ctx, userID, scholarshipID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) int); ok {
		r1 = rf(ctx, userID, scholarshipID)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int64, int64) error); ok {
		r2 = rf(ctx, userID, scholarshipID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Create provides a mock function with given fields: ctx, scholarship
func (_m *ScholarshipRepository) Create(ctx context.Context, scholarship entity.Scholarship) (entity.Scholarship, error) {
	ret := _m.Called(ctx, scholarship)

	var r0 entity.Scholarship
	if rf, ok := ret.Get(0).(func(context.Context, entity.Scholarship) entity.Scholarship); ok {
		r0 = rf(ctx, scholarship)
	} else {
		r0 = ret.Get(0).(entity.Scholarship)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.Scholarship) error); ok {
		r1 = rf(ctx, scholarship)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, filter
func (_m *ScholarshipRepository) Fetch(ctx context.Context, filter entity.ScholarshipFilter) ([]entity.Scholarship, string, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.Scholarship
	if rf, ok := ret.Get(0).(func(context.Context, entity.ScholarshipFilter) []entity.Scholarship); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Scholarship)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, entity.ScholarshipFilter) string); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, entity.ScholarshipFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FetchScholarshipBackoffice provides a mock function with given fields: ctx, filter
func (_m *ScholarshipRepository) FetchScholarshipBackoffice(ctx context.Context, filter entity.ScholarshipFilterBackoffice) ([]entity.Scholarship, string, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.Scholarship
	if rf, ok := ret.Get(0).(func(context.Context, entity.ScholarshipFilterBackoffice) []entity.Scholarship); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Scholarship)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, entity.ScholarshipFilterBackoffice) string); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, entity.ScholarshipFilterBackoffice) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *ScholarshipRepository) GetByID(ctx context.Context, ID int64) (entity.Scholarship, error) {
	ret := _m.Called(ctx, ID)

	var r0 entity.Scholarship
	if rf, ok := ret.Get(0).(func(context.Context, int64) entity.Scholarship); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(entity.Scholarship)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MyScholarship provides a mock function with given fields: ctx, userID, filter
func (_m *ScholarshipRepository) MyScholarship(ctx context.Context, userID int64, filter entity.ScholarshipFilter) ([]entity.Applicant, string, error) {
	ret := _m.Called(ctx, userID, filter)

	var r0 []entity.Applicant
	if rf, ok := ret.Get(0).(func(context.Context, int64, entity.ScholarshipFilter) []entity.Applicant); ok {
		r0 = rf(ctx, userID, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Applicant)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, int64, entity.ScholarshipFilter) string); ok {
		r1 = rf(ctx, userID, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int64, entity.ScholarshipFilter) error); ok {
		r2 = rf(ctx, userID, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
