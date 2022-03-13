package resolver

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// BankResolver .
type BankResolver struct {
	Bank entity.Bank
}

// ID ...
func (b BankResolver) ID() *int32 {
	ID := int32(b.Bank.ID)
	return &ID
}

// Name ...
func (b BankResolver) Name() *string {
	return &b.Bank.Name
}

// Code ...
func (b BankResolver) Code() *string {
	return &b.Bank.Code
}

// BankFeedResolver .
type BankFeedResolver struct {
	BankFeed entity.BankFeed
}

// Cursor .
func (b BankFeedResolver) Cursor() *string {
	return &b.BankFeed.Cursor
}

// Banks .
func (b BankFeedResolver) Banks() *[]*BankResolver {
	banksResolver := make([]*BankResolver, 0)

	for _, bank := range b.BankFeed.Banks {
		bank := bank
		banksResolver = append(banksResolver, &BankResolver{
			Bank: bank,
		})
	}

	return &banksResolver
}
