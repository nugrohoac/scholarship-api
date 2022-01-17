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

// Fetch provides a mock function with given fields: ctx, filter
func (_m *UserRepository) Fetch(ctx context.Context, filter scholarship_api.UserFilter) ([]scholarship_api.User, string, error) {
	ret := _m.Called(ctx, filter)

	var r0 []scholarship_api.User
	if rf, ok := ret.Get(0).(func(context.Context, scholarship_api.UserFilter) []scholarship_api.User); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]scholarship_api.User)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, scholarship_api.UserFilter) string); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, scholarship_api.UserFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Login provides a mock function with given fields: ctx, email
func (_m *UserRepository) Login(ctx context.Context, email string) (scholarship_api.User, error) {
	ret := _m.Called(ctx, email)

	var r0 scholarship_api.User
	if rf, ok := ret.Get(0).(func(context.Context, string) scholarship_api.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(scholarship_api.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetStatus provides a mock function with given fields: ctx, ID, status
func (_m *UserRepository) SetStatus(ctx context.Context, ID int64, status int) error {
	ret := _m.Called(ctx, ID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int) error); ok {
		r0 = rf(ctx, ID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

// UpdateByID provides a mock function with given fields: ctx, ID, user
func (_m *UserRepository) UpdateByID(ctx context.Context, ID int64, user scholarship_api.User) (scholarship_api.User, error) {
	ret := _m.Called(ctx, ID, user)

	var r0 scholarship_api.User
	if rf, ok := ret.Get(0).(func(context.Context, int64, scholarship_api.User) scholarship_api.User); ok {
		r0 = rf(ctx, ID, user)
	} else {
		r0 = ret.Get(0).(scholarship_api.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, scholarship_api.User) error); ok {
		r1 = rf(ctx, ID, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
