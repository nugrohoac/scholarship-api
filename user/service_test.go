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

var cursor = "cursor"

func TestUserServiceStore(t *testing.T) {
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)
	user := users[0]

	tests := map[string]struct {
		paramUser    sa.User
		fetchUser    testdata.FuncCaller
		storeUser    testdata.FuncCaller
		sendEmail    testdata.FuncCaller
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
			sendEmail: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, user.Email},
				Output:   []interface{}{nil},
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
			sendEmail:    testdata.FuncCaller{},
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
			sendEmail:    testdata.FuncCaller{},
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
			sendEmail: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, user.Email},
				Output:   []interface{}{nil},
			},
			expectedResp: sa.User{},
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)
			jwtHashMock := new(mocks.JwtHash)
			emailRepoMock := new(mocks.EmailRepository)

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

			if test.sendEmail.IsCalled {
				emailRepoMock.On("SendActivateUser", test.sendEmail.Input...).
					Return(test.sendEmail.Output...).
					Once()
			}

			userService := _user.NewUserService(userRepoMock, jwtHashMock, emailRepoMock)
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

	loginResponse := sa.LoginResponse{Token: tokenHash, User: user}

	tests := map[string]struct {
		paramEmail    string
		paramPassword string
		login         testdata.FuncCaller
		jwtEncode     testdata.FuncCaller
		expectedResp  sa.LoginResponse
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
			expectedResp: loginResponse,
			expectedErr:  nil,
		},
		"error login": {
			paramEmail:    user.Email,
			paramPassword: password,
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email},
				Output:   []interface{}{sa.User{}, errors.New("internal server error")},
			},
			jwtEncode:    testdata.FuncCaller{},
			expectedResp: loginResponse,
			expectedErr:  errors.New("internal server error"),
		},
		"invalid password": {
			paramEmail:    user.Email,
			paramPassword: "wrong password",
			login: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, email},
				Output:   []interface{}{user, nil},
			},
			jwtEncode:    testdata.FuncCaller{},
			expectedResp: loginResponse,
			expectedErr:  bcrypt.ErrMismatchedHashAndPassword,
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
			expectedResp: sa.LoginResponse{},
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)
			jwtHashMock := new(mocks.JwtHash)
			emailRepoMock := new(mocks.EmailRepository)

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

			userService := _user.NewUserService(userRepoMock, jwtHashMock, emailRepoMock)
			response, err := userService.Login(context.Background(), test.paramEmail, test.paramPassword)
			userRepoMock.AssertExpectations(t)
			jwtHashMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, response)
		})
	}
}

func TestUserServiceUpdateByID(t *testing.T) {
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	user := users[0]
	userInvalid := user
	userInvalid.Email = "email@invalid.com"

	ctx := sa.SetUserOnContext(context.Background(), user)
	ctxInvalid := sa.SetUserOnContext(context.Background(), userInvalid)

	tests := map[string]struct {
		paramCtx     context.Context
		paramID      int64
		paramUser    sa.User
		fetchUser    testdata.FuncCaller
		updateUser   testdata.FuncCaller
		expectedResp sa.User
		expectedErr  error
	}{
		"success": {
			paramCtx:  ctx,
			paramID:   user.ID,
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctx, sa.UserFilter{IDs: []int64{user.ID}}},
				Output:   []interface{}{users, cursor, nil},
			},
			updateUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctx, user.ID, user},
				Output:   []interface{}{user, nil},
			},
			expectedResp: user,
			expectedErr:  nil,
		},
		"error fetch user": {
			paramCtx:  ctx,
			paramID:   user.ID,
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctx, sa.UserFilter{IDs: []int64{user.ID}}},
				Output:   []interface{}{nil, "", errors.New("internal server error")},
			},
			updateUser:   testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  errors.New("internal server error"),
		},
		"user not found": {
			paramCtx:  ctx,
			paramID:   user.ID,
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctx, sa.UserFilter{IDs: []int64{user.ID}}},
				Output:   []interface{}{nil, "", nil},
			},
			updateUser:   testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  sa.ErrNotFound{Message: "user not found"},
		},
		"error get user on context": {
			paramCtx:  context.Background(),
			paramID:   user.ID,
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{context.Background(), sa.UserFilter{IDs: []int64{user.ID}}},
				Output:   []interface{}{users, cursor, nil},
			},
			updateUser:   testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  sa.ErrBadRequest{Message: "failed casting key to string"},
		},
		"error user is not sync": {
			paramCtx:  ctxInvalid,
			paramID:   user.ID,
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctxInvalid, sa.UserFilter{IDs: []int64{user.ID}}},
				Output:   []interface{}{users, cursor, nil},
			},
			updateUser:   testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  sa.ErrUnAuthorize{Message: "user is not sync"},
		},
		"error update user": {
			paramCtx:  ctx,
			paramID:   user.ID,
			paramUser: user,
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctx, sa.UserFilter{IDs: []int64{user.ID}}},
				Output:   []interface{}{users, cursor, nil},
			},
			updateUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctx, user.ID, user},
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
			emailRepoMock := new(mocks.EmailRepository)

			if test.fetchUser.IsCalled {
				userRepoMock.On("Fetch", test.fetchUser.Input...).
					Return(test.fetchUser.Output...).
					Once()
			}

			if test.updateUser.IsCalled {
				userRepoMock.On("UpdateByID", test.updateUser.Input...).
					Return(test.updateUser.Output...).
					Once()
			}

			userService := _user.NewUserService(userRepoMock, jwtHashMock, emailRepoMock)
			userResp, err := userService.UpdateByID(test.paramCtx, test.paramID, test.paramUser)

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

func TestUserServiceActivateStatus(t *testing.T) {
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	user := users[0]
	userInvalid := user
	userInvalid.Email = "email@invalid.com"

	token := "token"

	tests := map[string]struct {
		paramToken   string
		jwtDecode    testdata.FuncCaller
		fetchUser    testdata.FuncCaller
		setStatus    testdata.FuncCaller
		expectedResp sa.User
		expectedErr  error
	}{
		"success": {
			paramToken: token,
			jwtDecode: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{token, mock.Anything},
				Output:   []interface{}{nil},
			},
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.UserFilter{Email: user.Email}},
				Output:   []interface{}{users, cursor, nil},
			},
			setStatus: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, user.ID, 1},
				Output:   []interface{}{nil},
			},
			expectedResp: user,
			expectedErr:  nil,
		},
		"error decode": {
			paramToken: token,
			jwtDecode: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{token, mock.Anything},
				Output:   []interface{}{errors.New("internal server error")},
			},
			fetchUser:    testdata.FuncCaller{},
			setStatus:    testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  errors.New("internal server error"),
		},
		"error fetch user": {
			paramToken: token,
			jwtDecode: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{token, mock.Anything},
				Output:   []interface{}{nil},
			},
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.UserFilter{Email: user.Email}},
				Output:   []interface{}{nil, "", errors.New("error")},
			},
			setStatus:    testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  errors.New("error"),
		},
		"user not found": {
			paramToken: token,
			jwtDecode: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{token, mock.Anything},
				Output:   []interface{}{nil},
			},
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.UserFilter{Email: user.Email}},
				Output:   []interface{}{nil, "", nil},
			},
			setStatus:    testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  sa.ErrNotFound{Message: "user not found"},
		},
		"user not sync": {
			paramToken: token,
			jwtDecode: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{token, mock.Anything},
				Output:   []interface{}{nil},
			},
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.UserFilter{Email: user.Email}},
				Output:   []interface{}{[]sa.User{userInvalid}, "", nil},
			},
			setStatus:    testdata.FuncCaller{},
			expectedResp: sa.User{},
			expectedErr:  sa.ErrUnAuthorize{Message: "user is not sync"},
		},
		"error set status": {
			paramToken: token,
			jwtDecode: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{token, mock.Anything},
				Output:   []interface{}{nil},
			},
			fetchUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.UserFilter{Email: user.Email}},
				Output:   []interface{}{users, cursor, nil},
			},
			setStatus: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, user.ID, 1},
				Output:   []interface{}{errors.New("error")},
			},
			expectedResp: sa.User{},
			expectedErr:  errors.New("error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userRepoMock := new(mocks.UserRepository)
			jwtHashMock := new(mocks.JwtHash)
			emailRepoMock := new(mocks.EmailRepository)

			if test.jwtDecode.IsCalled {
				jwtHashMock.On("Decode", test.jwtDecode.Input...).
					Return(test.jwtDecode.Output...).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(*sa.Claim)
						arg.Email = user.Email
						arg.Name = user.Name
						arg.Type = user.Type
						arg.Status = user.Status
					}).Once()
			}

			if test.fetchUser.IsCalled {
				userRepoMock.On("Fetch", test.fetchUser.Input...).
					Return(test.fetchUser.Output...).
					Once()
			}

			if test.setStatus.IsCalled {
				userRepoMock.On("SetStatus", test.setStatus.Input...).
					Return(test.setStatus.Output...).
					Once()
			}

			userService := _user.NewUserService(userRepoMock, jwtHashMock, emailRepoMock)
			message, err := userService.ActivateStatus(context.Background(), token)
			userRepoMock.AssertExpectations(t)
			jwtHashMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, message)
		})
	}
}
