package query

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"strings"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// BankQuery .
type BankQuery struct {
	bankService business.BankService
}

// FetchBank ...
func (b *BankQuery) FetchBank(ctx context.Context, param entity.InputBankFilter) (*resolver.BankFeedResolver, error) {
	filter := entity.BankFilter{}
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
func NewBankQuery(bankService business.BankService) BankQuery {
	return BankQuery{bankService: bankService}
}
