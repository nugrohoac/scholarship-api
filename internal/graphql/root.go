package graphql

import "github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"

// RootGraphql .
type RootGraphql struct {
	query.BankQuery
}

// NewRootGraphql ...
func NewRootGraphql(bankQuery query.BankQuery) *RootGraphql {
	return &RootGraphql{bankQuery}
}
