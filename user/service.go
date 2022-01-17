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

// UpdateByID ...
// update for sponsor currently
func (u userService) UpdateByID(ctx context.Context, ID int64, user sa.User) (sa.User, error) {
	users, _, err := u.userRepo.Fetch(ctx, sa.UserFilter{IDs: []int64{ID}})
	if err != nil {
		return sa.User{}, err
	}

	if len(users) == 0 {
		return sa.User{}, sa.ErrNotFound{Message: "user not found"}
	}

	userOnCtx, err := sa.GetUserOnContext(ctx)
	if err != nil {
		return sa.User{}, err
	}

	if userOnCtx.Email != users[0].Email {
		return sa.User{}, sa.ErrUnAuthorize{Message: "user is not sync"}
	}

	return u.userRepo.UpdateByID(ctx, ID, user)
}

// ActivateStatus ...
func (u userService) ActivateStatus(ctx context.Context, ID int64) (string, error) {
	user, err := sa.GetUserOnContext(ctx)
	if err != nil {
		return "", err
	}

	users, _, err := u.userRepo.Fetch(ctx, sa.UserFilter{IDs: []int64{ID}})
	if err != nil {
		return "", err
	}

	if user.Email != users[0].Email {
		return "", sa.ErrUnAuthorize{Message: "user is not sync"}
	}

	if err = u.userRepo.SetStatus(ctx, ID, 1); err != nil {
		return "", err
	}

	return "success", nil
}

// NewUserService .
func NewUserService(userRepo sa.UserRepository, jwtHash sa.JwtHash) sa.UserService {
	return userService{userRepo: userRepo, jwtHash: jwtHash}
}
