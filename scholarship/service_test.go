package scholarship_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"

	"testing"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	_service "github.com/Nusantara-Muda/scholarship-api/scholarship"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestScholarshipServiceCreate(t *testing.T) {
	var (
		scholarship sa.Scholarship
		user        sa.User
	)

	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "user", &user)

	scholarship.Sponsor = user
	scholarship.SponsorID = user.ID

	ctxValid := sa.SetUserOnContext(context.Background(), user)
	scholarshipNotMatch := scholarship
	scholarshipNotMatch.SponsorID = 2

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
				Output:   []interface{}{scholarship, nil},
			},
			expectedResp: scholarship,
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
