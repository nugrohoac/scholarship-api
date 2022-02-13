package resolver

import sa "github.com/Nusantara-Muda/scholarship-api"

// BankTransferResolver ...
type BankTransferResolver struct {
	BankTransfer sa.BankTransfer
}

// Name ....
func (b BankTransferResolver) Name() *string {
	return &b.BankTransfer.Name
}

// AccountName ...
func (b BankTransferResolver) AccountName() *string {
	return &b.BankTransfer.AccountName
}

// AccountNo ...
func (b BankTransferResolver) AccountNo() *int32 {
	accountNo := int32(b.BankTransfer.AccountNo)
	return &accountNo
}

// Image ...
func (b BankTransferResolver) Image() *ImageResolver {
	return &ImageResolver{Image: b.BankTransfer.Image}
}
