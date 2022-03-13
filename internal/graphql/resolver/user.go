package resolver

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// UserResolver ...
type UserResolver struct {
	User entity.User
}

// ID ...
func (u UserResolver) ID() *int32 {
	ID := int32(u.User.ID)

	return &ID
}

// Name ...
func (u UserResolver) Name() *string {
	return &u.User.Name
}

// Type ...
func (u UserResolver) Type() *string {
	if u.User.Type == "" {
		return nil
	}

	return &u.User.Type
}

// Email ...
func (u UserResolver) Email() *string {
	return &u.User.Email
}

// PhoneNo ...
func (u UserResolver) PhoneNo() *string {
	return &u.User.PhoneNo
}

// Photo ...
func (u UserResolver) Photo() *ImageResolver {
	return &ImageResolver{Image: u.User.Photo}
}

// CompanyName ...
func (u UserResolver) CompanyName() *string {
	return &u.User.CompanyName
}

// Status ...
func (u UserResolver) Status() *int32 {
	status := int32(u.User.Status)
	return &status
}

// CountryID ...
func (u UserResolver) CountryID() *int32 {
	countryID := u.User.CountryID
	return &countryID
}

// PostalCode ...
func (u UserResolver) PostalCode() *string {
	return &u.User.PostalCode
}

// Address ...
func (u UserResolver) Address() *string {
	return &u.User.Address
}

// Gender ...
func (u UserResolver) Gender() *string {
	return &u.User.Gender
}

// Ethnic ...
func (u UserResolver) Ethnic() *string {
	return &u.User.Ethnic
}

// BankID ...
func (u UserResolver) BankID() *int32 {
	bankID := u.User.BankID
	return &bankID
}

// BankAccountNo ...
func (u UserResolver) BankAccountNo() *string {
	return &u.User.BankAccountNo
}

// BankAccountName ...
func (u UserResolver) BankAccountName() *string {
	return &u.User.BankAccountName
}

// CardIdentities ...
func (u UserResolver) CardIdentities() *[]*CardIdentityResolver {
	cards := make([]*CardIdentityResolver, 0)

	for _, c := range u.User.CardIdentities {
		c := c
		cards = append(cards, &CardIdentityResolver{CardIdentity: c})
	}

	return &cards
}

// LoginResponseResolver ...
type LoginResponseResolver struct {
	LoginResponse entity.LoginResponse
}

// Token ...
func (l LoginResponseResolver) Token() *string {
	return &l.LoginResponse.Token
}

// User ...
func (l LoginResponseResolver) User() *UserResolver {
	return &UserResolver{User: l.LoginResponse.User}
}