package scholarship_api

import "time"

type Bank struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

// BankFilter ...
type BankFilter struct {
	Limit  int
	Cursor string
}

// BankFeed ...
type BankFeed struct {
	Cursor string `json:"cursor"`
	Banks  []Bank `json:"banks"`
}
