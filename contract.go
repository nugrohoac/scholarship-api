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
	Fetch(ctx context.Context, filter UserFilter) ([]User, string, error)
	Login(ctx context.Context, email string) (User, error)
	UpdateByID(ctx context.Context, ID int64, user User) (User, error)
	SetStatus(ctx context.Context, ID int64, status int) error
	ResetPassword(ctx context.Context, email, password string) error
}

// UserService ....
type UserService interface {
	Store(ctx context.Context, user User) (User, error)
	Login(ctx context.Context, email, password string) (LoginResponse, error)
	UpdateByID(ctx context.Context, ID int64, user User) (User, error)
	ActivateStatus(ctx context.Context, token string) (User, error)
	ResendEmailVerification(ctx context.Context, email string) (string, error)
	ResetPassword(ctx context.Context, password string) (string, error)
	ForgotPassword(ctx context.Context, email string) (string, error)
}

// CountryRepository .
type CountryRepository interface {
	Fetch(ctx context.Context, filter CountryFilter) ([]Country, string, error)
}

// CountryService .
type CountryService interface {
	Fetch(ctx context.Context, filter CountryFilter) (CountryFeed, error)
}

// JwtHash ...
type JwtHash interface {
	Encode(user User) (string, error)
	Decode(tokenString string, claim *Claim) error
}

// EmailRepository ...
type EmailRepository interface {
	SendActivateUser(ctx context.Context, email, token string) error
	SendForgotPassword(ctx context.Context, email, token string) error
}
