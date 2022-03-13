package payment

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"
)

type paymentService struct {
	paymentRepo     business.PaymentRepository
	scholarshipRepo business.ScholarshipRepository
}

// SubmitTransfer ....
func (p paymentService) SubmitTransfer(ctx context.Context, payment entity.Payment) (entity.Payment, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return entity.Payment{}, err
	}

	scholarship, err := p.scholarshipRepo.GetByID(ctx, payment.ScholarshipID)
	if err != nil {
		return entity.Payment{}, err
	}

	if user.ID != scholarship.SponsorID {
		return entity.Payment{}, errors.ErrUnAuthorize{Message: "user is not owner of scholarship"}
	}

	if scholarship.Status != 0 {
		return entity.Payment{}, errors.ErrBadRequest{Message: "scholarship was paid"}
	}

	payments, err := p.paymentRepo.Fetch(ctx, []int64{payment.ScholarshipID})
	if err != nil {
		return entity.Payment{}, err
	}

	if len(payments) == 0 {
		return entity.Payment{}, errors.ErrNotFound{Message: "payment not found"}
	}

	currentPayment := payments[0]
	if payment.TransferDate.After(currentPayment.Deadline) {
		return entity.Payment{}, errors.ErrNotAllowed{Message: "payment out of range deadline"}
	}

	return p.paymentRepo.SubmitTransfer(ctx, payment)
}

// NewPaymentService ....
func NewPaymentService(paymentRepo business.PaymentRepository, scholarshipRepo business.ScholarshipRepository) business.PaymentService {
	return paymentService{
		paymentRepo:     paymentRepo,
		scholarshipRepo: scholarshipRepo,
	}
}
