package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

// 0 status un verification
// 1 status verify but un complete profile
// 2 status verify and complete profile

type userService struct {
	userRepo  sa.UserRepository
	jwtHash   sa.JwtHash
	emailRepo sa.EmailRepository
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

	user, err = u.userRepo.Store(ctx, user)
	if err != nil {
		return sa.User{}, err
	}

	if err = u.sendEmail(user); err != nil {
		return sa.User{}, err
	}

	return user, nil
}

func (u userService) sendEmail(user sa.User) error {
	token, err := u.jwtHash.Encode(user)
	if err != nil {
		return err
	}

	if err = u.emailRepo.SendActivateUser(context.Background(), user.Email, token); err != nil {
		return err
	}

	return nil
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

	// 2 = active and profile is complete
	user.Status = 2
	user, err = u.userRepo.UpdateByID(ctx, ID, user)
	if err != nil {
		return sa.User{}, err
	}

	user.Email = userOnCtx.Email
	user.Type = userOnCtx.Type
	user.Name = userOnCtx.Name

	return user, nil
}

// ActivateStatus ...
func (u userService) ActivateStatus(ctx context.Context, token string) (sa.User, error) {
	var c sa.Claim
	if err := u.jwtHash.Decode(token, &c); err != nil {
		return sa.User{}, err
	}

	users, _, err := u.userRepo.Fetch(ctx, sa.UserFilter{Email: c.Email})
	if err != nil {
		return sa.User{}, err
	}

	if len(users) == 0 {
		return sa.User{}, sa.ErrNotFound{Message: "user not found"}
	}

	user := users[0]

	if user.Status == 1 {
		return sa.User{}, sa.ErrNotAllowed{Message: "user has ben activated"}
	}

	if err = u.userRepo.SetStatus(ctx, user.ID, 1); err != nil {
		return sa.User{}, err
	}

	user.Status = 1

	return user, nil
}

// ResendEmailVerification ...
func (u userService) ResendEmailVerification(ctx context.Context, email string) (string, error) {
	users, _, err := u.userRepo.Fetch(ctx, sa.UserFilter{Email: email})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", sa.ErrNotFound{Message: "user is not found"}
	}

	token, err := u.jwtHash.Encode(users[0])
	if err != nil {
		return "", err
	}

	if err = u.emailRepo.SendActivateUser(ctx, email, token); err != nil {
		return "", err
	}

	return "success", nil
}

// ForgotPassword ...
func (u userService) ForgotPassword(ctx context.Context, email string) (string, error) {
	users, _, err := u.userRepo.Fetch(ctx, sa.UserFilter{Email: email})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", sa.ErrNotFound{Message: "user is not found"}
	}

	if users[0].Status == 0 {
		return "", sa.ErrNotAllowed{Message: "account is inactivate, please choose resend email activation"}
	}

	token, err := u.jwtHash.Encode(users[0])
	if err != nil {
		return "", err
	}

	if err = u.emailRepo.SendForgotPassword(ctx, email, token); err != nil {
		return "", err
	}

	return "success", nil
}

// ResetPassword ...
func (u userService) ResetPassword(ctx context.Context, password string) (sa.User, error) {
	user, err := sa.GetUserOnContext(ctx)
	if err != nil {
		return sa.User{}, err
	}

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return sa.User{}, err
	}

	if err = u.userRepo.ResetPassword(ctx, user.Email, string(bytePassword)); err != nil {
		return sa.User{}, err
	}

	return user, nil
}

// SetupEducation .
func (u userService) SetupEducation(ctx context.Context, user sa.User) (sa.User, error) {
	userCtx, err := sa.GetUserOnContext(ctx)
	if err != nil {
		return sa.User{}, err
	}

	if userCtx.ID != user.ID {
		return sa.User{}, sa.ErrUnAuthorize{Message: "user is not match"}
	}

	if userCtx.Type != sa.Student {
		return sa.User{}, sa.ErrNotAllowed{Message: "user type is not student"}
	}

	userLogin, err := u.userRepo.Login(ctx, userCtx.Email)
	if err != nil {
		return sa.User{}, err
	}

	// check status, should 2
	if userLogin.Status != 2 {
		return sa.User{}, sa.ErrNotAllowed{Message: "user status is not complete profile"}
	}

	user.Name = userCtx.Name
	user.Email = userCtx.Email
	user.Type = userCtx.Type
	user.Status = userLogin.Status

	// look at readme.md to get more status
	user.Status = 3

	return u.userRepo.SetupEducation(ctx, user)
}

// NewUserService .
func NewUserService(userRepo sa.UserRepository, jwtHash sa.JwtHash, emailRepo sa.EmailRepository) sa.UserService {
	return userService{userRepo: userRepo, jwtHash: jwtHash, emailRepo: emailRepo}
}
