package scholarship_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	_service "github.com/Nusantara-Muda/scholarship-api/scholarship"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

var cursor = "next-cursor"

func TestScholarshipServiceCreate(t *testing.T) {

	var (
		scholarship sa.Scholarship
		user        sa.User
		payment     sa.Payment
	)

	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "user", &user)
	testdata.GoldenJSONUnmarshal(t, "payment", &payment)

	user.Status = 2

	scholarship.Sponsor = user
	scholarship.SponsorID = user.ID

	userUnCompleteProfile := user
	userUnCompleteProfile.Status = 0

	ctxValid := sa.SetUserOnContext(context.Background(), user)
	ctxUnComplete := sa.SetUserOnContext(context.Background(), userUnCompleteProfile)

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

	tests := map[string]struct {
		paramCtx          context.Context
		paramScholarship  sa.Scholarship
		createScholarship testdata.FuncCaller
		expectedResp      sa.Scholarship
		expectedErr       error
	}{
		"error get sponsor on context": {
			paramCtx:          context.Background(),
			paramScholarship:  scholarship,
			createScholarship: testdata.FuncCaller{},
			expectedResp:      sa.Scholarship{},
			expectedErr:       sa.ErrBadRequest{Message: "context doesn't contain user"},
		},
		"sponsor is not match with token": {
			paramCtx:          ctxValid,
			paramScholarship:  scholarshipNotMatch,
			createScholarship: testdata.FuncCaller{},
			expectedResp:      sa.Scholarship{},
			expectedErr:       sa.ErrUnAuthorize{Message: "sponsor id is not match"},
		},
		"sponsor status un complete": {
			paramCtx:          ctxUnComplete,
			paramScholarship:  scholarship,
			createScholarship: testdata.FuncCaller{},
			expectedResp:      sa.Scholarship{},
			expectedErr:       sa.ErrNotAllowed{Message: "sponsor un complete profile"},
		},
		"funding end before funding start": {
			paramCtx:          ctxValid,
			paramScholarship:  scholarshipInvalid,
			createScholarship: testdata.FuncCaller{},
			expectedResp:      sa.Scholarship{},
			expectedErr:       sa.ErrBadRequest{Message: "scholarship funding end before funding start"},
		},
		"failed create scholarship": {
			paramCtx:         ctxValid,
			paramScholarship: scholarship,
			createScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctxValid, scholarship},
				Output:   []interface{}{sa.Scholarship{}, errors.New("error")},
			},
			expectedResp: sa.Scholarship{},
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
			expectedResp: scholarshipResp,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			scholarshipRepoMock := new(mocks.ScholarshipRepository)

			if test.createScholarship.IsCalled {
				scholarshipRepoMock.On("Create", test.createScholarship.Input...).
					Return(test.createScholarship.Output...).
					Once()
			}

			scholarshipService := _service.NewScholarshipService(scholarshipRepoMock)
			scholarshipResp, err := scholarshipService.Create(test.paramCtx, test.paramScholarship)
			scholarshipRepoMock.AssertExpectations(t)

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
		scholarships = make([]sa.Scholarship, 0)
		sponsor      sa.User
	)

	testdata.GoldenJSONUnmarshal(t, "scholarships", &scholarships)
	testdata.GoldenJSONUnmarshal(t, "user", &sponsor)

	scholarships[0].SponsorID = 1
	scholarships[1].SponsorID = 1

	tests := map[string]struct {
		paramFilter            sa.ScholarshipFilter
		fetchScholarship       testdata.FuncCaller
		fetchScholarshipCursor testdata.FuncCaller
		expectedResp           sa.ScholarshipFeed
		expectedErr            error
	}{
		"error fetch scholarship": {
			paramFilter: sa.ScholarshipFilter{},
			fetchScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.ScholarshipFilter{}},
				Output:   []interface{}{nil, "", errors.New("error")},
			},
			fetchScholarshipCursor: testdata.FuncCaller{},
			expectedResp:           sa.ScholarshipFeed{},
			expectedErr:            errors.New("error"),
		},
		"error fetch scholarship for cursor": {
			paramFilter: sa.ScholarshipFilter{},
			fetchScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.ScholarshipFilter{}},
				Output:   []interface{}{scholarships, cursor, nil},
			},
			fetchScholarshipCursor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.ScholarshipFilter{Cursor: cursor, Limit: 1}},
				Output:   []interface{}{nil, "", errors.New("error")},
			},
			expectedResp: sa.ScholarshipFeed{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramFilter: sa.ScholarshipFilter{},
			fetchScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.ScholarshipFilter{}},
				Output:   []interface{}{scholarships, cursor, nil},
			},
			fetchScholarshipCursor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.ScholarshipFilter{Cursor: cursor, Limit: 1}},
				Output:   []interface{}{[]sa.Scholarship{scholarships[0]}, cursor, nil},
			},
			expectedResp: sa.ScholarshipFeed{Cursor: cursor, Scholarships: scholarships},
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			scholarshipRepoMock := new(mocks.ScholarshipRepository)

			if test.fetchScholarship.IsCalled {
				scholarshipRepoMock.On("Fetch", test.fetchScholarship.Input...).
					Return(test.fetchScholarship.Output...).
					Once()
			}

			if test.fetchScholarshipCursor.IsCalled {
				scholarshipRepoMock.On("Fetch", test.fetchScholarshipCursor.Input...).
					Return(test.fetchScholarshipCursor.Output...).
					Once()
			}

			scholarshipService := _service.NewScholarshipService(scholarshipRepoMock)
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
