// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	mock "github.com/stretchr/testify/mock"
)

// SchoolRepository is an autogenerated mock type for the SchoolRepository type
type SchoolRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, school
func (_m *SchoolRepository) Create(ctx context.Context, school entity.School) (entity.School, error) {
	ret := _m.Called(ctx, school)

	var r0 entity.School
	if rf, ok := ret.Get(0).(func(context.Context, entity.School) entity.School); ok {
		r0 = rf(ctx, school)
	} else {
		r0 = ret.Get(0).(entity.School)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.School) error); ok {
		r1 = rf(ctx, school)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, filter
func (_m *SchoolRepository) Fetch(ctx context.Context, filter entity.SchoolFilter) ([]entity.School, string, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.School
	if rf, ok := ret.Get(0).(func(context.Context, entity.SchoolFilter) []entity.School); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.School)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, entity.SchoolFilter) string); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, entity.SchoolFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUserSchool provides a mock function with given fields: ctx, userID
func (_m *SchoolRepository) GetUserSchool(ctx context.Context, userID int64) ([]entity.UserSchool, error) {
	ret := _m.Called(ctx, userID)

	var r0 []entity.UserSchool
	if rf, ok := ret.Get(0).(func(context.Context, int64) []entity.UserSchool); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.UserSchool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
