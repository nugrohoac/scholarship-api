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
	schoolRepository      business.SchoolRepository
	userRepo              business.UserRepository
	assessmentRepo        business.AssessmentRepository
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

	applicantFeed := entity.ApplicantFeed{
		Cursor:     cursor,
		Applicants: applicants,
	}

	filter.Limit = 1
	filter.Cursor = cursor

	applicants, _, err = a.applicantRepository.Fetch(ctx, filter)
	if err != nil {
		return entity.ApplicantFeed{}, err
	}

	if len(applicants) == 0 {
		applicantFeed.Cursor = ""
	}

	return applicantFeed, nil
}

// GetByID .
func (a applicantService) GetByID(ctx context.Context, ID int64) (entity.Applicant, error) {
	_, err := common.GetUserOnContext(ctx)
	if err != nil {
		return entity.Applicant{}, err
	}

	// business process will do check base on user on context
	applicant, err := a.applicantRepository.GetByID(ctx, ID)
	if err != nil {
		return entity.Applicant{}, err
	}

	userSchools, err := a.schoolRepository.GetUserSchool(ctx, applicant.UserID)
	if err != nil {
		return entity.Applicant{}, err
	}

	applicant.User.UserSchools = userSchools

	userDocuments, err := a.userRepo.GetDocuments(ctx, applicant.UserID)
	if err != nil {
		return entity.Applicant{}, err
	}

	applicant.User.UserDocuments = userDocuments

	scores, err := a.assessmentRepo.GetScoreByApplicantIDs(ctx, []int64{ID})
	if err != nil {
		return entity.Applicant{}, err
	}

	applicant.Scores = scores

	return applicant, nil
}

// UpdateStatus .
func (a applicantService) UpdateStatus(ctx context.Context, ID int64, status int32) (string, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return "", err
	}

	applicant, err := a.applicantRepository.GetByID(ctx, ID)
	if err != nil {
		return "", err
	}

	if applicant.Status == status {
		return "", errors.ErrNotAllowed{Message: "status is same with current status"}
	}

	scholarship, err := a.scholarshipRepository.GetByID(ctx, applicant.ScholarshipID)
	if err != nil {
		return "", err
	}

	if scholarship.SponsorID != user.ID {
		return "", errors.ErrNotAllowed{Message: "sponsor not scholarship owner"}
	}

	if err = a.applicantRepository.UpdateStatus(ctx, ID, status); err != nil {
		return "", err
	}

	return "success", nil
}

// NewApplicantService .
func NewApplicantService(
	applicantRepository business.ApplicantRepository,
	scholarshipRepository business.ScholarshipRepository,
	schoolRepository business.SchoolRepository,
	userRepo business.UserRepository,
	assessmentRepo business.AssessmentRepository,
) business.ApplicantService {
	return applicantService{
		applicantRepository:   applicantRepository,
		scholarshipRepository: scholarshipRepository,
		schoolRepository:      schoolRepository,
		userRepo:              userRepo,
		assessmentRepo:        assessmentRepo,
	}
}
