package query_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/stretchr/testify/require"
	"testing"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
)

var cursor = "next-cursor"

func TestScholarshipQuery_GetScholarshipBySponsor(t *testing.T) {
	var (
		sponsor      sa.User
		scholarships = make([]sa.Scholarship, 0)
	)

	testdata.GoldenJSONUnmarshal(t, "user", &sponsor)
	testdata.GoldenJSONUnmarshal(t, "scholarships", &scholarships)

	scholarships[0].SponsorID = sponsor.ID
	scholarships[1].SponsorID = sponsor.ID

	scholarshipFeedResolver := resolver.ScholarshipFeedResolver{
		ScholarshipFeed: struct {
			Cursor       string
			Scholarships []sa.Scholarship
		}{Cursor: cursor, Scholarships: scholarships},
	}

	tests := map[string]struct {
		param            sa.InputScholarshipFilter
		fetchScholarship testdata.FuncCaller
		expectedResp     *resolver.ScholarshipFeedResolver
		expectedErr      error
	}{
		"success": {
			param: sa.InputScholarshipFilter{},
			fetchScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.ScholarshipFilter{}},
				Output:   []interface{}{sa.ScholarshipFeed{Cursor: cursor, Scholarships: scholarships}, nil},
			},
			expectedResp: &scholarshipFeedResolver,
			expectedErr:  nil,
		},
		"error": {
			param: sa.InputScholarshipFilter{},
			fetchScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.ScholarshipFilter{}},
				Output:   []interface{}{sa.ScholarshipFeed{}, errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			scholarshipServiceMock := new(mocks.ScholarshipService)

			if test.fetchScholarship.IsCalled {
				scholarshipServiceMock.On("Fetch", test.fetchScholarship.Input...).
					Return(test.fetchScholarship.Output...).
					Once()
			}

			scholarshipQuery := query.NewScholarshipQuery(scholarshipServiceMock)
			resp, err := scholarshipQuery.FetchScholarship(context.Background(), test.param)
			scholarshipServiceMock.AssertExpectations(t)

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

func TestScholarshipQuery_GetScholarshipByID(t *testing.T) {
	var (
		scholarship  sa.Scholarship
		payment      sa.Payment
		bankTransfer sa.BankTransfer
	)

	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "payment", &payment)
	testdata.GoldenJSONUnmarshal(t, "bank_transfer", &bankTransfer)

	scholarship.Payment = payment
	scholarship.Payment.BankTransfer = bankTransfer

	response := resolver.ScholarshipResolver{Scholarship: scholarship}

	tests := map[string]struct {
		paramID            struct{ ID int32 }
		getScholarshipByID testdata.FuncCaller
		expectedResp       *resolver.ScholarshipResolver
		expectedErr        error
	}{
		"error": {
			paramID: struct {
				ID int32
			}{ID: int32(scholarship.ID)},
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarship.ID},
				Output:   []interface{}{sa.Scholarship{}, errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramID: struct {
				ID int32
			}{ID: int32(scholarship.ID)},
			getScholarshipByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarship.ID},
				Output:   []interface{}{scholarship, nil},
			},
			expectedResp: &response,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			scholarshipServiceMock := new(mocks.ScholarshipService)

			if test.getScholarshipByID.IsCalled {
				scholarshipServiceMock.On("GetByID", test.getScholarshipByID.Input...).
					Return(test.getScholarshipByID.Output...).
					Once()
			}

			scholarshipQuery := query.NewScholarshipQuery(scholarshipServiceMock)
			scholarshipResp, err := scholarshipQuery.GetScholarshipByID(context.Background(), test.paramID)
			scholarshipServiceMock.AssertExpectations(t)

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
