package bank

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type service struct {
	bankRepo sa.BankRepository
}

// Fetch bank service
// This is without validation
// General for public access
func (s service) Fetch(ctx context.Context, filter sa.BankFilter) (sa.BankFeed, error) {
	banks, cursor, err := s.bankRepo.Fetch(ctx, filter)
	if err != nil {
		return sa.BankFeed{}, err
	}

	bankFeed := sa.BankFeed{Cursor: cursor, Banks: banks}

	return bankFeed, nil
}

// NewBankService .
func NewBankService(bankRepo sa.BankRepository) sa.BankService {
	return service{bankRepo: bankRepo}
}
