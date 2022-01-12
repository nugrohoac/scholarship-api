package query

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

// UserQuery ...
type UserQuery struct {
	userService sa.UserService
}

// Login ...
func (u UserQuery) Login(ctx context.Context, param sa.InputLogin) (*string, error) {
	token, err := u.userService.Login(ctx, param.Email, param.Password)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// NewUserQuery ...
func NewUserQuery(userService sa.UserService) UserQuery {
	return UserQuery{userService: userService}
}
