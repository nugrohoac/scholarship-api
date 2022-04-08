package mutation

import (
	"context"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// AssessmentMutation .
type AssessmentMutation struct {
	assessmentService business.AssessmentService
}

// SubmitAssessment .
func (a AssessmentMutation) SubmitAssessment(ctx context.Context, param entity.InputAssessment) (*string, error) {
	var (
		applicantEligibilities []entity.ApplicantEligibility
		applicantScores        []entity.ApplicantScore
	)

	if param.ApplicantEligibilities != nil {
		for _, eligibility := range *param.ApplicantEligibilities {
			applicantEligibilities = append(applicantEligibilities, entity.ApplicantEligibility{
				ApplicantID:   int64(param.ApplicantID),
				RequirementID: int64(eligibility.RequirementID),
				Value:         eligibility.Value,
			})
		}
	}

	for _, score := range param.ApplicantScores {
		applicantScores = append(applicantScores, entity.ApplicantScore{
			ApplicantID: int64(param.ApplicantID),
			Name:        score.Name,
			Value:       score.Value,
		})
	}

	message, err := a.assessmentService.Submit(ctx, int64(param.ApplicantID), applicantEligibilities, applicantScores)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// NewAssessmentMutation .
func NewAssessmentMutation(assessmentService business.AssessmentService) AssessmentMutation {
	return AssessmentMutation{assessmentService: assessmentService}
}
