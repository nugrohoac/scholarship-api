package graphql

import "github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"

// RootGraphql .
type RootGraphql struct {
	query.BankQuery
	query.CountryQuery
}

// NewRootGraphql ...
func NewRootGraphql(bankQuery query.BankQuery, countryQuery query.CountryQuery) *RootGraphql {
	return &RootGraphql{
		bankQuery,
		countryQuery,
	}
}
