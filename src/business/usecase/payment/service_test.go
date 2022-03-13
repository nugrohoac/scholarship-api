package payment_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	errors2 "github.com/Nusantara-Muda/scholarship-api/src/business/errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/Nusantara-Muda/scholarship-api/mocks"
	_service "github.com/Nusantara-Muda/scholarship-api/src/business/usecase/payment"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestPaymentSubmitTransfer(t *testing.T) {
	var (
		payments    = make([]entity.Payment, 0)
		scholarship entity.Scholarship
		user        entity.User
	)

	testdata.GoldenJSONUnmarshal(t, "payments", &payments)
	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "user", &user)

	scholarship.ID = payments[0].ScholarshipID
	scholarship.SponsorID = user.ID
	scholarshipPaid := scholarship
	scholarshipPaid.Status = 1

	otherUser := user
	otherUser.ID = 9999

	ctxValid := common.SetUserOnContext(context.Background(), user)
	ctxOtherUser := common.SetUserOnContext(context.Background(), otherUser)
	ctxInvalid := context.Background()

	payment := payments[0]
	payment.Deadline = time.Time{}

	payments[0].TransferDate = time.Time{}
	payments[0].BankAccountName = ""
	payments[0].Image = entity.Image{}

	paymentOutOfDeadline := payments[0]
	paymentOutOfDeadline.TransferDate = payments[0].Deadline.Add(32 * time.Hour)

	tests := map[string]struct {
		paramCtx           context.Context
		paramPayment       entity.Payment
		getScholarshipByID testdata.FuncCaller
		fetchPayments      testdata.FuncCaller
		submitTransfer     testdata.FuncCaller
		expectedResp       entity.Payment
		expectedErr        error
	}{
		"error get user on context": {
			paramCtx:           ctxInvalid,
			paramPayment:       entity.Payment{},
			getScholarshipByID: testdata.FuncCaller{},
			fetchPayments:      testdata.FuncCaller{},
			submitTransfer:     testdata.FuncCaller{},
			expectedResp:       entity.Payment{},
			expectedErr:        errors2.ErrBadRequest{Message: "context doesn't contain user"},
		},
		"error get scholarship by id": {
			paramCtx:     ctxValid,
			paramPayment: payment,
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment.ScholarshipID},
				Output:   []interface{}{entity.Scholarship{}, errors.New("error")},
			},
			fetchPayments:  testdata.FuncCaller{},
			submitTransfer: testdata.FuncCaller{},
			expectedResp:   entity.Payment{},
			expectedErr:    errors.New("error"),
		},
		"user not owner of scholarship": {
			paramCtx:     ctxOtherUser,
			paramPayment: payment,
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment.ScholarshipID},
				Output:   []interface{}{scholarship, nil},
			},
			fetchPayments:  testdata.FuncCaller{},
			submitTransfer: testdata.FuncCaller{},
			expectedResp:   entity.Payment{},
			expectedErr:    errors2.ErrUnAuthorize{Message: "user is not owner of scholarship"},
		},
		"scholarship was paid": {
			paramCtx:     ctxValid,
			paramPayment: payment,
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment.ScholarshipID},
				Output:   []interface{}{scholarshipPaid, nil},
			},
			fetchPayments:  testdata.FuncCaller{},
			submitTransfer: testdata.FuncCaller{},
			expectedResp:   entity.Payment{},
			expectedErr:    errors2.ErrBadRequest{Message: "scholarship was paid"},
		},
		"error fetch payments": {
			paramCtx:     ctxValid,
			paramPayment: payment,
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment.ScholarshipID},
				Output:   []interface{}{scholarship, nil},
			},
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{payment.ScholarshipID}},
				Output:   []interface{}{nil, errors.New("error")},
			},
			submitTransfer: testdata.FuncCaller{},
			expectedResp:   entity.Payment{},
			expectedErr:    errors.New("error"),
		},
		"payment not found": {
			paramCtx:     ctxValid,
			paramPayment: payment,
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment.ScholarshipID},
				Output:   []interface{}{scholarship, nil},
			},
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{payment.ScholarshipID}},
				Output:   []interface{}{nil, nil},
			},
			submitTransfer: testdata.FuncCaller{},
			expectedResp:   entity.Payment{},
			expectedErr:    errors2.ErrNotFound{Message: "payment not found"},
		},
		"payment out of range": {
			paramCtx:     ctxValid,
			paramPayment: paymentOutOfDeadline,
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment.ScholarshipID},
				Output:   []interface{}{scholarship, nil},
			},
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{paymentOutOfDeadline.ScholarshipID}},
				Output:   []interface{}{payments, nil},
			},
			submitTransfer: testdata.FuncCaller{},
			expectedResp:   entity.Payment{},
			expectedErr:    errors2.ErrNotAllowed{Message: "payment out of range deadline"},
		},
		"error submit payment": {
			paramCtx:     ctxValid,
			paramPayment: payment,
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment.ScholarshipID},
				Output:   []interface{}{scholarship, nil},
			},
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{payment.ScholarshipID}},
				Output:   []interface{}{payments, nil},
			},
			submitTransfer: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment},
				Output:   []interface{}{entity.Payment{}, errors.New("error")},
			},
			expectedResp: entity.Payment{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramCtx:     ctxValid,
			paramPayment: payment,
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, payment.ScholarshipID},
				Output:   []interface{}{scholarship, nil},
			},
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
			scholarshipRepoMock := new(mocks.ScholarshipRepository)

			if test.getScholarshipByID.IsCalled {
				scholarshipRepoMock.On("GetByID", test.getScholarshipByID.Input...).
					Return(test.getScholarshipByID.Output...).
					Once()
			}

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

			paymentService := _service.NewPaymentService(paymentRepoMock, scholarshipRepoMock)
			paymentResp, err := paymentService.SubmitTransfer(test.paramCtx, test.paramPayment)
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
