// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	scholarship_api "github.com/Nusantara-Muda/scholarship-api"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// ActivateStatus provides a mock function with given fields: ctx, ID
func (_m *UserService) ActivateStatus(ctx context.Context, ID int64) (string, error) {
	ret := _m.Called(ctx, ID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, int64) string); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, email, password
func (_m *UserService) Login(ctx context.Context, email string, password string) (scholarship_api.LoginResponse, error) {
	ret := _m.Called(ctx, email, password)

	var r0 scholarship_api.LoginResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string) scholarship_api.LoginResponse); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(scholarship_api.LoginResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, user
func (_m *UserService) Store(ctx context.Context, user scholarship_api.User) (scholarship_api.User, error) {
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
func (_m *UserService) UpdateByID(ctx context.Context, ID int64, user scholarship_api.User) (scholarship_api.User, error) {
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
