package mutation_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestPaymentSubmitTransferPayment(t *testing.T) {
	var payment sa.Payment
	testdata.GoldenJSONUnmarshal(t, "payment", &payment)

	input := sa.InputSubmitTransfer{
		ScholarshipID:   int32(payment.ScholarshipID),
		TransferDate:    payment.TransferDate.Format(time.RFC3339Nano),
		BankAccountName: payment.BankAccountName,
		Image: sa.InputImage{
			URL:    payment.Image.URL,
			Width:  payment.Image.Width,
			Height: payment.Image.Height,
		},
	}

	_payment := sa.Payment{
		ScholarshipID:   payment.ScholarshipID,
		TransferDate:    payment.TransferDate,
		BankAccountName: payment.BankAccountName,
		Image:           payment.Image,
	}

	response := resolver.PaymentResolver{Payment: _payment}

	tests := map[string]struct {
		param         sa.InputSubmitTransfer
		submitPayment testdata.FuncCaller
		expectedResp  *resolver.PaymentResolver
		expectedErr   error
	}{
		"error": {
			param: input,
			submitPayment: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, _payment},
				Output:   []interface{}{sa.Payment{}, errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
		"success": {
			param: input,
			submitPayment: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, _payment},
				Output:   []interface{}{_payment, nil},
			},
			expectedResp: &response,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			paymentServiceMock := new(mocks.PaymentService)

			if test.submitPayment.IsCalled {
				paymentServiceMock.On("SubmitTransfer", test.submitPayment.Input...).
					Return(test.submitPayment.Output...).
					Once()
			}

			paymentMutation := mutation.NewPaymentMutation(paymentServiceMock)
			response, err := paymentMutation.SubmitTransferPayment(context.Background(), test.param)
			paymentServiceMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, response)
		})
	}
}
