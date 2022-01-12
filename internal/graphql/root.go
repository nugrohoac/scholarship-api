package graphql

import (
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
)

// RootGraphql .
type RootGraphql struct {
	query.BankQuery
	query.CountryQuery
	mutation.UserMutation
}

// NewRootGraphql ...
func NewRootGraphql(bankQuery query.BankQuery, countryQuery query.CountryQuery, userMutation mutation.UserMutation) *RootGraphql {
	return &RootGraphql{
		bankQuery,
		countryQuery,
		userMutation,
	}
}
