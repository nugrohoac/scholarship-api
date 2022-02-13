// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	scholarship_api "github.com/Nusantara-Muda/scholarship-api"
	mock "github.com/stretchr/testify/mock"
)

// PaymentRepository is an autogenerated mock type for the PaymentRepository type
type PaymentRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, scholarshipIDs
func (_m *PaymentRepository) Fetch(ctx context.Context, scholarshipIDs []int64) ([]scholarship_api.Payment, error) {
	ret := _m.Called(ctx, scholarshipIDs)

	var r0 []scholarship_api.Payment
	if rf, ok := ret.Get(0).(func(context.Context, []int64) []scholarship_api.Payment); ok {
		r0 = rf(ctx, scholarshipIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]scholarship_api.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []int64) error); ok {
		r1 = rf(ctx, scholarshipIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
