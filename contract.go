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

// UserRepository ....
type UserRepository interface {
	Store(ctx context.Context, user User) (User, error)
}

// UserService ....
type UserService interface {
	Store(ctx context.Context, user User) (User, error)
}
