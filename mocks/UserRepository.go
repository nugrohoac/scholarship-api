// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	scholarship_api "github.com/Nusantara-Muda/scholarship-api"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Store provides a mock function with given fields: ctx, user
func (_m *UserRepository) Store(ctx context.Context, user scholarship_api.User) (scholarship_api.User, error) {
	ret := _m.Called(ctx, user)

	var r0 scholarship_api.User
	if rf, ok := ret.Get(0).(func(context.Context, scholarship_api.User) scholarship_api.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(scholarship_api.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, scholarship_api.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
