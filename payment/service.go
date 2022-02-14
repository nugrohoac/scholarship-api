package payment

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type paymentService struct {
	paymentRepo sa.PaymentRepository
}

// SubmitTransfer ....
func (p paymentService) SubmitTransfer(ctx context.Context, payment sa.Payment) (sa.Payment, error) {
	payments, err := p.paymentRepo.Fetch(ctx, []int64{payment.ScholarshipID})
	if err != nil {
		return sa.Payment{}, err
	}

	if len(payments) == 0 {
		return sa.Payment{}, sa.ErrNotFound{Message: "payment not found"}
	}

	currentPayment := payments[0]
	if payment.TransferDate.After(currentPayment.Deadline) {
		return sa.Payment{}, sa.ErrNotAllowed{Message: "payment out of range deadline"}
	}

	return p.paymentRepo.SubmitTransfer(ctx, payment)
}

// NewPaymentService ....
func NewPaymentService(paymentRepo sa.PaymentRepository) sa.PaymentService {
	return paymentService{paymentRepo: paymentRepo}
}
