package major_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/major"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestMajorService_Fetch(t *testing.T) {
	var (
		majors    []entity.Major
		majorFeed entity.MajorFeed
		cursor    = "cursor"
	)

	testdata.GoldenJSONUnmarshal(t, "majors", &majors)
	majorFeed.Cursor = cursor
	majorFeed.Majors = majors

	tests := map[string]struct {
		paramFilter  entity.MajorFilter
		fetchMajor   testdata.FuncCaller
		expectedResp entity.MajorFeed
		expectedErr  error
	}{
		"error": {
			paramFilter: entity.MajorFilter{},
			fetchMajor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.MajorFilter{}},
				Output:   []interface{}{nil, "", errors.New("error")},
			},
			expectedResp: entity.MajorFeed{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramFilter: entity.MajorFilter{},
			fetchMajor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.MajorFilter{}},
				Output:   []interface{}{majors, cursor, nil},
			},
			expectedResp: majorFeed,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			majorRepoMock := new(mocks.MajorRepository)

			if test.fetchMajor.IsCalled {
				majorRepoMock.On("Fetch", test.fetchMajor.Input...).
					Return(test.fetchMajor.Output...).
					Once()
			}

			majorService := major.NewMajorService(majorRepoMock)
			response, err := majorService.Fetch(context.Background(), test.paramFilter)
			majorRepoMock.AssertExpectations(t)

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
