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
	query.ScholarshipQuery
	query.DegreeQuery
	mutation.UserMutation
	mutation.ScholarshipMutation
	mutation.PaymentMutation
}

// NewRootGraphql ...
func NewRootGraphql(
	bankQuery query.BankQuery,
	countryQuery query.CountryQuery,
	userQuery query.UserQuery,
	scholarshipQuery query.ScholarshipQuery,
	degreeQuery query.DegreeQuery,
	userMutation mutation.UserMutation,
	scholarshipMutation mutation.ScholarshipMutation,
	paymentMutation mutation.PaymentMutation,

) *RootGraphql {
	return &RootGraphql{
		bankQuery,
		countryQuery,
		userQuery,
		scholarshipQuery,
		degreeQuery,
		userMutation,
		scholarshipMutation,
		paymentMutation,
	}
}
