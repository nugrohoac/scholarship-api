package major_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/major"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestMajorService_Fetch(t *testing.T) {
	var (
		majors    []sa.Major
		majorFeed sa.MajorFeed
		cursor    = "cursor"
	)

	testdata.GoldenJSONUnmarshal(t, "majors", &majors)
	majorFeed.Cursor = cursor
	majorFeed.Majors = majors

	tests := map[string]struct {
		paramFilter  sa.MajorFilter
		fetchMajor   testdata.FuncCaller
		expectedResp sa.MajorFeed
		expectedErr  error
	}{
		"error": {
			paramFilter: sa.MajorFilter{},
			fetchMajor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.MajorFilter{}},
				Output:   []interface{}{nil, "", errors.New("error")},
			},
			expectedResp: sa.MajorFeed{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramFilter: sa.MajorFilter{},
			fetchMajor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.MajorFilter{}},
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
