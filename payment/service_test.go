package payment_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	_service "github.com/Nusantara-Muda/scholarship-api/payment"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestPaymentSubmitTransfer(t *testing.T) {
	payments := make([]sa.Payment, 0)
	testdata.GoldenJSONUnmarshal(t, "payments", &payments)

	payment := payments[0]
	payment.Deadline = time.Time{}

	payments[0].TransferDate = time.Time{}
	payments[0].BankAccountName = ""
	payments[0].Image = sa.Image{}

	paymentOutOfDeadline := payments[0]
	paymentOutOfDeadline.TransferDate = payments[0].Deadline.Add(32 * time.Hour)

	tests := map[string]struct {
		paramPayment   sa.Payment
		fetchPayments  testdata.FuncCaller
		submitTransfer testdata.FuncCaller
		expectedResp   sa.Payment
		expectedErr    error
	}{
		"error fetch payment": {
			paramPayment: payment,
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{payment.ScholarshipID}},
				Output:   []interface{}{[]sa.Payment{}, errors.New("error")},
			},
			submitTransfer: testdata.FuncCaller{},
			expectedResp:   sa.Payment{},
			expectedErr:    errors.New("error"),
		},
		"payment not found": {
			paramPayment: payment,
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{payment.ScholarshipID}},
				Output:   []interface{}{[]sa.Payment{}, nil},
			},
			submitTransfer: testdata.FuncCaller{},
			expectedResp:   sa.Payment{},
			expectedErr:    sa.ErrNotFound{Message: "payment not found"},
		},
		"payment out of range": {
			paramPayment: paymentOutOfDeadline,
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{paymentOutOfDeadline.ScholarshipID}},
				Output:   []interface{}{payments, nil},
			},
			submitTransfer: testdata.FuncCaller{},
			expectedResp:   sa.Payment{},
			expectedErr:    sa.ErrNotAllowed{Message: "payment out of range deadline"},
		},
		"error submit payment": {
			paramPayment: payment,
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{payment.ScholarshipID}},
				Output:   []interface{}{payments, nil},
			},
			submitTransfer: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment},
				Output:   []interface{}{sa.Payment{}, errors.New("error")},
			},
			expectedResp: sa.Payment{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramPayment: payment,
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{payment.ScholarshipID}},
				Output:   []interface{}{payments, nil},
			},
			submitTransfer: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment},
				Output:   []interface{}{payment, nil},
			},
			expectedResp: payment,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			paymentRepoMock := new(mocks.PaymentRepository)

			if test.fetchPayments.IsCalled {
				paymentRepoMock.On("Fetch", test.fetchPayments.Input...).
					Return(test.fetchPayments.Output...).
					Once()
			}

			if test.submitTransfer.IsCalled {
				paymentRepoMock.On("SubmitTransfer", test.submitTransfer.Input...).
					Return(test.submitTransfer.Output...).
					Once()
			}

			paymentService := _service.NewPaymentService(paymentRepoMock)
			paymentResp, err := paymentService.SubmitTransfer(context.Background(), test.paramPayment)
			paymentRepoMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, paymentResp)
		})
	}
}
