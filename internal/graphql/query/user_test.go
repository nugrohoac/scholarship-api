package query

import (
	"context"
	"errors"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	email    = "jhon@wick.com"
	password = "password"
)

func TestUserQueryLogin(t *testing.T) {
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	loginResp := sa.LoginResponse{
		Token: "token",
		User:  users[0],
	}

	loginResponseResolver := resolver.LoginResponseResolver{
		LoginResponse: loginResp,
	}

	tests := map[string]struct {
		paramLogin   sa.InputLogin
		login        testdata.FuncCaller
		expectedResp *resolver.LoginResponseResolver
		expectedErr  error
	}{
		"success": {
			paramLogin: sa.InputLogin{Email: email, Password: password},
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email, password},
				Output:   []interface{}{loginResp, nil},
			},
			expectedResp: &loginResponseResolver,
			expectedErr:  nil,
		},
		"error": {
			paramLogin: sa.InputLogin{Email: email, Password: password},
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email, password},
				Output:   []interface{}{sa.LoginResponse{}, errors.New("internal server error")},
			},
			expectedResp: &loginResponseResolver,
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userServiceMock := new(mocks.UserService)

			if test.login.IsCalled {
				userServiceMock.On("Login", test.login.Input...).
					Return(test.login.Output...).
					Once()
			}

			userQuery := NewUserQuery(userServiceMock)
			tokenResp, err := userQuery.Login(context.Background(), test.paramLogin)
			userServiceMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, tokenResp)
		})
	}
}

func TestUserQueryResendEmailVerification(t *testing.T) {
	message := "success"

	tests := map[string]struct {
		paramEmail   struct{ Email string }
		resendEmail  testdata.FuncCaller
		expectedResp *string
		expectedErr  error
	}{
		"success": {
			paramEmail: struct {
				Email string
			}{Email: email},
			resendEmail: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email},
				Output:   []interface{}{message, nil},
			},
			expectedResp: &message,
			expectedErr:  nil,
		},
		"error": {
			paramEmail: struct {
				Email string
			}{Email: email},
			resendEmail: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email},
				Output:   []interface{}{"", errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userServiceMock := new(mocks.UserService)

			if test.resendEmail.IsCalled {
				userServiceMock.On("ResendEmailVerification", test.resendEmail.Input...).
					Return(test.resendEmail.Output...).
					Once()
			}

			userQuery := NewUserQuery(userServiceMock)
			resp, err := userQuery.ResendEmailVerification(context.Background(), test.paramEmail)
			userServiceMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, resp)
		})
	}
}
