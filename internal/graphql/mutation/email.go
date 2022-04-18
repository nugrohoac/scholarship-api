package mutation

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
)

// EmailMutation .
type EmailMutation struct {
	emailService business.EmailService
}

// ConfirmAwardee .
func (e EmailMutation) ConfirmAwardee(ctx context.Context, param struct{ ScholarshipID int32 }) (*string, error) {
	message, err := e.emailService.ConfirmAwardee(ctx, int64(param.ScholarshipID))
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// NewEmailMutation .
func NewEmailMutation(emailService business.EmailService) EmailMutation {
	return EmailMutation{emailService: emailService}
}
