package email

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type emailService struct {
	emailRepo     business.EmailRepository
	applicantRepo business.ApplicantRepository
}

// NotifyFundingConformation .
func (e emailService) NotifyFundingConformation(ctx context.Context) error {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return err
	}

	// get applicant which status is 2 = awardee
	applicants, _, err := e.applicantRepo.Fetch(ctx, entity.FilterApplicant{
		SponsorID: user.ID,
		Cursor:    "",
		Status:    []int32{2},
	})

	//TODO implement me
	panic("implement me")
}

// NewEmailService .
func NewEmailService(
	emailRepo business.EmailRepository,
	applicantRepo business.ApplicantRepository) business.EmailService {
	return emailService{
		emailRepo:     nil,
		applicantRepo: nil,
	}
}
