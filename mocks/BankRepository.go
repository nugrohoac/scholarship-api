// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"

	mock "github.com/stretchr/testify/mock"
)

// BankRepository is an autogenerated mock type for the BankRepository type
type BankRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, filter
func (_m *BankRepository) Fetch(ctx context.Context, filter entity.BankFilter) ([]entity.Bank, string, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.Bank
	if rf, ok := ret.Get(0).(func(context.Context, entity.BankFilter) []entity.Bank); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Bank)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, entity.BankFilter) string); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, entity.BankFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
