package bank_transfer

import sa "github.com/Nusantara-Muda/scholarship-api"

type bankTransferRepo struct {
	bankTransfer sa.BankTransfer
}

// Get ...
func (b bankTransferRepo) Get() sa.BankTransfer {
	return b.bankTransfer
}

// NewBankTransfer ...
func NewBankTransfer(bankTransfer sa.BankTransfer) sa.BankTransferRepsitory {
	return bankTransferRepo{
		bankTransfer: bankTransfer,
	}
}
