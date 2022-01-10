package graphql

import "github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"

// RootGraphql .
type RootGraphql struct {
	BankQuery query.BankQuery
}

// NewRootGraphql ...
func NewRootGraphql(bankQuery query.BankQuery) RootGraphql {
	return RootGraphql{BankQuery: bankQuery}
}
