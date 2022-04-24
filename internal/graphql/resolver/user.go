package resolver

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"time"
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

// Country .
func (u UserResolver) Country() *CountryResolver {
	return &CountryResolver{Country: u.User.Country}
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
	if u.User.Gender == "" {
		return nil
	}

	return &u.User.Gender
}

// Rating .
func (u UserResolver) Rating() *float64 {
	return &u.User.Rating
}

// EthnicID ...
func (u UserResolver) EthnicID() *int32 {
	return &u.User.EthnicID
}

// BirthDate .
func (u UserResolver) BirthDate() *string {
	bd := u.User.BirthDate.Format(time.RFC3339)
	return &bd
}

// BirthPlace .
func (u UserResolver) BirthPlace() *string {
	return &u.User.BirthPlace
}

// Ethnic .
func (u UserResolver) Ethnic() *EthnicResolver {
	return &EthnicResolver{Ethnic: u.User.Ethnic}
}

// BankID ...
func (u UserResolver) BankID() *int32 {
	bankID := u.User.BankID
	return &bankID
}

// Bank .
func (u UserResolver) Bank() *BankResolver {
	return &BankResolver{Bank: u.User.Bank}
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

// CreatedAt ...
func (u UserResolver) CreatedAt() *string {
	time := u.User.CreatedAt.Format("2006-01-02 15:04:05")
	return &time
}

// CareerGoal .
func (u UserResolver) CareerGoal() *string {
	return &u.User.CareerGoal
}

// UserSchools .
func (u UserResolver) UserSchools() *[]*UserSchoolResolver {
	us := make([]*UserSchoolResolver, 0)

	for _, _us := range u.User.UserSchools {
		_us := _us
		us = append(us, &UserSchoolResolver{UserSchool: _us})
	}

	return &us
}

// UserDocuments .
func (u UserResolver) UserDocuments() *[]*UserDocumentResolver {
	uds := make([]*UserDocumentResolver, 0)

	for _, ud := range u.User.UserDocuments {
		ud := ud
		uds = append(uds, &UserDocumentResolver{userDocument: ud})
	}

	return &uds
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
