package graphql

import (
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
)

// RootGraphql .
type RootGraphql struct {
	query.BankQuery
	mutation.UserMutation
}

// NewRootGraphql ...
func NewRootGraphql(bankQuery query.BankQuery, userMutation mutation.UserMutation) *RootGraphql {
	return &RootGraphql{bankQuery, userMutation}
}
