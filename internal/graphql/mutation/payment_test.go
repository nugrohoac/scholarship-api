package mutation_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestPaymentSubmitTransferPayment(t *testing.T) {
	var payment entity.Payment
	testdata.GoldenJSONUnmarshal(t, "payment", &payment)

	input := entity.InputSubmitTransfer{
		ScholarshipID:   int32(payment.ScholarshipID),
		TransferDate:    payment.TransferDate.Format(time.RFC3339Nano),
		BankAccountName: payment.BankAccountName,
		Image: entity.InputImage{
			URL:    payment.Image.URL,
			Width:  payment.Image.Width,
			Height: payment.Image.Height,
		},
	}

	_payment := entity.Payment{
		ScholarshipID:   payment.ScholarshipID,
		TransferDate:    payment.TransferDate,
		BankAccountName: payment.BankAccountName,
		Image:           payment.Image,
	}

	response := resolver.PaymentResolver{Payment: _payment}

	tests := map[string]struct {
		param         entity.InputSubmitTransfer
		submitPayment testdata.FuncCaller
		expectedResp  *resolver.PaymentResolver
		expectedErr   error
	}{
		"error": {
			param: input,
			submitPayment: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, _payment},
				Output:   []interface{}{entity.Payment{}, errors.New("error")},
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
