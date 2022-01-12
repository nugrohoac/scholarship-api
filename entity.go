package scholarship_api

import "time"

// Bank ...
type Bank struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

// InputBankFilter ...
type InputBankFilter struct {
	Limit  *int32
	Cursor *string
	Name   *string
}

// BankFilter ...
type BankFilter struct {
	Limit  int
	Cursor string
	Name   string
}

// BankFeed ...
type BankFeed struct {
	Cursor string `json:"cursor"`
	Banks  []Bank `json:"banks"`
}

// Image ...
type Image struct {
	URL     string `json:"url"`
	Width   int32  `json:"width"`
	Height  int32  `json:"height"`
	Mime    string `json:"mime"`
	Caption string `json:"caption"`
}

// InputImage ...
type InputImage struct {
	URL     string
	Width   int32
	Height  int32
	Mime    *string
	Caption *string
}

// User ....
type User struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Type            string    `json:"type"`
	Email           string    `json:"email"`
	PhoneNo         string    `json:"phone_no"`
	Photo           Image     `json:"photo"`
	CompanyName     string    `json:"company_name"`
	Password        string    `json:"-"`
	Status          int       `json:"status"`
	CountryID       int32     `json:"country_id"`
	PostalCode      string    `json:"postal_code"`
	Address         string    `json:"address"`
	Gender          string    `json:"gender"`
	Ethnic          string    `json:"ethnic"`
	BankID          int32     `json:"bank_id"`
	BankAccountNo   string    `json:"bank_account_no"`
	BankAccountName string    `json:"bank_account_name"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"-"`
}

// InputRegisterUser .
type InputRegisterUser struct {
	Name            string
	Type            string
	Email           string
	PhoneNo         string
	Photo           InputImage
	CompanyName     *string
	CountryID       int32
	PostalCode      string
	Address         string
	Gender          *string
	Ethnic          *string
	BankID          int32
	BankAccountNo   string
	BankAccountName string
}

// Country ...
type Country struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// CountryFeed ...
type CountryFeed struct {
	Cursor    string
	Countries []Country `json:"countries"`
}

// CountryFilter ...
type CountryFilter struct {
	Limit  int
	Cursor string
	Name   string
}

// InputCountryFilter ...
type InputCountryFilter struct {
	Limit  *int32
	Cursor *string
	Name   *string
}
