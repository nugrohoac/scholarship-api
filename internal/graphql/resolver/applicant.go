package resolver

import (
	"time"

	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// ApplicantResolver .
type ApplicantResolver struct {
	Applicant entity.Applicant
}

// ID .
func (a ApplicantResolver) ID() *int32 {
	ID := int32(a.Applicant.ID)

	return &ID
}

// ScholarshipID .
func (a ApplicantResolver) ScholarshipID() *int32 {
	scholarshipID := int32(a.Applicant.ScholarshipID)

	return &scholarshipID
}

// UserID .
func (a ApplicantResolver) UserID() *int32 {
	userID := int32(a.Applicant.UserID)

	return &userID
}

// Status .
func (a ApplicantResolver) Status() *int32 {
	return &a.Applicant.Status
}

// ApplyDate .
func (a ApplicantResolver) ApplyDate() *string {
	applyDate := a.Applicant.ApplyDate.Format(time.RFC3339)
	return &applyDate
}

// User .
func (a ApplicantResolver) User() *UserResolver {
	return &UserResolver{User: a.Applicant.User}
}

// Essay .
func (a ApplicantResolver) Essay() *string {
	return &a.Applicant.Essay
}

// RecommendationLetter .
func (a ApplicantResolver) RecommendationLetter() *ImageResolver {
	return &ImageResolver{Image: a.Applicant.RecommendationLetter}
}

// Scores .
func (a ApplicantResolver) Scores() *[]*ScoreResolver {
	scores := make([]*ScoreResolver, 0)

	for _, score := range a.Applicant.Scores {
		score := score

		scores = append(scores, &ScoreResolver{Score: score})
	}

	return &scores
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

		applicants = append(applicants, &ApplicantResolver{Applicant: applicant})
	}

	return &applicants
}

// ScoreResolver .
type ScoreResolver struct {
	Score entity.ApplicantScore
}

// Name .
func (s ScoreResolver) Name() *string {
	return &s.Score.Name
}

// Value .
func (s ScoreResolver) Value() *int32 {
	return &s.Score.Value
}
