package query

import (
	"context"
	"strings"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// BankQuery .
type BankQuery struct {
	bankService sa.BankService
}

// FetchBank ...
func (b *BankQuery) FetchBank(ctx context.Context, param sa.InputBankFilter) (*resolver.BankFeedResolver, error) {
	filter := sa.BankFilter{}
	if param.Limit != nil {
		filter.Limit = int(*param.Limit)
	}

	if param.Cursor != nil {
		filter.Cursor = *param.Cursor
	}

	if param.Name != nil {
		filter.Name = strings.ToLower(*param.Name)
	}

	bankFeed, err := b.bankService.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.BankFeedResolver{BankFeed: bankFeed}, nil
}

// NewBankQuery ...
func NewBankQuery(bankService sa.BankService) BankQuery {
	return BankQuery{bankService: bankService}
}
