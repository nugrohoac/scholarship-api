package query

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

// UserQuery ...
type UserQuery struct {
	userService sa.UserService
}

// Login ...
func (u UserQuery) Login(ctx context.Context, param sa.InputLogin) (*resolver.LoginResponseResolver, error) {
	loginResponse, err := u.userService.Login(ctx, param.Email, param.Password)
	if err != nil {
		return nil, err
	}

	return &resolver.LoginResponseResolver{LoginResponse: loginResponse}, nil
}

// NewUserQuery ...
func NewUserQuery(userService sa.UserService) UserQuery {
	return UserQuery{userService: userService}
}
