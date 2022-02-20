package mutation

import (
	"context"
	"time"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// PaymentMutation ...
type PaymentMutation struct {
	paymentService sa.PaymentService
}

// SubmitTransferPayment ...
func (p PaymentMutation) SubmitTransferPayment(ctx context.Context, param sa.InputSubmitTransfer) (*resolver.PaymentResolver, error) {
	payment := sa.Payment{
		ScholarshipID:   int64(param.ScholarshipID),
		TransferDate:    time.Time{},
		BankAccountName: param.BankAccountName,
		BankAccountNo:   param.BankAccountNo,
		Image: sa.Image{
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
func NewPaymentMutation(paymentService sa.PaymentService) PaymentMutation {
	return PaymentMutation{paymentService: paymentService}
}
