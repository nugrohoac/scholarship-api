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
