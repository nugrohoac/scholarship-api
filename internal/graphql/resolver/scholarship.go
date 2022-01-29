package resolver

import (
	"time"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

// ScholarshipResolver ...
type ScholarshipResolver struct {
	Scholarship sa.Scholarship
}

// ID ...
func (s ScholarshipResolver) ID() *int32 {
	ID := int32(s.Scholarship.ID)
	return &ID
}

// SponsorID ...
func (s ScholarshipResolver) SponsorID() *int32 {
	ID := int32(s.Scholarship.SponsorID)
	return &ID
}

// Sponsor ...
func (s ScholarshipResolver) Sponsor() *UserResolver {
	return &UserResolver{User: s.Scholarship.Sponsor}
}

// Name ...
func (s ScholarshipResolver) Name() *string {
	return &s.Scholarship.Name
}

// Amount ...
func (s ScholarshipResolver) Amount() *int32 {
	amount := int32(s.Scholarship.Amount)
	return &amount
}

// Status ...
func (s ScholarshipResolver) Status() *int32 {
	status := int32(s.Scholarship.Status)
	return &status
}

// Image ...
func (s ScholarshipResolver) Image() *ImageResolver {
	return &ImageResolver{Image: s.Scholarship.Image}
}

// Awardee ...
func (s ScholarshipResolver) Awardee() *int32 {
	awardee := int32(s.Scholarship.Awardee)
	return &awardee
}

// CurrentApplicant ...
func (s ScholarshipResolver) CurrentApplicant() *int32 {
	ca := int32(s.Scholarship.CurrentApplicant)
	return &ca
}

// Deadline ...
func (s ScholarshipResolver) Deadline() *string {
	deadline := s.Scholarship.Deadline.Format(time.RFC3339)
	return &deadline
}

// EligibilityDescription ...
func (s ScholarshipResolver) EligibilityDescription() *string {
	return &s.Scholarship.EligibilityDescription
}

// SubsidyDescription ...
func (s ScholarshipResolver) SubsidyDescription() *string {
	return &s.Scholarship.SubsidyDescription
}

// FundingStart ...
func (s ScholarshipResolver) FundingStart() *string {
	fs := s.Scholarship.FundingStart.Format(time.RFC3339)
	return &fs
}

// FundingEnd ...
func (s ScholarshipResolver) FundingEnd() *string {
	fs := s.Scholarship.FundingEnd.Format(time.RFC3339)
	return &fs
}

// Requirements ...
func (s ScholarshipResolver) Requirements() *[]*RequirementResolver {
	rs := make([]*RequirementResolver, 0)

	for _, r := range s.Scholarship.Requirements {
		r := r
		rs = append(rs, &RequirementResolver{Requirement: r})
	}

	return &rs
}

// ScholarshipFeedResolver ...
type ScholarshipFeedResolver struct {
	ScholarshipFeed sa.ScholarshipFeed
}

// Cursor ...
func (s ScholarshipFeedResolver) Cursor() *string {
	return &s.ScholarshipFeed.Cursor
}

// Scholarships ...
func (s ScholarshipFeedResolver) Scholarships() *[]*ScholarshipResolver {
	scholarshipResolvers := make([]*ScholarshipResolver, 0)

	for _, scholarship := range s.ScholarshipFeed.Scholarships {
		scholarship := scholarship

		scholarshipResolvers = append(scholarshipResolvers, &ScholarshipResolver{Scholarship: scholarship})
	}

	return &scholarshipResolvers
}