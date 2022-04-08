package assessment

import (
	"context"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"
)

type assessmentService struct {
	assessmentRepository  business.AssessmentRepository
	applicantRepository   business.ApplicantRepository
	scholarshipRepository business.ScholarshipRepository
}

// Submit .
func (a assessmentService) Submit(ctx context.Context, ApplicantID int64, eligibilities []entity.ApplicantEligibility, scores []entity.ApplicantScore) (string, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return "", err
	}

	applicant, err := a.applicantRepository.GetByID(ctx, ApplicantID)
	if err != nil {
		return "", err
	}

	scholarship, err := a.scholarshipRepository.GetByID(ctx, applicant.ScholarshipID)
	if err != nil {
		return "", err
	}

	if scholarship.SponsorID != user.ID {
		return "", errors.ErrNotAllowed{Message: "sponsor not scholarship owner"}
	}

	if err = a.assessmentRepository.Submit(ctx, ApplicantID, eligibilities, scores); err != nil {
		return "", err
	}

	return "success", nil
}

// NewAssessmentService .
func NewAssessmentService(
	assessmentRepository business.AssessmentRepository,
	applicantRepository business.ApplicantRepository,
	scholarshipRepository business.ScholarshipRepository,
) business.AssessmentService {
	return assessmentService{
		assessmentRepository:  assessmentRepository,
		applicantRepository:   applicantRepository,
		scholarshipRepository: scholarshipRepository,
	}
}
