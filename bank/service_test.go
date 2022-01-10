package bank_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/bank"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

var cursor = "next-cursor"

func TestService_Fetch(t *testing.T) {

	banks := make([]sa.Bank, 0)
	testdata.GoldenJSONUnmarshal(t, "banks", &banks)

	tests := map[string]struct {
		paramFilter  sa.BankFilter
		fetchBank    testdata.FuncCaller
		expectedResp sa.BankFeed
		expectedErr  error
	}{
		"success": {
			paramFilter: sa.BankFilter{Limit: 10},
			fetchBank: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.BankFilter{Limit: 10}},
				Output:   []interface{}{banks, cursor, nil},
			},
			expectedResp: sa.BankFeed{Cursor: cursor, Banks: banks},
			expectedErr:  nil,
		},
		"error": {
			paramFilter: sa.BankFilter{},
			fetchBank: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.BankFilter{}},
				Output:   []interface{}{nil, "", errors.New("internal server error")},
			},
			expectedResp: sa.BankFeed{},
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			bankRepoMock := new(mocks.BankRepository)

			if test.fetchBank.IsCalled {
				bankRepoMock.On("Fetch", test.fetchBank.Input...).
					Return(test.fetchBank.Output...).
					Once()
			}

			bankService := bank.NewBankService(bankRepoMock)
			bankFeed, err := bankService.Fetch(context.Background(), test.paramFilter)
			bankRepoMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, bankFeed)
		})
	}
}
