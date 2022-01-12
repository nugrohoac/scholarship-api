package postgresql_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

type countrySuite struct {
	postgresql.TestSuite
}

func TestCountryRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test country")
	}

	suite.Run(t, new(countrySuite))
}

func (c countrySuite) TestCountryRepoFetch() {
	countries := make([]sa.Country, 0)
	testdata.GoldenJSONUnmarshal(c.T(), "countries", &countries)
	postgresql.SeedCountries(c.DBConn, c.T(), countries)

	// without filter
	countryRepo := postgresql.NewCountryRepository(c.DBConn)
	countriesResp, cursor, err := countryRepo.Fetch(context.Background(), sa.CountryFilter{})
	require.NoError(c.T(), err)
	require.Equal(c.T(), "MjAyMi0wMS0xMVQwNTozNDowMy4xOTFa", cursor)
	require.Equal(c.T(), 3, len(countriesResp))

	// limit
	countriesResp, cursor, err = countryRepo.Fetch(context.Background(), sa.CountryFilter{Limit: 1})
	require.NoError(c.T(), err)
	require.Equal(c.T(), "MjAyMi0wMS0xMVQwNTozNDowNS4xOTFa", cursor)
	require.Equal(c.T(), 1, len(countriesResp))
	require.Equal(c.T(), "Singapura", countriesResp[0].Name)

	// limit and cursor
	countriesResp, cursor, err = countryRepo.Fetch(context.Background(), sa.CountryFilter{
		Limit:  1,
		Cursor: "MjAyMi0wMS0xMVQwNTozNDowNS4xOTFa",
	})
	require.NoError(c.T(), err)
	require.Equal(c.T(), "MjAyMi0wMS0xMVQwNTozNDowNC4xOTFa", cursor)
	require.Equal(c.T(), 1, len(countriesResp))
	require.Equal(c.T(), "Malaysia", countriesResp[0].Name)

	// name
	countriesResp, cursor, err = countryRepo.Fetch(context.Background(), sa.CountryFilter{
		Name: "indo",
	})
	require.NoError(c.T(), err)
	require.Equal(c.T(), "MjAyMi0wMS0xMVQwNTozNDowMy4xOTFa", cursor)
	require.Equal(c.T(), 1, len(countriesResp))
	require.Equal(c.T(), "Indonesia", countriesResp[0].Name)
}
