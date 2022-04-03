package scholarship_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	errors2 "github.com/Nusantara-Muda/scholarship-api/src/business/errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/Nusantara-Muda/scholarship-api/mocks"
	_service "github.com/Nusantara-Muda/scholarship-api/src/business/usecase/scholarship"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

var cursor = "next-cursor"

func TestScholarshipServiceCreate(t *testing.T) {

	var (
		scholarship  entity.Scholarship
		user         entity.User
		payment      entity.Payment
		bankTransfer entity.BankTransfer
	)

	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "user", &user)
	testdata.GoldenJSONUnmarshal(t, "payment", &payment)
	testdata.GoldenJSONUnmarshal(t, "bank_transfer", &bankTransfer)

	user.Status = 2

	scholarship.Sponsor = user
	scholarship.SponsorID = user.ID

	userUnCompleteProfile := user
	userUnCompleteProfile.Status = 0

	ctxValid := common.SetUserOnContext(context.Background(), user)
	ctxUnComplete := common.SetUserOnContext(context.Background(), userUnCompleteProfile)

	scholarshipNotMatch := scholarship
	scholarshipNotMatch.SponsorID = 2

	scholarshipInvalid := scholarship
	fundingStart := scholarshipInvalid.FundingEnd
	scholarshipInvalid.FundingEnd = scholarshipInvalid.FundingStart
	scholarshipInvalid.FundingStart = fundingStart

	scholarshipResp := scholarship
	scholarshipResp.Payment.ID = payment.ScholarshipID
	scholarshipResp.Payment.ScholarshipID = scholarship.ID
	scholarshipResp.Payment.Deadline = payment.Deadline
	scholarshipResp.Payment.BankTransfer = bankTransfer

	tests := map[string]struct {
		paramCtx          context.Context
		paramScholarship  entity.Scholarship
		createScholarship testdata.FuncCaller
		getBankTransfer   testdata.FuncCaller
		expectedResp      entity.Scholarship
		expectedErr       error
	}{
		"error get sponsor on context": {
			paramCtx:          context.Background(),
			paramScholarship:  scholarship,
			createScholarship: testdata.FuncCaller{},
			expectedResp:      entity.Scholarship{},
			expectedErr:       errors2.ErrBadRequest{Message: "context doesn't contain user"},
		},
		"sponsor is not match with token": {
			paramCtx:          ctxValid,
			paramScholarship:  scholarshipNotMatch,
			createScholarship: testdata.FuncCaller{},
			expectedResp:      entity.Scholarship{},
			expectedErr:       errors2.ErrUnAuthorize{Message: "sponsor id is not match"},
		},
		"sponsor status un complete": {
			paramCtx:          ctxUnComplete,
			paramScholarship:  scholarship,
			createScholarship: testdata.FuncCaller{},
			expectedResp:      entity.Scholarship{},
			expectedErr:       errors2.ErrNotAllowed{Message: "sponsor un complete profile"},
		},
		"funding end before funding start": {
			paramCtx:          ctxValid,
			paramScholarship:  scholarshipInvalid,
			createScholarship: testdata.FuncCaller{},
			expectedResp:      entity.Scholarship{},
			expectedErr:       errors2.ErrBadRequest{Message: "scholarship funding end before funding start"},
		},
		"failed create scholarship": {
			paramCtx:         ctxValid,
			paramScholarship: scholarship,
			createScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctxValid, scholarship},
				Output:   []interface{}{entity.Scholarship{}, errors.New("error")},
			},
			expectedResp: entity.Scholarship{},
			expectedErr:  errors.New("error"),
		},
		"failed get bank transfer": {
			paramCtx:         ctxValid,
			paramScholarship: scholarship,
			createScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctxValid, scholarship},
				Output:   []interface{}{scholarshipResp, nil},
			},
			getBankTransfer: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything},
				Output:   []interface{}{entity.BankTransfer{}, errors.New("error")},
			},
			expectedResp: entity.Scholarship{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramCtx:         ctxValid,
			paramScholarship: scholarship,
			createScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctxValid, scholarship},
				Output:   []interface{}{scholarshipResp, nil},
			},
			getBankTransfer: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything},
				Output:   []interface{}{bankTransfer, nil},
			},
			expectedResp: scholarshipResp,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			scholarshipRepoMock := new(mocks.ScholarshipRepository)
			bankTransferRepoMock := new(mocks.BankTransferRepository)
			paymentRepoMock := new(mocks.PaymentRepository)
			requirementDescRepoMock := new(mocks.RequirementDescriptionRepository)

			if test.createScholarship.IsCalled {
				scholarshipRepoMock.On("Create", test.createScholarship.Input...).
					Return(test.createScholarship.Output...).
					Once()
			}

			if test.getBankTransfer.IsCalled {
				bankTransferRepoMock.On("Get", test.getBankTransfer.Input...).
					Return(test.getBankTransfer.Output...).
					Once()
			}

			scholarshipService := _service.NewScholarshipService(scholarshipRepoMock, bankTransferRepoMock, paymentRepoMock, requirementDescRepoMock)
			scholarshipResp, err := scholarshipService.Create(test.paramCtx, test.paramScholarship)
			scholarshipRepoMock.AssertExpectations(t)
			bankTransferRepoMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, scholarshipResp)
		})

	}

}

func TestScholarshipServiceFetch(t *testing.T) {
	var (
		scholarships = make([]entity.Scholarship, 0)
		sponsor      entity.User
	)

	testdata.GoldenJSONUnmarshal(t, "scholarships", &scholarships)
	testdata.GoldenJSONUnmarshal(t, "user", &sponsor)

	scholarships[0].SponsorID = 1
	scholarships[1].SponsorID = 1

	requirementDesc := map[int64][]string{}
	requirementDesc[scholarships[0].ID] = scholarships[0].RequirementDescriptions
	requirementDesc[scholarships[1].ID] = scholarships[1].RequirementDescriptions

	scholarshipIDs := []int64{scholarships[0].ID, scholarships[1].ID}

	tests := map[string]struct {
		paramFilter            entity.ScholarshipFilter
		fetchScholarship       testdata.FuncCaller
		fetchRequirementDesc   testdata.FuncCaller
		fetchScholarshipCursor testdata.FuncCaller
		expectedResp           entity.ScholarshipFeed
		expectedErr            error
	}{
		"error fetch scholarship": {
			paramFilter: entity.ScholarshipFilter{},
			fetchScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.ScholarshipFilter{}},
				Output:   []interface{}{nil, "", errors.New("error")},
			},
			fetchRequirementDesc:   testdata.FuncCaller{},
			fetchScholarshipCursor: testdata.FuncCaller{},
			expectedResp:           entity.ScholarshipFeed{},
			expectedErr:            errors.New("error"),
		},
		"error etch requirement description": {
			paramFilter: entity.ScholarshipFilter{},
			fetchScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.ScholarshipFilter{}},
				Output:   []interface{}{nil, "", errors.New("error")},
			},
			fetchRequirementDesc: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarshipIDs},
				Output:   []interface{}{nil, errors.New("error")},
			},
			fetchScholarshipCursor: testdata.FuncCaller{},
			expectedResp:           entity.ScholarshipFeed{},
			expectedErr:            errors.New("error"),
		},
		"error fetch scholarship for cursor": {
			paramFilter: entity.ScholarshipFilter{},
			fetchScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.ScholarshipFilter{}},
				Output:   []interface{}{scholarships, cursor, nil},
			},
			fetchRequirementDesc: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarshipIDs},
				Output:   []interface{}{requirementDesc, nil},
			},
			fetchScholarshipCursor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.ScholarshipFilter{Cursor: cursor, Limit: 1}},
				Output:   []interface{}{nil, "", errors.New("error")},
			},
			expectedResp: entity.ScholarshipFeed{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramFilter: entity.ScholarshipFilter{},
			fetchScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.ScholarshipFilter{}},
				Output:   []interface{}{scholarships, cursor, nil},
			},
			fetchRequirementDesc: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarshipIDs},
				Output:   []interface{}{requirementDesc, nil},
			},
			fetchScholarshipCursor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.ScholarshipFilter{Cursor: cursor, Limit: 1}},
				Output:   []interface{}{[]entity.Scholarship{scholarships[0]}, cursor, nil},
			},
			expectedResp: entity.ScholarshipFeed{Cursor: cursor, Scholarships: scholarships},
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			scholarshipRepoMock := new(mocks.ScholarshipRepository)
			bankTransferRepoMock := new(mocks.BankTransferRepository)
			paymentRepoMock := new(mocks.PaymentRepository)
			requirementDescRepoMock := new(mocks.RequirementDescriptionRepository)

			if test.fetchScholarship.IsCalled {
				scholarshipRepoMock.On("Fetch", test.fetchScholarship.Input...).
					Return(test.fetchScholarship.Output...).
					Once()
			}

			if test.fetchRequirementDesc.IsCalled {
				requirementDescRepoMock.On("Fetch", test.fetchRequirementDesc.Input...).
					Return(test.fetchRequirementDesc.Output...).
					Once()
			}

			if test.fetchScholarshipCursor.IsCalled {
				scholarshipRepoMock.On("Fetch", test.fetchScholarshipCursor.Input...).
					Return(test.fetchScholarshipCursor.Output...).
					Once()
			}

			scholarshipService := _service.NewScholarshipService(scholarshipRepoMock, bankTransferRepoMock, paymentRepoMock, requirementDescRepoMock)
			resp, err := scholarshipService.Fetch(context.Background(), test.paramFilter)
			scholarshipRepoMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, resp)
		})
	}
}

func TestScholarshipServiceGetByID(t *testing.T) {
	var (
		scholarship  entity.Scholarship
		payment      entity.Payment
		bankTransfer entity.BankTransfer
	)

	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "payment", &payment)
	testdata.GoldenJSONUnmarshal(t, "bank_transfer", &bankTransfer)

	payment.ScholarshipID = scholarship.ID

	scholarshipPaid := scholarship
	scholarshipPaid.Status = 1

	scholarshipUnPaid := scholarship
	scholarshipUnPaid.Status = 0

	scholarshipResponse := scholarshipUnPaid
	scholarshipResponse.Payment = payment
	scholarshipResponse.Payment.BankTransfer = bankTransfer

	requirementDesc := map[int64][]string{}
	requirementDesc[scholarship.ID] = scholarship.RequirementDescriptions

	tests := map[string]struct {
		paramID              int64
		getScholarship       testdata.FuncCaller
		fetchRequirementDesc testdata.FuncCaller
		fetchPayments        testdata.FuncCaller
		getBankTransfer      testdata.FuncCaller
		expectedResp         entity.Scholarship
		expectedErr          error
	}{
		"error get scholarship": {
			paramID: scholarship.ID,
			getScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarship.ID},
				Output:   []interface{}{entity.Scholarship{}, errors.New("error")},
			},
			fetchPayments:   testdata.FuncCaller{},
			getBankTransfer: testdata.FuncCaller{},
			expectedResp:    entity.Scholarship{},
			expectedErr:     errors.New("error"),
		},
		"scholarship not found": {
			paramID: scholarship.ID,
			getScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarshipUnPaid.ID},
				Output:   []interface{}{entity.Scholarship{}, nil},
			},
			fetchRequirementDesc: testdata.FuncCaller{},
			fetchPayments:        testdata.FuncCaller{},
			getBankTransfer:      testdata.FuncCaller{},
			expectedResp:         entity.Scholarship{},
			expectedErr:          errors2.ErrNotFound{Message: "scholarship is not found"},
		},
		"error fetch requirement description": {
			paramID: scholarship.ID,
			getScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarshipUnPaid.ID},
				Output:   []interface{}{scholarshipUnPaid, nil},
			},
			fetchRequirementDesc: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{scholarship.ID}},
				Output:   []interface{}{nil, errors.New("error")},
			},
			fetchPayments:   testdata.FuncCaller{},
			getBankTransfer: testdata.FuncCaller{},
			expectedResp:    entity.Scholarship{},
			expectedErr:     errors.New("error"),
		},
		"error fetch payment": {
			paramID: scholarship.ID,
			getScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarshipUnPaid.ID},
				Output:   []interface{}{scholarshipUnPaid, nil},
			},
			fetchRequirementDesc: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{scholarship.ID}},
				Output:   []interface{}{requirementDesc, nil},
			},
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{scholarshipUnPaid.ID}},
				Output:   []interface{}{nil, errors.New("error")},
			},
			getBankTransfer: testdata.FuncCaller{},
			expectedResp:    entity.Scholarship{},
			expectedErr:     errors.New("error"),
		},
		"error get bank transfer": {
			paramID: scholarship.ID,
			getScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarshipUnPaid.ID},
				Output:   []interface{}{scholarshipUnPaid, nil},
			},
			fetchRequirementDesc: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{scholarship.ID}},
				Output:   []interface{}{requirementDesc, nil},
			},
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{scholarshipUnPaid.ID}},
				Output:   []interface{}{[]entity.Payment{payment}, nil},
			},
			getBankTransfer: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything},
				Output:   []interface{}{entity.BankTransfer{}, errors.New("error")},
			},
			expectedResp: entity.Scholarship{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramID: scholarship.ID,
			getScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarshipUnPaid.ID},
				Output:   []interface{}{scholarshipUnPaid, nil},
			},
			fetchRequirementDesc: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{scholarship.ID}},
				Output:   []interface{}{requirementDesc, nil},
			},
			fetchPayments: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, []int64{scholarshipUnPaid.ID}},
				Output:   []interface{}{[]entity.Payment{payment}, nil},
			},
			getBankTransfer: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything},
				Output:   []interface{}{bankTransfer, nil},
			},
			expectedResp: scholarshipResponse,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			scholarshipRepoMock := new(mocks.ScholarshipRepository)
			paymentRepoMock := new(mocks.PaymentRepository)
			bankTransferRepoMock := new(mocks.BankTransferRepository)
			requirementDescRepoMock := new(mocks.RequirementDescriptionRepository)

			if test.getScholarship.IsCalled {
				scholarshipRepoMock.On("GetByID", test.getScholarship.Input...).
					Return(test.getScholarship.Output...).
					Once()
			}

			if test.fetchRequirementDesc.IsCalled {
				requirementDescRepoMock.On("Fetch", test.fetchRequirementDesc.Input...).
					Return(test.fetchRequirementDesc.Output...).
					Once()
			}

			if test.fetchPayments.IsCalled {
				paymentRepoMock.On("Fetch", test.fetchPayments.Input...).
					Return(test.fetchPayments.Output...).
					Once()
			}

			if test.getBankTransfer.IsCalled {
				bankTransferRepoMock.On("Get", test.getBankTransfer.Input...).
					Return(test.getBankTransfer.Output...).
					Once()
			}

			scholarshipService := _service.NewScholarshipService(scholarshipRepoMock, bankTransferRepoMock, paymentRepoMock, requirementDescRepoMock)
			scholarshipResp, err := scholarshipService.GetByID(context.Background(), test.paramID)
			scholarshipRepoMock.AssertExpectations(t)
			bankTransferRepoMock.AssertExpectations(t)
			paymentRepoMock.AssertExpectations(t)
			requirementDescRepoMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, scholarshipResp)
		})
	}
}
