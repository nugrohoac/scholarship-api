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
)

func TestUserMutationRegisterUser(t *testing.T) {
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	user := sa.User{
		Type:     sa.Sponsor,
		Email:    users[0].Email,
		PhoneNo:  users[0].PhoneNo,
		Password: "password",
	}

	userResolver := resolver.UserResolver{User: user}
	inputRegisterUser := sa.InputRegisterUser{
		Type:     users[0].Type,
		Email:    users[0].Email,
		PhoneNo:  users[0].PhoneNo,
		Password: "password",
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
