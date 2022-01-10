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

type bankSuite struct {
	postgresql.TestSuite
}

func TestBankRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test bank")
	}

	suite.Run(t, new(bankSuite))
}

func (b bankSuite) TestBankRepoFetch() {
	banks := make([]sa.Bank, 0)
	testdata.GoldenJSONUnmarshal(b.T(), "banks", &banks)
	postgresql.SeedBanks(b.DBConn, b.T(), banks)

	// normal case
	bankRepo := postgresql.NewBankRepository(b.DBConn)
	bankResp, cursor, err := bankRepo.Fetch(context.Background(), sa.BankFilter{})
	require.NoError(b.T(), err)
	require.Equal(b.T(), "MjAyMi0wMS0wN1QyMjo1NTozMy4wODda", cursor)
	require.Equal(b.T(), 3, len(bankResp))

	// case with limit
	bankResp, cursor, err = bankRepo.Fetch(context.Background(), sa.BankFilter{Limit: 1})
	require.NoError(b.T(), err)
	require.Equal(b.T(), "MjAyMi0wMS0wN1QyMjo1NTozNS4wODda", cursor)
	require.Equal(b.T(), 1, len(bankResp))

	// case with limit and cursor
	bankResp, cursor, err = bankRepo.Fetch(context.Background(), sa.BankFilter{
		Limit:  1,
		Cursor: "MjAyMi0wMS0wN1QyMjo1NTozNS4wODda",
	})

	require.NoError(b.T(), err)
	require.Equal(b.T(), "MjAyMi0wMS0wN1QyMjo1NTozNC4wODda", cursor)
	require.Equal(b.T(), 1, len(bankResp))
	require.Equal(b.T(), "Mandiri", bankResp[0].Name)

	// case with name
	bankResp, cursor, err = bankRepo.Fetch(context.Background(), sa.BankFilter{
		Name: "bca",
	})

	require.NoError(b.T(), err)
	require.Equal(b.T(), 1, len(bankResp))
	require.Equal(b.T(), "BCA", bankResp[0].Name)
	require.Equal(b.T(), "MjAyMi0wMS0wN1QyMjo1NTozMy4wODda", cursor)
}
