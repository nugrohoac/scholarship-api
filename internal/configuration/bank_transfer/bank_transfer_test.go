package bank_transfer_test

import (
	"testing"

	"github.com/Nusantara-Muda/scholarship-api/internal/configuration/bank_transfer"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/require"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

func TestBankTransferRepo_Get(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		var bankTransfer sa.BankTransfer
		testdata.GoldenJSONUnmarshal(t, "bank_transfer", &bankTransfer)

		bankTransferRepo := bank_transfer.NewBankTransfer(bankTransfer)
		response := bankTransferRepo.Get()
		require.Equal(t, response, bankTransfer)
	})
}
