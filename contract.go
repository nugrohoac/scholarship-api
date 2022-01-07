package scholarship_api

import "context"

type BankRepository interface {
	Fetch(ctx context.Context, filter BankFilter) ([]Bank, string, error)
}

type BankService interface {
	Fetch(ctx context.Context, filter BankFilter) (BankFeed, error)
}
