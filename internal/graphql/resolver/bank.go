package resolver

import sa "github.com/Nusantara-Muda/scholarship-api"

// BankResolver .
type BankResolver struct {
	Bank sa.Bank
}

func (b BankResolver) id() int64 {
	return b.Bank.ID
}

func (b BankResolver) name() string {
	return b.Bank.Name
}

func (b BankResolver) code() string {
	return b.Bank.Code
}

// BankFeedResolver .
type BankFeedResolver struct {
	BankFeed sa.BankFeed
}

func (b BankFeedResolver) cursor() string {
	return b.BankFeed.Cursor
}

func (b BankFeedResolver) banks() []BankResolver {
	banksResolver := make([]BankResolver, 0)

	for _, bank := range b.BankFeed.Banks {
		banksResolver = append(banksResolver, BankResolver{
			Bank: bank,
		})
	}

	return banksResolver
}
