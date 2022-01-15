package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type userService struct {
	userRepo sa.UserRepository
	jwtHash  sa.JwtHash
}

// Store ...
func (u userService) Store(ctx context.Context, user sa.User) (sa.User, error) {
	users, _, err := u.userRepo.Fetch(ctx, sa.UserFilter{Email: user.Email})
	if err != nil {
		return sa.User{}, err
	}

	if len(users) > 0 {
		return sa.User{}, sa.ErrorDuplicate{Message: "email already exist"}
	}

	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return sa.User{}, err
	}

	user.Password = string(bytesPassword)

	return u.userRepo.Store(ctx, user)
}

// Login ....
func (u userService) Login(ctx context.Context, email, password string) (sa.LoginResponse, error) {
	user, err := u.userRepo.Login(ctx, email)
	if err != nil {
		return sa.LoginResponse{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return sa.LoginResponse{}, err
	}

	token, err := u.jwtHash.Encode(user)
	if err != nil {
		return sa.LoginResponse{}, err
	}

	return sa.LoginResponse{Token: token, User: user}, nil
}

// NewUserService .
func NewUserService(userRepo sa.UserRepository, jwtHash sa.JwtHash) sa.UserService {
	return userService{userRepo: userRepo, jwtHash: jwtHash}
}
