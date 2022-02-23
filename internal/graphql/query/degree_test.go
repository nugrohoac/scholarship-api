package query_test

import (
	"context"
	"errors"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDegreeQuery_GetDegree(t *testing.T) {
	var degrees []sa.Degree
	testdata.GoldenJSONUnmarshal(t, "degrees", &degrees)

	response := make([]*resolver.DegreeResolver, 0)

	for _, degree := range degrees {
		degree := degree
		response = append(response, &resolver.DegreeResolver{Degree: degree})
	}

	tests := map[string]struct {
		fetchDegree  testdata.FuncCaller
		expectedResp *[]*resolver.DegreeResolver
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
			expectedResp: &response,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			degreeRepoMock := new(mocks.DegreeService)

			if test.fetchDegree.IsCalled {
				degreeRepoMock.On("Fetch", test.fetchDegree.Input...).
					Return(test.fetchDegree.Output...).
					Once()
			}

			degreeQuery := query.NewDegreeQuery(degreeRepoMock)
			_response, err := degreeQuery.FetchDegree(context.Background())
			degreeRepoMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, _response)
		})
	}
}
