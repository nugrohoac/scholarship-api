// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"

	mock "github.com/stretchr/testify/mock"
)

// MajorRepository is an autogenerated mock type for the MajorRepository type
type MajorRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, filter
func (_m *MajorRepository) Fetch(ctx context.Context, filter entity.MajorFilter) ([]entity.Major, string, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.Major
	if rf, ok := ret.Get(0).(func(context.Context, entity.MajorFilter) []entity.Major); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Major)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, entity.MajorFilter) string); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, entity.MajorFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
