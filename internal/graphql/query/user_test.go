package query

import (
	"context"
	"errors"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserQueryLogin(t *testing.T) {
	var (
		email    = "jhon@wick.com"
		password = "password"
		token    = "token"
	)

	tests := map[string]struct {
		paramLogin    sa.InputLogin
		login         testdata.FuncCaller
		expectedToken *string
		expectedErr   error
	}{
		"success": {
			paramLogin: sa.InputLogin{Email: email, Password: password},
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email, password},
				Output:   []interface{}{token, nil},
			},
			expectedToken: &token,
			expectedErr:   nil,
		},
		"error": {
			paramLogin: sa.InputLogin{Email: email, Password: password},
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email, password},
				Output:   []interface{}{"", errors.New("internal server error")},
			},
			expectedToken: &token,
			expectedErr:   errors.New("internal server error"),
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
			require.Equal(t, test.expectedToken, tokenResp)
		})
	}
}
