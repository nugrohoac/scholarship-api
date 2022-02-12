package resolver

import (
	sa "github.com/Nusantara-Muda/scholarship-api"
	"time"
)

// PaymentResolver .
type PaymentResolver struct {
	Payment sa.Payment
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

// Image ...
func (p PaymentResolver) Image() *ImageResolver {
	return &ImageResolver{Image: p.Payment.Image}
}
