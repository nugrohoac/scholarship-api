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
func (b *BankQuery) FetchBank(ctx context.Context, filter sa.InputBankFilter) (*resolver.BankFeedResolver, error) {
	_filter := sa.BankFilter{}
	if filter.Limit != nil {
		_filter.Limit = int(*filter.Limit)
	}

	if filter.Cursor != nil {
		_filter.Cursor = *filter.Cursor
	}

	if filter.Name != nil {
		_filter.Name = strings.ToLower(*filter.Name)
	}

	bankFeed, err := b.bankService.Fetch(ctx, _filter)
	if err != nil {
		return nil, err
	}

	return &resolver.BankFeedResolver{BankFeed: bankFeed}, nil
}

// NewBankQuery ...
func NewBankQuery(bankService sa.BankService) BankQuery {
	return BankQuery{bankService: bankService}
}
