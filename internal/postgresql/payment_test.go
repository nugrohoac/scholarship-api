package postgresql_test

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

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

func (p paymentSuite) TestPaymentFetch() {
	t := p.T()
	var payments []entity.Payment
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

func (p paymentSuite) TestPaymentSubmitTransfer() {
	var (
		scholarship entity.Scholarship
		payment     entity.Payment
		t           = p.T()
	)

	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "payment", &payment)

	postgresql.SeedScholarship(p.DBConn, t, []entity.Scholarship{scholarship})

	payment.ScholarshipID = scholarship.ID

	oldPayment := payment
	oldPayment.TransferDate = time.Time{}
	oldPayment.BankAccountName = ""
	oldPayment.Image = entity.Image{}
	postgresql.SeedPayments(p.DBConn, t, []entity.Payment{oldPayment})

	paymentRepo := postgresql.NewPaymentRepository(p.DBConn)
	paymentResp, err := paymentRepo.SubmitTransfer(context.Background(), payment)
	require.NoError(t, err)

	payments, err := paymentRepo.Fetch(context.Background(), []int64{payment.ID})
	require.NoError(t, err)
	require.Len(t, payments, 1)
	require.Equal(t, paymentResp.BankAccountName, payments[0].BankAccountName)
	require.Equal(t, paymentResp.Image, payments[0].Image)

	paymentRespTransferDate := paymentResp.TransferDate.Format("2006-01-02T15:04:05")
	currentTransferDate := payments[0].TransferDate.Format("2006-01-02T15:04:05")
	require.Equal(t, paymentRespTransferDate, currentTransferDate)

	var status int
	row := p.DBConn.QueryRow("SELECT status FROM scholarship WHERE id=$1", payment.ScholarshipID)
	err = row.Scan(&status)
	require.NoError(t, err)
	require.Equal(t, 1, status)
}
