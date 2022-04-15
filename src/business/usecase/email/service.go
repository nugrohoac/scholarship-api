package email

import (
	"context"
	"fmt"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"golang.org/x/text/message"
)

type emailService struct {
	emailRepo       business.EmailRepository
	applicantRepo   business.ApplicantRepository
	scholarshipRepo business.ScholarshipRepository
	jwtHash         business.JwtHash
	printer         *message.Printer
}

// NotifyFundingConformation .
func (e emailService) NotifyFundingConformation(ctx context.Context, scholarshipID int64) (string, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return "", err
	}

	scholarship, err := e.scholarshipRepo.GetByID(ctx, scholarshipID)
	if err != nil {
		return "", err
	}

	amount := scholarship.Amount / scholarship.Awardee
	amountStr := e.printer.Sprintf("%d\n", amount)

	// get applicant which status is 2 = awardee
	applicants, _, err := e.applicantRepo.Fetch(ctx, entity.FilterApplicant{
		ScholarshipID: scholarshipID,
		Status:        []int32{2},
	})
	if err != nil {
		return "", err
	}

	if len(applicants) > 0 {
		data := ""

		for _, applicant := range applicants {
			data = data + fmt.Sprintf("<li>%s - Rp %s</li>\n", applicant.User.Name, amountStr)
		}

		token, err := e.jwtHash.Encode(user)
		if err != nil {
			return "", err
		}

		if err = e.emailRepo.NotifyFundingConformation(ctx, user.Email, token, data); err != nil {
			return "", err
		}
	}

	return "please check email", nil
}

// NewEmailService .
func NewEmailService(
	emailRepo business.EmailRepository,
	applicantRepo business.ApplicantRepository,
	scholarshipRepo business.ScholarshipRepository,
	jwtHash business.JwtHash,
	printer *message.Printer,
) business.EmailService {
	return emailService{
		emailRepo:       emailRepo,
		applicantRepo:   applicantRepo,
		scholarshipRepo: scholarshipRepo,
		jwtHash:         jwtHash,
		printer:         printer,
	}
}
