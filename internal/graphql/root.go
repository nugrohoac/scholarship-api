package graphql

import (
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
)

// RootGraphql .
type RootGraphql struct {
	query.BankQuery
	query.CountryQuery
	query.UserQuery
	mutation.UserMutation
}

// NewRootGraphql ...
func NewRootGraphql(
	bankQuery query.BankQuery,
	countryQuery query.CountryQuery,
	userQuery query.UserQuery,
	userMutation mutation.UserMutation,
) *RootGraphql {
	return &RootGraphql{
		bankQuery,
		countryQuery,
		userQuery,
		userMutation,
	}
}
