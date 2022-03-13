package bank_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/bank"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

var cursor = "next-cursor"

func TestService_Fetch(t *testing.T) {

	banks := make([]entity.Bank, 0)
	testdata.GoldenJSONUnmarshal(t, "banks", &banks)

	tests := map[string]struct {
		paramFilter  entity.BankFilter
		fetchBank    testdata.FuncCaller
		expectedResp entity.BankFeed
		expectedErr  error
	}{
		"success": {
			paramFilter: entity.BankFilter{Limit: 10},
			fetchBank: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.BankFilter{Limit: 10}},
				Output:   []interface{}{banks, cursor, nil},
			},
			expectedResp: entity.BankFeed{Cursor: cursor, Banks: banks},
			expectedErr:  nil,
		},
		"error": {
			paramFilter: entity.BankFilter{},
			fetchBank: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.BankFilter{}},
				Output:   []interface{}{nil, "", errors.New("internal server error")},
			},
			expectedResp: entity.BankFeed{},
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
