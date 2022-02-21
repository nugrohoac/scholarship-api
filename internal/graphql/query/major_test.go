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

func TestMajorQuery_GetMajor(t *testing.T) {
	var (
		majors = make([]sa.Major, 0)
		cursor = "cursor"
	)

	testdata.GoldenJSONUnmarshal(t, "majors", &majors)
	majorFeedResolver := resolver.MajorFeedResolver{
		MajorFeed: sa.MajorFeed{
			Cursor: cursor,
			Majors: majors,
		},
	}

	tests := map[string]struct {
		paramFilter  sa.InputMajorFilter
		fetchMajor   testdata.FuncCaller
		expectedResp *resolver.MajorFeedResolver
		expectedErr  error
	}{
		"error": {
			paramFilter: sa.InputMajorFilter{},
			fetchMajor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.MajorFilter{}},
				Output:   []interface{}{sa.MajorFeed{}, errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramFilter: sa.InputMajorFilter{},
			fetchMajor: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.MajorFilter{}},
				Output:   []interface{}{sa.MajorFeed{Cursor: cursor, Majors: majors}, nil},
			},
			expectedResp: &majorFeedResolver,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			majorServiceMock := new(mocks.MajorService)

			if test.fetchMajor.IsCalled {
				majorServiceMock.On("Fetch", test.fetchMajor.Input...).
					Return(test.fetchMajor.Output...).
					Once()
			}

			majorQuery := query.NewMajorQuery(majorServiceMock)
			_response, err := majorQuery.FetchMajor(context.Background(), test.paramFilter)
			majorServiceMock.AssertExpectations(t)

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
