package postgresql_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/suite"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

type paymentSuite struct {
	postgresql.TestSuite
}

func TestPaymentRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test payment")
	}

	suite.Run(t, new(paymentSuite))
}

func (p paymentSuite) TestPaymentRepo_Fetch() {
	t := p.T()
	var payments []sa.Payment
	testdata.GoldenJSONUnmarshal(t, "payments", &payments)
	postgresql.SeedPayments(p.DBConn, t, payments)

	paymentRepo := postgresql.NewPaymentRepository(p.DBConn)
	paymentsResp, err := paymentRepo.Fetch(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, 2, len(paymentsResp))
	require.Equal(t, payments[1].ID, paymentsResp[0].ID)
	require.Equal(t, payments[0].ID, paymentsResp[1].ID)

	paymentsResp, err = paymentRepo.Fetch(context.Background(), []int64{payments[0].ID})
	require.NoError(t, err)
	require.Equal(t, 1, len(paymentsResp))
	require.Equal(t, payments[0].ID, paymentsResp[0].ID)
}
