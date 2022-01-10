package country_test

import (
	"context"
	"errors"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/country"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var cursor = "next-cursor"

func TestCountryServiceFetch(t *testing.T) {

	countries := make([]sa.Country, 0)
	testdata.GoldenJSONUnmarshal(t, "countries", &countries)

	tests := map[string]struct {
		paramFilter  sa.CountryFilter
		fetchCountry testdata.FuncCaller
		expectedResp sa.CountryFeed
		expectedErr  error
	}{
		"success": {
			paramFilter: sa.CountryFilter{Limit: 10},
			fetchCountry: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.CountryFilter{Limit: 10}},
				Output:   []interface{}{countries, cursor, nil},
			},
			expectedResp: sa.CountryFeed{
				Cursor:    cursor,
				Countries: countries,
			},
			expectedErr: nil,
		},
		"error": {
			paramFilter: sa.CountryFilter{},
			fetchCountry: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.CountryFilter{}},
				Output:   []interface{}{nil, "", errors.New("internal server error")},
			},
			expectedResp: sa.CountryFeed{},
			expectedErr:  errors.New("internal server error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			countryRepoMock := new(mocks.CountryRepository)

			if test.fetchCountry.IsCalled {
				countryRepoMock.On("Fetch", test.fetchCountry.Input...).
					Return(test.fetchCountry.Output...).
					Once()
			}

			countryService := country.NewCountryService(countryRepoMock)
			countryResp, err := countryService.Fetch(context.Background(), test.paramFilter)
			countryRepoMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, countryResp)
		})
	}
}
