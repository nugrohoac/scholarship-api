package postgresql_test

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/require"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type bankTransferSuite struct {
	postgresql.TestSuite
}

func TestBankTransferRepo(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test country")
	}

	suite.Run(t, new(bankTransferSuite))
}

func (b bankTransferSuite) TestBankTransferRepo_Get() {
	var (
		bankTransfer sa.BankTransfer
		t            = b.T()
	)
	testdata.GoldenJSONUnmarshal(t, "bank_transfer", &bankTransfer)
	postgresql.SeedBankTransfer(b.DBConn, t, bankTransfer)

	bankTransferRepo := postgresql.NewBankTransferRepository(b.DBConn)
	response, err := bankTransferRepo.Get(context.Background())
	require.NoError(t, err)
	require.Equal(t, bankTransfer.Name, response.Name)
	require.Equal(t, bankTransfer.AccountName, response.AccountName)
	require.Equal(t, bankTransfer.AccountNo, response.AccountNo)
	require.Equal(t, bankTransfer.Image, response.Image)
}
