// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"

	mock "github.com/stretchr/testify/mock"
)

// CountryService is an autogenerated mock type for the CountryService type
type CountryService struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, filter
func (_m *CountryService) Fetch(ctx context.Context, filter entity.CountryFilter) (entity.CountryFeed, error) {
	ret := _m.Called(ctx, filter)

	var r0 entity.CountryFeed
	if rf, ok := ret.Get(0).(func(context.Context, entity.CountryFilter) entity.CountryFeed); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(entity.CountryFeed)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CountryFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
