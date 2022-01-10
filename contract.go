package scholarship_api

import "context"

// BankRepository .
type BankRepository interface {
	Fetch(ctx context.Context, filter BankFilter) ([]Bank, string, error)
}

// BankService .
type BankService interface {
	Fetch(ctx context.Context, filter BankFilter) (BankFeed, error)
}

// CountryRepository .
type CountryRepository interface {
	Fetch(ctx context.Context, filter CountryFilter) ([]Country, string, error)
}

// CountryService .
type CountryService interface {
	Fetch(ctx context.Context, filter CountryFilter) (CountryFeed, error)
}
