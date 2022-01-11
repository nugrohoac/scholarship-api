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
		Name:    param.Name,
		Type:    param.Type,
		Email:   param.Email,
		PhoneNo: param.PhoneNo,
		Photo: sa.Image{
			URL:    param.Photo.URL,
			Width:  param.Photo.Width,
			Height: param.Photo.Height,
		},
		CountryID:       param.CountryID,
		PostalCode:      param.PostalCode,
		Address:         param.Address,
		BankID:          param.BankID,
		BankAccountNo:   param.BankAccountNo,
		BankAccountName: param.BankAccountName,
	}

	if param.Photo.Mime != nil {
		user.Photo.Mime = *param.Photo.Mime
	}

	if param.Photo.Caption != nil {
		user.Photo.Mime = *param.Photo.Caption
	}

	if param.CompanyName != nil {
		user.Photo.Mime = *param.CompanyName
	}

	if param.Gender != nil {
		user.Gender = *param.Gender
	}

	if param.Ethnic != nil {
		user.Gender = *param.Ethnic
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
