package email

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"

	"golang.org/x/text/message"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"
)

type emailService struct {
	emailRepo       business.EmailRepository
	applicantRepo   business.ApplicantRepository
	scholarshipRepo business.ScholarshipRepository
	jwtHash72Hour   business.JwtHash
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

	if scholarship.Status >= 6 {
		return "", errors.ErrNotAllowed{Message: "scholarship has been blazing email"}
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

		token, err := e.jwtHash72Hour.Encode(user)
		if err != nil {
			return "", err
		}

		if err = e.emailRepo.NotifyFundingConformation(ctx, user.Email, token, scholarshipID, data); err != nil {
			return "", err
		}
	}

	return "please check email", nil
}

// BlazingToAwardee .
func (e emailService) BlazingToAwardee(ctx context.Context, scholarshipID int64) (string, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return "", err
	}

	scholarship, err := e.scholarshipRepo.GetByID(ctx, scholarshipID)
	if err != nil {
		return "", err
	}

	if user.ID != scholarship.SponsorID {
		return "", errors.ErrNotAllowed{Message: "sponsor doesnt own of scholarship"}
	}

	applicants, _, err := e.applicantRepo.Fetch(ctx, entity.FilterApplicant{
		ScholarshipID: scholarshipID,
		Status:        []int32{2},
	})
	if err != nil {
		return "", err
	}

	mapEmailToken := map[string]string{}
	userIDs := make([]int64, 0)
	for _, applicant := range applicants {
		token, err := e.jwtHash72Hour.Encode(applicant.User)
		if err != nil {
			return "", err
		}

		mapEmailToken[applicant.User.Email] = token
		userIDs = append(userIDs, applicant.UserID)
	}

	go func() {
		if err = e.applicantRepo.SetStatusWaitForConfirmation(context.Background(), userIDs, scholarshipID); err != nil {
			logrus.Error("failed st status waiting for confirmation : ", err)
			return
		}

		if err = e.emailRepo.BlazingToAwardee(context.Background(), mapEmailToken, scholarship); err != nil {
			logrus.Error("error blazing email to awardee", err)
		}
	}()

	return "please check email of applicants", nil
}

// ConfirmAwardee .
func (e emailService) ConfirmAwardee(ctx context.Context, scholarshipID int64) (string, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return "", err
	}

	applicants, _, err := e.applicantRepo.Fetch(ctx, entity.FilterApplicant{
		ScholarshipID: scholarshipID,
		Status:        []int32{3},
		UserID:        user.ID,
	})
	if err != nil {
		return "", err
	}

	if len(applicants) == 0 {
		return "", errors.ErrNotFound{Message: "applicant not found"}
	}

	if err = e.applicantRepo.SetStatusConfirmation(ctx, user.ID, scholarshipID); err != nil {
		return "", err
	}

	scholarship, err := e.scholarshipRepo.GetByID(ctx, scholarshipID)
	if err != nil {
		return "", err
	}

	if err = e.emailRepo.SuccessConfirmAwardee(ctx, user, scholarship.Name); err != nil {
		return "", err
	}

	return "success", nil
}

// NewEmailService .
func NewEmailService(
	emailRepo business.EmailRepository,
	applicantRepo business.ApplicantRepository,
	scholarshipRepo business.ScholarshipRepository,
	jwtHash72Hour business.JwtHash,
	printer *message.Printer,
) business.EmailService {
	return emailService{
		emailRepo:       emailRepo,
		applicantRepo:   applicantRepo,
		scholarshipRepo: scholarshipRepo,
		jwtHash72Hour:   jwtHash72Hour,
		printer:         printer,
	}
}
