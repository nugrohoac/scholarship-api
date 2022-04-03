package mutation

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"time"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// PaymentMutation ...
type PaymentMutation struct {
	paymentService business.PaymentService
}

// SubmitTransferPayment ...
func (p PaymentMutation) SubmitTransferPayment(ctx context.Context, param entity.InputSubmitTransfer) (*resolver.PaymentResolver, error) {
	payment := entity.Payment{
		ScholarshipID:   int64(param.ScholarshipID),
		TransferDate:    time.Time{},
		BankAccountName: param.BankAccountName,
		BankAccountNo:   param.BankAccountNo,
		Image: entity.Image{
			URL:    param.Image.URL,
			Width:  param.Image.Width,
			Height: param.Image.Height,
		},
	}

	transferDate, err := time.Parse(time.RFC3339, param.TransferDate)
	if err != nil {
		return nil, err
	}

	payment.TransferDate = transferDate

	payment, err = p.paymentService.SubmitTransfer(ctx, payment)
	if err != nil {
		return nil, err
	}

	return &resolver.PaymentResolver{Payment: payment}, err
}

// NewPaymentMutation ...
func NewPaymentMutation(paymentService business.PaymentService) PaymentMutation {
	return PaymentMutation{paymentService: paymentService}
}
