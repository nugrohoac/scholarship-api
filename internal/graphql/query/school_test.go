package query_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestSchoolQuery_FetchSchool(t *testing.T) {
	var schools []sa.School
	testdata.GoldenJSONUnmarshal(t, "schools", &schools)

	schoolFeed := sa.SchoolFeed{
		Cursor:  "cursor",
		Schools: schools,
	}

	tests := map[string]struct {
		paramFilter  sa.InputSchoolFilter
		fetchSchool  testdata.FuncCaller
		expectedResp *resolver.SchoolFeedResolver
		expectedErr  error
	}{
		"error": {
			paramFilter: sa.InputSchoolFilter{},
			fetchSchool: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.SchoolFilter{}},
				Output:   []interface{}{sa.SchoolFeed{}, errors.New("school")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("school"),
		},
		"success": {
			paramFilter: sa.InputSchoolFilter{},
			fetchSchool: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.SchoolFilter{}},
				Output:   []interface{}{schoolFeed, nil},
			},
			expectedResp: &resolver.SchoolFeedResolver{SchoolFeed: schoolFeed},
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			schoolServiceMock := new(mocks.SchoolService)

			if test.fetchSchool.IsCalled {
				schoolServiceMock.On("Fetch", test.fetchSchool.Input...).
					Return(test.fetchSchool.Output...).
					Once()
			}

			schoolQuery := query.NewSchoolQuery(schoolServiceMock)
			response, err := schoolQuery.FetchSchool(context.Background(), test.paramFilter)
			schoolServiceMock.AssertExpectations(t)

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
