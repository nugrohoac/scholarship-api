package applicant

import (
	"context"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"
)

type applicantService struct {
	applicantRepository   business.ApplicantRepository
	scholarshipRepository business.ScholarshipRepository
}

// Fetch .
func (a applicantService) Fetch(ctx context.Context, filter entity.FilterApplicant) (entity.ApplicantFeed, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return entity.ApplicantFeed{}, err
	}

	if user.ID != filter.SponsorID {
		return entity.ApplicantFeed{}, errors.ErrNotAllowed{Message: "user id is not match"}
	}

	scholarship, err := a.scholarshipRepository.GetByID(ctx, filter.ScholarshipID)
	if err != nil {
		return entity.ApplicantFeed{}, err
	}

	if scholarship.SponsorID != user.ID {
		return entity.ApplicantFeed{}, errors.ErrNotAllowed{Message: "access denied"}
	}

	applicants, cursor, err := a.applicantRepository.Fetch(ctx, filter)
	if err != nil {
		return entity.ApplicantFeed{}, err
	}

	return entity.ApplicantFeed{Cursor: cursor, Applicants: applicants}, nil
}

// NewApplicantService .
func NewApplicantService(applicantRepository business.ApplicantRepository, scholarshipRepository business.ScholarshipRepository) business.ApplicantService {
	return applicantService{
		applicantRepository:   applicantRepository,
		scholarshipRepository: scholarshipRepository,
	}
}