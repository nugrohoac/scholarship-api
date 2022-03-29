package resolver

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"strconv"
)

// BankTransferResolver ...
type BankTransferResolver struct {
	BankTransfer entity.BankTransfer
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
func (b BankTransferResolver) AccountNo() *string {
	accountNo := strconv.Itoa(b.BankTransfer.AccountNo)
	return &accountNo
}

// Image ...
func (b BankTransferResolver) Image() *ImageResolver {
	return &ImageResolver{Image: b.BankTransfer.Image}
}
