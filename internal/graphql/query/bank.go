package query

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// BankQuery .
type BankQuery struct {
	bankService sa.BankService
}

type bfilter struct {
	Limit  *int
	Cursor *string
}

// FetchBank ...
func (b BankQuery) FetchBank(ctx context.Context, filter *bfilter) (resolver.BankFeedResolver, error) {
	//bankFeed, err := b.bankService.Fetch(ctx, *filter)
	//if err != nil {
	//	return resolver.BankFeedResolver{}, err
	//}

	return resolver.BankFeedResolver{BankFeed: sa.BankFeed{}}, nil
}

// NewBankQuery ...
func NewBankQuery(bankService sa.BankService) BankQuery {
	return BankQuery{bankService: bankService}
}
