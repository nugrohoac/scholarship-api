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

// UpdateUser ...
func (u UserMutation) UpdateUser(ctx context.Context, param sa.InputUpdateUser) (*resolver.UserResolver, error) {
	cardIdentities := make([]sa.CardIdentity, 0)
	for _, c := range param.CardIdentities {
		cardIdentities = append(cardIdentities, sa.CardIdentity{
			Type: c.Type,
			No:   c.No,
			Image: sa.Image{
				URL:    c.Image.URL,
				Width:  c.Image.Width,
				Height: c.Image.Height,
			},
			UserID: int64(param.ID),
		})
	}

	user := sa.User{
		ID:   int64(param.ID),
		Name: param.Name,
		Photo: sa.Image{
			URL:    param.Photo.URL,
			Width:  param.Photo.Width,
			Height: param.Photo.Height,
		},
		CompanyName:     param.CompanyName,
		CountryID:       param.CountryID,
		PostalCode:      param.PostalCode,
		Address:         param.Address,
		CardIdentities:  cardIdentities,
		BankID:          param.BankID,
		BankAccountNo:   param.BankAccountNo,
		BankAccountName: param.BankAccountName,
	}

	userUpdated, err := u.userService.UpdateByID(ctx, user.ID, user)
	if err != nil {
		return nil, err
	}

	return &resolver.UserResolver{User: userUpdated}, nil
}

// NewUserMutation ...
func NewUserMutation(userService sa.UserService) UserMutation {
	return UserMutation{userService: userService}
}
