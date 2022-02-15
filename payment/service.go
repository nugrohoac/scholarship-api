package payment

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type paymentService struct {
	paymentRepo     sa.PaymentRepository
	scholarshipRepo sa.ScholarshipRepository
}

// SubmitTransfer ....
func (p paymentService) SubmitTransfer(ctx context.Context, payment sa.Payment) (sa.Payment, error) {
	user, err := sa.GetUserOnContext(ctx)
	if err != nil {
		return sa.Payment{}, err
	}

	scholarship, err := p.scholarshipRepo.GetByID(ctx, payment.ScholarshipID)
	if err != nil {
		return sa.Payment{}, err
	}

	if user.ID != scholarship.SponsorID {
		return sa.Payment{}, sa.ErrUnAuthorize{Message: "user is not owner of scholarship"}
	}

	if scholarship.Status != 0 {
		return sa.Payment{}, sa.ErrBadRequest{Message: "scholarship was paid"}
	}

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
func NewPaymentService(paymentRepo sa.PaymentRepository, scholarshipRepo sa.ScholarshipRepository) sa.PaymentService {
	return paymentService{
		paymentRepo:     paymentRepo,
		scholarshipRepo: scholarshipRepo,
	}
}
