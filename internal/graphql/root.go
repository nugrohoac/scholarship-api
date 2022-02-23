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
	query.MajorQuery
	query.SchoolQuery
	mutation.UserMutation
	mutation.ScholarshipMutation
	mutation.PaymentMutation
	mutation.SchoolMutation
}

// NewRootGraphql ...
func NewRootGraphql(
	bankQuery query.BankQuery,
	countryQuery query.CountryQuery,
	userQuery query.UserQuery,
	scholarshipQuery query.ScholarshipQuery,
	degreeQuery query.DegreeQuery,
	majorQuery query.MajorQuery,
	schoolQuery query.SchoolQuery,
	userMutation mutation.UserMutation,
	scholarshipMutation mutation.ScholarshipMutation,
	paymentMutation mutation.PaymentMutation,
	schoolMutation mutation.SchoolMutation,

) *RootGraphql {
	return &RootGraphql{
		bankQuery,
		countryQuery,
		userQuery,
		scholarshipQuery,
		degreeQuery,
		majorQuery,
		schoolQuery,
		userMutation,
		scholarshipMutation,
		paymentMutation,
		schoolMutation,
	}
}
