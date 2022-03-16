package resolver

import (
	"time"

	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// ApplicantResolver .
type ApplicantResolver struct {
	applicant entity.Applicant
}

// ID .
func (a ApplicantResolver) ID() *int32 {
	ID := int32(a.applicant.ID)

	return &ID
}

// ScholarshipID .
func (a ApplicantResolver) ScholarshipID() *int32 {
	scholarshipID := int32(a.applicant.ScholarshipID)

	return &scholarshipID
}

// UserID .
func (a ApplicantResolver) UserID() *int32 {
	userID := int32(a.applicant.UserID)

	return &userID
}

// Status .
func (a ApplicantResolver) Status() *int32 {
	return &a.applicant.Status
}

// ApplyDate .
func (a ApplicantResolver) ApplyDate() *string {
	applyDate := a.applicant.ApplyDate.Format(time.RFC3339)
	return &applyDate
}

// User .
func (a ApplicantResolver) User() *UserResolver {
	return &UserResolver{User: a.applicant.User}
}

// ApplicantFeedResolver .
type ApplicantFeedResolver struct {
	ApplicantFeed entity.ApplicantFeed
}

// Cursor .
func (a ApplicantFeedResolver) Cursor() *string {
	return &a.ApplicantFeed.Cursor
}

// Applicants .
func (a ApplicantFeedResolver) Applicants() *[]*ApplicantResolver {
	applicants := make([]*ApplicantResolver, 0)

	for _, applicant := range a.ApplicantFeed.Applicants {
		applicant := applicant

		applicants = append(applicants, &ApplicantResolver{applicant: applicant})
	}

	return &applicants
}
