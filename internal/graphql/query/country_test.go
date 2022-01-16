package query_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestCountryQuery(t *testing.T) {
	countries := make([]sa.Country, 0)
	testdata.GoldenJSONUnmarshal(t, "countries", &countries)

	var (
		limit  int32 = 10
		cursor       = "cursor"
	)

	countryFeed := sa.CountryFeed{
		Cursor:    cursor,
		Countries: countries,
	}

	countryFeedResolver := resolver.CountryFeedResolver{
		CountryFeed: countryFeed,
	}

	tests := map[string]struct {
		paramFilter  sa.InputCountryFilter
		fetchCountry testdata.FuncCaller
		expectedResp *resolver.CountryFeedResolver
		expectedErr  error
	}{
		"success": {
			paramFilter: sa.InputCountryFilter{
				Limit:  &limit,
				Cursor: &cursor,
			},
			fetchCountry: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.CountryFilter{Limit: int(limit), Cursor: cursor}},
				Output:   []interface{}{countryFeed, nil},
			},
			expectedResp: &countryFeedResolver,
			expectedErr:  nil,
		},
		"error": {
			paramFilter: sa.InputCountryFilter{
				Limit:  &limit,
				Cursor: &cursor,
			},
			fetchCountry: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.CountryFilter{Limit: int(limit), Cursor: cursor}},
				Output:   []interface{}{sa.CountryFeed{}, errors.New("internal server error")},
			},
			expectedResp: &countryFeedResolver,
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			countryServiceMock := new(mocks.CountryService)

			if test.fetchCountry.IsCalled {
				countryServiceMock.On("Fetch", test.fetchCountry.Input...).
					Return(test.fetchCountry.Output...).
					Once()
			}

			countryQuery := query.NewCountryQuery(countryServiceMock)
			countryFeedResp, err := countryQuery.FetchCountry(context.Background(), test.paramFilter)
			countryServiceMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, countryFeedResp)
		})
	}
}
