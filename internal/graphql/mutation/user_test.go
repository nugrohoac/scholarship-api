package mutation_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserMutationRegisterUser(t *testing.T) {
	users := make([]entity.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	user := entity.User{
		Type:     entity.Sponsor,
		Email:    users[0].Email,
		PhoneNo:  users[0].PhoneNo,
		Password: "password",
	}

	userResolver := resolver.UserResolver{User: user}
	inputRegisterUser := entity.InputRegisterUser{
		Type:     users[0].Type,
		Email:    users[0].Email,
		PhoneNo:  users[0].PhoneNo,
		Password: "password",
	}

	tests := map[string]struct {
		paramUser    entity.InputRegisterUser
		storeUser    testdata.FuncCaller
		expectedResp *resolver.UserResolver
		expectedErr  error
	}{
		"success": {
			paramUser: inputRegisterUser,
			storeUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, user},
				Output:   []interface{}{user, nil},
			},
			expectedResp: &userResolver,
			expectedErr:  nil,
		},
		"error": {
			paramUser: inputRegisterUser,
			storeUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, user},
				Output:   []interface{}{entity.User{}, errors.New("internal server error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userServiceMock := new(mocks.UserService)

			if test.storeUser.IsCalled {
				userServiceMock.On("Store", test.storeUser.Input...).
					Return(test.storeUser.Output...).
					Once()
			}

			userMutation := mutation.NewUserMutation(userServiceMock)
			userResolverResp, err := userMutation.RegisterUser(context.Background(), test.paramUser)
			userServiceMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, userResolverResp)
		})
	}
}

func TestUserMutationUpdateUser(t *testing.T) {
	users := make([]entity.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	cardIdentities := make([]entity.CardIdentity, 0)
	testdata.GoldenJSONUnmarshal(t, "card_identities", &cardIdentities)

	for i := range cardIdentities {
		cardIdentities[i].ID = 0
		cardIdentities[i].CreatedAt = time.Time{}
	}

	inputCardIdentities := make([]entity.InputCardIdentity, 0)

	user := entity.User{
		ID:   users[0].ID,
		Name: users[0].Name,
		Photo: entity.Image{
			URL:    users[0].Photo.URL,
			Width:  users[0].Photo.Width,
			Height: users[0].Photo.Height,
		},
		CompanyName:     users[0].CompanyName,
		CountryID:       users[0].CountryID,
		PostalCode:      users[0].PostalCode,
		Address:         users[0].Address,
		CardIdentities:  cardIdentities,
		BankID:          users[0].BankID,
		BankAccountNo:   users[0].BankAccountNo,
		BankAccountName: users[0].BankAccountName,
	}

	for _, c := range user.CardIdentities {
		inputCardIdentities = append(inputCardIdentities, entity.InputCardIdentity{
			Type: c.Type,
			No:   c.No,
			Image: entity.InputImage{
				URL:    c.Image.URL,
				Width:  c.Image.Width,
				Height: c.Image.Height,
			},
		})
	}

	inputUser := entity.InputUpdateUser{
		ID:   int32(users[0].ID),
		Name: users[0].Name,
		Photo: &entity.InputImage{
			URL:    users[0].Photo.URL,
			Width:  users[0].Photo.Width,
			Height: users[0].Photo.Height,
		},
		CompanyName:     &users[0].CompanyName,
		CountryID:       users[0].CountryID,
		Address:         users[0].Address,
		PostalCode:      users[0].PostalCode,
		CardIdentities:  inputCardIdentities,
		BankID:          users[0].BankID,
		BankAccountNo:   users[0].BankAccountNo,
		BankAccountName: users[0].BankAccountName,
	}

	userResolver := resolver.UserResolver{User: user}

	tests := map[string]struct {
		paramInput   entity.InputUpdateUser
		updateUser   testdata.FuncCaller
		expectedResp *resolver.UserResolver
		expectedErr  error
	}{
		"success:": {
			paramInput: inputUser,
			updateUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, user.ID, user},
				Output:   []interface{}{user, nil},
			},
			expectedResp: &userResolver,
			expectedErr:  nil,
		},
		"error": {
			paramInput: inputUser,
			updateUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, user.ID, user},
				Output:   []interface{}{entity.User{}, errors.New("internal server error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userServiceMock := new(mocks.UserService)

			if test.updateUser.IsCalled {
				userServiceMock.On("UpdateByID", test.updateUser.Input...).
					Return(test.updateUser.Output...).
					Once()
			}

			userMutation := mutation.NewUserMutation(userServiceMock)
			userResp, err := userMutation.UpdateUser(context.Background(), test.paramInput)
			userServiceMock.AssertExpectations(t)

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

func TestUserMutationActivateUser(t *testing.T) {
	users := make([]entity.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	token := "token"
	user := users[0]
	userResolver := resolver.UserResolver{User: user}

	tests := map[string]struct {
		paramToken     struct{ Token string }
		activateStatus testdata.FuncCaller
		expectedResp   *resolver.UserResolver
		expectedErr    error
	}{
		"success": {
			paramToken: struct {
				Token string
			}{Token: token},
			activateStatus: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, token},
				Output:   []interface{}{user, nil},
			},
			expectedResp: &userResolver,
			expectedErr:  nil,
		},
		"error": {
			paramToken: struct {
				Token string
			}{Token: token},
			activateStatus: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, token},
				Output:   []interface{}{entity.User{}, errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userServiceMock := new(mocks.UserService)

			if test.activateStatus.IsCalled {
				userServiceMock.On("ActivateStatus", test.activateStatus.Input...).
					Return(test.activateStatus.Output...).
					Once()
			}

			userMutation := mutation.NewUserMutation(userServiceMock)
			message, err := userMutation.ActivateUser(context.Background(), test.paramToken)
			userServiceMock.AssertExpectations(t)

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

func TestUserMutationResetPassword(t *testing.T) {
	var user entity.User
	testdata.GoldenJSONUnmarshal(t, "user", &user)

	password := "new password"
	userResolver := resolver.UserResolver{User: user}

	tests := map[string]struct {
		paramPassword struct{ Password string }
		resetPasswd   testdata.FuncCaller
		expectedResp  *resolver.UserResolver
		expectedErr   error
	}{
		"success": {
			paramPassword: struct {
				Password string
			}{Password: password},
			resetPasswd: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, password},
				Output:   []interface{}{user, nil},
			},
			expectedResp: &userResolver,
			expectedErr:  nil,
		},
		"error": {
			paramPassword: struct {
				Password string
			}{Password: password},
			resetPasswd: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, password},
				Output:   []interface{}{entity.User{}, errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			userServiceMock := new(mocks.UserService)

			if test.resetPasswd.IsCalled {
				userServiceMock.On("ResetPassword", test.resetPasswd.Input...).
					Return(test.resetPasswd.Output...).
					Once()
			}

			userMutation := mutation.NewUserMutation(userServiceMock)
			_message, err := userMutation.ResetPassword(context.Background(), test.paramPassword)
			userServiceMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, _message)
		})
	}
}
