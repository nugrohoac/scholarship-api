package query

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
)

// EmailQuery .
type EmailQuery struct {
	emailService business.EmailService
}

// NotifyFundingConfirmation .
func (e EmailQuery) NotifyFundingConfirmation(ctx context.Context, param struct{ ScholarshipID int32 }) (*string, error) {
	message, err := e.emailService.NotifyFundingConformation(ctx, int64(param.ScholarshipID))
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// NewEmailQuery .
func NewEmailQuery(emailService business.EmailService) EmailQuery {
	return EmailQuery{emailService: emailService}
}
