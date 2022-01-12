package query_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestBankQuery_Fetch(t *testing.T) {
	var (
		limit  int32 = 10
		cursor       = "cursor"
	)

	banks := make([]sa.Bank, 0)
	testdata.GoldenJSONUnmarshal(t, "banks", &banks)

	bankFeed := sa.BankFeed{
		Cursor: cursor,
		Banks:  banks,
	}

	bankResolver := resolver.BankFeedResolver{BankFeed: bankFeed}

	tests := map[string]struct {
		paramCtx     context.Context
		paramFilter  sa.InputBankFilter
		fetchBank    testdata.FuncCaller
		expectedResp *resolver.BankFeedResolver
		expectedErr  error
	}{
		"success": {
			paramCtx: context.Background(),
			paramFilter: sa.InputBankFilter{
				Limit:  &limit,
				Cursor: &cursor,
			},
			fetchBank: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.BankFilter{Limit: int(limit), Cursor: cursor}},
				Output:   []interface{}{bankFeed, nil},
			},
			expectedResp: &bankResolver,
			expectedErr:  nil,
		},
		"error": {
			paramCtx: context.Background(),
			paramFilter: sa.InputBankFilter{
				Limit:  &limit,
				Cursor: &cursor,
			},
			fetchBank: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.BankFilter{Limit: int(limit), Cursor: cursor}},
				Output:   []interface{}{sa.BankFeed{}, errors.New("internal server error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			bankServiceMock := new(mocks.BankService)

			if test.fetchBank.IsCalled {
				bankServiceMock.On("Fetch", test.fetchBank.Input...).
					Return(test.fetchBank.Output...).
					Once()
			}

			bankQuery := query.NewBankQuery(bankServiceMock)
			bankResp, err := bankQuery.FetchBank(test.paramCtx, test.paramFilter)
			bankServiceMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, bankResp)
		})
	}
}
