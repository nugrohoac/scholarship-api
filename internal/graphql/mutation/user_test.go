package mutation_test

import (
	"context"
	"errors"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserMutationRegisterUser(t *testing.T) {
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	users[0].CompanyName = ""
	users[0].CreatedAt = time.Time{}

	userResolver := resolver.UserResolver{User: users[0]}
	inputRegisterUser := sa.InputRegisterUser{
		Name:    users[0].Name,
		Type:    users[0].Type,
		Email:   users[0].Email,
		PhoneNo: users[0].PhoneNo,
		Photo: sa.InputImage{
			URL:    users[0].Photo.URL,
			Width:  users[0].Photo.Width,
			Height: users[0].Photo.Height,
		},
		CountryID:       users[0].CountryID,
		PostalCode:      users[0].PostalCode,
		Address:         users[0].Address,
		BankID:          users[0].BankID,
		BankAccountNo:   users[0].BankAccountNo,
		BankAccountName: users[0].BankAccountName,
	}

	tests := map[string]struct {
		paramUser    sa.InputRegisterUser
		storeUser    testdata.FuncCaller
		expectedResp *resolver.UserResolver
		expectedErr  error
	}{
		"success": {
			paramUser: inputRegisterUser,
			storeUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, users[0]},
				Output:   []interface{}{users[0], nil},
			},
			expectedResp: &userResolver,
			expectedErr:  nil,
		},
		"error": {
			paramUser: inputRegisterUser,
			storeUser: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, users[0]},
				Output:   []interface{}{sa.User{}, errors.New("internal server error")},
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
