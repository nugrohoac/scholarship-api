package graphql

import (
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query/backoffice"
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
	query.EthnicQuery
	mutation.UserMutation
	mutation.ScholarshipMutation
	mutation.PaymentMutation
	mutation.SchoolMutation
	backoffice.SponsorQuery
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
	ethnicQuery query.EthnicQuery,
	userMutation mutation.UserMutation,
	scholarshipMutation mutation.ScholarshipMutation,
	paymentMutation mutation.PaymentMutation,
	schoolMutation mutation.SchoolMutation,
	sponsorQuery backoffice.SponsorQuery,

) *RootGraphql {
	return &RootGraphql{
		bankQuery,
		countryQuery,
		userQuery,
		scholarshipQuery,
		degreeQuery,
		majorQuery,
		schoolQuery,
		ethnicQuery,
		userMutation,
		scholarshipMutation,
		paymentMutation,
		schoolMutation,
		sponsorQuery,
	}
}
