package mutation

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// UserMutation ...
type UserMutation struct {
	userService sa.UserService
}

// RegisterUser ...
func (u UserMutation) RegisterUser(ctx context.Context, param sa.InputRegisterUser) (*resolver.UserResolver, error) {
	user := sa.User{
		Type:     param.Type,
		Email:    param.Email,
		PhoneNo:  param.PhoneNo,
		Password: param.Password,
	}

	user, err := u.userService.Store(ctx, user)
	if err != nil {
		return nil, err
	}

	return &resolver.UserResolver{User: user}, err
}

// NewUserMutation ...
func NewUserMutation(userService sa.UserService) UserMutation {
	return UserMutation{userService: userService}
}
