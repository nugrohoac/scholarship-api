package country_test

import (
	"context"
	"errors"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/country"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var cursor = "next-cursor"

func TestCountryServiceFetch(t *testing.T) {

	countries := make([]entity.Country, 0)
	testdata.GoldenJSONUnmarshal(t, "countries", &countries)

	tests := map[string]struct {
		paramFilter  entity.CountryFilter
		fetchCountry testdata.FuncCaller
		expectedResp entity.CountryFeed
		expectedErr  error
	}{
		"success": {
			paramFilter: entity.CountryFilter{Limit: 10},
			fetchCountry: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.CountryFilter{Limit: 10}},
				Output:   []interface{}{countries, cursor, nil},
			},
			expectedResp: entity.CountryFeed{
				Cursor:    cursor,
				Countries: countries,
			},
			expectedErr: nil,
		},
		"error": {
			paramFilter: entity.CountryFilter{},
			fetchCountry: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, entity.CountryFilter{}},
				Output:   []interface{}{nil, "", errors.New("internal server error")},
			},
			expectedResp: entity.CountryFeed{},
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
