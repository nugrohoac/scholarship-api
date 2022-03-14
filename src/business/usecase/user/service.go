package user

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"
	"golang.org/x/crypto/bcrypt"
)

// 0 status un verification
// 1 status verify but un complete profile
// 2 status verify and complete profile

type userService struct {
	userRepo  business.UserRepository
	jwtHash   business.JwtHash
	emailRepo business.EmailRepository
}

// Store ...
func (u userService) Store(ctx context.Context, user entity.User) (entity.User, error) {
	users, _, err := u.userRepo.Fetch(ctx, entity.UserFilter{Email: user.Email})
	if err != nil {
		return entity.User{}, err
	}

	if len(users) > 0 {
		return entity.User{}, errors.ErrorDuplicate{Message: "email already exist"}
	}

	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return entity.User{}, err
	}

	user.Password = string(bytesPassword)

	user, err = u.userRepo.Store(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	if err = u.sendEmail(user); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u userService) sendEmail(user entity.User) error {
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
func (u userService) Login(ctx context.Context, email, password string) (entity.LoginResponse, error) {
	user, err := u.userRepo.Login(ctx, email)
	if err != nil {
		return entity.LoginResponse{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return entity.LoginResponse{}, err
	}

	token, err := u.jwtHash.Encode(user)
	if err != nil {
		return entity.LoginResponse{}, err
	}

	return entity.LoginResponse{Token: token, User: user}, nil
}

// UpdateByID ...
// update for sponsor currently
func (u userService) UpdateByID(ctx context.Context, ID int64, user entity.User) (entity.User, error) {
	users, _, err := u.userRepo.Fetch(ctx, entity.UserFilter{IDs: []int64{ID}})
	if err != nil {
		return entity.User{}, err
	}

	if len(users) == 0 {
		return entity.User{}, errors.ErrNotFound{Message: "user not found"}
	}

	userOnCtx, err := common.GetUserOnContext(ctx)
	if err != nil {
		return entity.User{}, err
	}

	if userOnCtx.Email != users[0].Email {
		return entity.User{}, errors.ErrUnAuthorize{Message: "user is not sync"}
	}

	// 2 = active and profile is complete
	user.Status = 2

	if users[0].Status > 2 {
		// 2 = active and profile is complete
		user.Status = users[0].Status
	}

	user, err = u.userRepo.UpdateByID(ctx, ID, user)
	if err != nil {
		return entity.User{}, err
	}

	user.Email = userOnCtx.Email
	user.Type = userOnCtx.Type
	user.Name = userOnCtx.Name

	return user, nil
}

// ActivateStatus ...
func (u userService) ActivateStatus(ctx context.Context, token string) (entity.User, error) {
	var c entity.Claim
	if err := u.jwtHash.Decode(token, &c); err != nil {
		return entity.User{}, err
	}

	users, _, err := u.userRepo.Fetch(ctx, entity.UserFilter{Email: c.Email})
	if err != nil {
		return entity.User{}, err
	}

	if len(users) == 0 {
		return entity.User{}, errors.ErrNotFound{Message: "user not found"}
	}

	user := users[0]

	if user.Status == 1 {
		return entity.User{}, errors.ErrNotAllowed{Message: "user has ben activated"}
	}

	if err = u.userRepo.SetStatus(ctx, user.ID, 1); err != nil {
		return entity.User{}, err
	}

	user.Status = 1

	return user, nil
}

// ResendEmailVerification ...
func (u userService) ResendEmailVerification(ctx context.Context, email string) (string, error) {
	users, _, err := u.userRepo.Fetch(ctx, entity.UserFilter{Email: email})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", errors.ErrNotFound{Message: "user is not found"}
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
	users, _, err := u.userRepo.Fetch(ctx, entity.UserFilter{Email: email})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", errors.ErrNotFound{Message: "user is not found"}
	}

	if users[0].Status == 0 {
		return "", errors.ErrNotAllowed{Message: "account is inactivate, please choose resend email activation"}
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
func (u userService) ResetPassword(ctx context.Context, password string) (entity.User, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return entity.User{}, err
	}

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return entity.User{}, err
	}

	if err = u.userRepo.ResetPassword(ctx, user.Email, string(bytePassword)); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// SetupEducation .
func (u userService) SetupEducation(ctx context.Context, user entity.User) (entity.User, error) {
	userCtx, err := common.GetUserOnContext(ctx)
	if err != nil {
		return entity.User{}, err
	}

	if userCtx.ID != user.ID {
		return entity.User{}, errors.ErrUnAuthorize{Message: "user is not match"}
	}

	if userCtx.Type != entity.Student {
		return entity.User{}, errors.ErrNotAllowed{Message: "user type is not student"}
	}

	userLogin, err := u.userRepo.Login(ctx, userCtx.Email)
	if err != nil {
		return entity.User{}, err
	}

	// check status, should 2
	if userLogin.Status != 2 {
		return entity.User{}, errors.ErrNotAllowed{Message: "user status is not complete profile"}
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
func NewUserService(userRepo business.UserRepository, jwtHash business.JwtHash, emailRepo business.EmailRepository) business.UserService {
	return userService{userRepo: userRepo, jwtHash: jwtHash, emailRepo: emailRepo}
}
