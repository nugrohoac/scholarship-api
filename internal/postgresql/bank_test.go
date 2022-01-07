package postgresql_test

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
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

	bankRepo := postgresql.NewBankRepository(b.DBConn)
	bankResp, cursor, err := bankRepo.Fetch(context.Background(), sa.BankFilter{})
	require.NoError(b.T(), err)
	require.NotEmpty(b.T(), cursor)
	require.Equal(b.T(), 3, len(bankResp))
}
