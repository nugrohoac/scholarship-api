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
	query.ApplicantQuery
	query.EmailQuery
	backoffice.SponsorQuery
	backoffice.StudentQuery
	mutation.UserMutation
	mutation.ScholarshipMutation
	mutation.PaymentMutation
	mutation.SchoolMutation
	mutation.AssessmentMutation
	mutation.ApplicantMutation
	mutation.EmailMutation
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
	applicantQuery query.ApplicantQuery,
	emailQuery query.EmailQuery,
	sponsorQuery backoffice.SponsorQuery,
	studentQuery backoffice.StudentQuery,
	userMutation mutation.UserMutation,
	scholarshipMutation mutation.ScholarshipMutation,
	paymentMutation mutation.PaymentMutation,
	schoolMutation mutation.SchoolMutation,
	assessmentMutation mutation.AssessmentMutation,
	applicantMutation mutation.ApplicantMutation,
	emailMutation mutation.EmailMutation,

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
		applicantQuery,
		emailQuery,
		sponsorQuery,
		studentQuery,
		userMutation,
		scholarshipMutation,
		paymentMutation,
		schoolMutation,
		assessmentMutation,
		applicantMutation,
		emailMutation,
	}
}
