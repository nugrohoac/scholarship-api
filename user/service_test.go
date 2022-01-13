package user_test

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	_user "github.com/Nusantara-Muda/scholarship-api/user"
)

func TestUserServiceStore(t *testing.T) {
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)
	user := users[0]

	tests := map[string]struct {
		paramUser    sa.User
		fetchUser    testdata.FuncCaller
		storeUser    testdata.FuncCaller
		expectedResp sa.User
		expectedErr  error
	}{
		"success": {
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.UserFilter{Email: users[0].Email}},
				Output:   []interface{}{nil, "", nil},
			},
			storeUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, mock.Anything},
				Output:   []interface{}{user, nil},
			},
			expectedResp: user,
			expectedErr:  nil,
		},
		"error fetch user": {
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.UserFilter{Email: users[0].Email}},
				Output:   []interface{}{nil, "", errors.New("internal server error")},
			},
			storeUser:    testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  errors.New("internal server error"),
		},
		"error email already exist": {
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.UserFilter{Email: users[0].Email}},
				Output:   []interface{}{[]sa.User{users[0]}, "cursor", nil},
			},
			storeUser:    testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  sa.ErrorDuplicate{Message: "email already exist"},
		},
		"error": {
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.UserFilter{Email: users[0].Email}},
				Output:   []interface{}{nil, "", nil},
			},
			storeUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, mock.Anything},
				Output:   []interface{}{sa.User{}, errors.New("internal server error")},
			},
			expectedResp: sa.User{},
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)
			jwtHashMock := new(mocks.JwtHash)

			if test.fetchUser.IsCalled {
				userRepoMock.On("Fetch", test.fetchUser.Input...).
					Return(test.fetchUser.Output...).
					Once()
			}

			if test.storeUser.IsCalled {
				userRepoMock.On("Store", test.storeUser.Input...).
					Return(test.storeUser.Output...).
					Once()
			}

			userService := _user.NewUserService(userRepoMock, jwtHashMock)
			userResp, err := userService.Store(context.Background(), test.paramUser)
			userRepoMock.AssertExpectations(t)
			jwtHashMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, userResp)
		})
	}
}

func TestUserServiceLogin(t *testing.T) {
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	user := users[0]
	user.Password = "$2a$14$uftSMoHxUlqPgn5k/PGSO.83RfJxXhrDmv1vUjwR6dH.kCiiv6/ka"

	email := user.Email
	password := "ini password"
	tokenHash := "token-hash"

	tests := map[string]struct {
		paramEmail    string
		paramPassword string
		login         testdata.FuncCaller
		jwtEncode     testdata.FuncCaller
		expectedToken string
		expectedErr   error
	}{
		"success": {
			paramEmail:    user.Email,
			paramPassword: password,
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email},
				Output:   []interface{}{user, nil},
			},
			jwtEncode: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{user},
				Output:   []interface{}{tokenHash, nil},
			},
			expectedToken: tokenHash,
			expectedErr:   nil,
		},
		"error login": {
			paramEmail:    user.Email,
			paramPassword: password,
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email},
				Output:   []interface{}{sa.User{}, errors.New("internal server error")},
			},
			jwtEncode:     testdata.FuncCaller{},
			expectedToken: "",
			expectedErr:   errors.New("internal server error"),
		},
		"invalid password": {
			paramEmail:    user.Email,
			paramPassword: "wrong password",
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email},
				Output:   []interface{}{user, nil},
			},
			jwtEncode:     testdata.FuncCaller{},
			expectedToken: "",
			expectedErr:   bcrypt.ErrMismatchedHashAndPassword,
		},
		"error hash token": {
			paramEmail:    user.Email,
			paramPassword: password,
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email},
				Output:   []interface{}{user, nil},
			},
			jwtEncode: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{user},
				Output:   []interface{}{"", errors.New("internal server error")},
			},
			expectedToken: "",
			expectedErr:   errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)
			jwtHashMock := new(mocks.JwtHash)

			if test.login.IsCalled {
				userRepoMock.On("Login", test.login.Input...).
					Return(test.login.Output...).
					Once()
			}

			if test.jwtEncode.IsCalled {
				jwtHashMock.On("Encode", test.jwtEncode.Input...).
					Return(test.jwtEncode.Output...).
					Once()
			}

			userService := _user.NewUserService(userRepoMock, jwtHashMock)
			token, err := userService.Login(context.Background(), test.paramEmail, test.paramPassword)
			userRepoMock.AssertExpectations(t)
			jwtHashMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.NotEmpty(t, token)
			require.Equal(t, test.expectedToken, token)
		})
	}
}
