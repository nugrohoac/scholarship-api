package query

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// UserQuery ...
type UserQuery struct {
	userService business.UserService
}

// Login ...
func (u UserQuery) Login(ctx context.Context, param entity.InputLogin) (*resolver.LoginResponseResolver, error) {
	loginResponse, err := u.userService.Login(ctx, param.Email, param.Password)
	if err != nil {
		return nil, err
	}

	return &resolver.LoginResponseResolver{LoginResponse: loginResponse}, nil
}

// ResendEmailVerification ...
func (u UserQuery) ResendEmailVerification(ctx context.Context, param struct{ Email string }) (*string, error) {
	message, err := u.userService.ResendEmailVerification(ctx, param.Email)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// ForgotPassword ...
func (u UserQuery) ForgotPassword(ctx context.Context, param struct{ Email string }) (*string, error) {
	message, err := u.userService.ForgotPassword(ctx, param.Email)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// NewUserQuery ...
func NewUserQuery(userService business.UserService) UserQuery {
	return UserQuery{userService: userService}
}
