package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type userService struct {
	userRepo sa.UserRepository
}

// Store ...
func (u userService) Store(ctx context.Context, user sa.User) (sa.User, error) {
	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return sa.User{}, err
	}

	user.Password = string(bytesPassword)

	return u.userRepo.Store(ctx, user)
}

// NewUserService .
func NewUserService(userRepo sa.UserRepository) sa.UserService {
	return userService{userRepo: userRepo}
}
