package resolver

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"time"
)

// PaymentResolver .
type PaymentResolver struct {
	Payment entity.Payment
}

// ID .
func (p PaymentResolver) ID() *int32 {
	ID := int32(p.Payment.ID)
	return &ID
}

// ScholarshipID .
func (p PaymentResolver) ScholarshipID() *int32 {
	scholarshipID := int32(p.Payment.ScholarshipID)
	return &scholarshipID
}

// BankTransfer .
func (p PaymentResolver) BankTransfer() *BankTransferResolver {
	return &BankTransferResolver{BankTransfer: p.Payment.BankTransfer}
}

// Deadline ...
func (p PaymentResolver) Deadline() *string {
	deadline := p.Payment.Deadline.Format(time.RFC3339)
	return &deadline
}

// TransferDate ...
func (p PaymentResolver) TransferDate() *string {
	transferDate := p.Payment.TransferDate.Format(time.RFC3339)
	return &transferDate
}

// BankAccountName ...
func (p PaymentResolver) BankAccountName() *string {
	return &p.Payment.BankAccountName
}

// BankAccountNo ...
func (p PaymentResolver) BankAccountNo() *string {
	return &p.Payment.BankAccountNo
}

// Image ...
func (p PaymentResolver) Image() *ImageResolver {
	return &ImageResolver{Image: p.Payment.Image}
}
