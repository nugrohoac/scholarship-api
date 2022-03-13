package bank

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type service struct {
	bankRepo business.BankRepository
}

// Fetch bank service
// This is without validation
// General for public access
func (s *service) Fetch(ctx context.Context, filter entity.BankFilter) (entity.BankFeed, error) {
	banks, cursor, err := s.bankRepo.Fetch(ctx, filter)
	if err != nil {
		return entity.BankFeed{}, err
	}

	bankFeed := entity.BankFeed{Cursor: cursor, Banks: banks}

	return bankFeed, nil
}

// NewBankService .
func NewBankService(bankRepo business.BankRepository) business.BankService {
	return &service{bankRepo: bankRepo}
}
