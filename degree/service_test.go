package degree_test

import (
	"context"
	"errors"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/degree"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDegreeServiceGet(t *testing.T) {
	var degrees []sa.Degree
	testdata.GoldenJSONUnmarshal(t, "degrees", &degrees)

	tests := map[string]struct {
		fetchDegree  testdata.FuncCaller
		expectedResp []sa.Degree
		expectedErr  error
	}{
		"error": {
			fetchDegree: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything},
				Output:   []interface{}{nil, errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
		"success": {
			fetchDegree: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything},
				Output:   []interface{}{degrees, nil},
			},
			expectedResp: degrees,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			degreeRepoMock := new(mocks.DegreeRepository)

			if test.fetchDegree.IsCalled {
				degreeRepoMock.On("Fetch", test.fetchDegree.Input...).
					Return(test.fetchDegree.Output...).
					Once()
			}

			degreeService := degree.NewDegreeService(degreeRepoMock)
			response, err := degreeService.Fetch(context.Background())
			degreeRepoMock.AssertExpectations(t)

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
