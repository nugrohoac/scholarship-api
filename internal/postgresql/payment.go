package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type paymentRepo struct {
	db *sql.DB
}

// Fetch .
func (p paymentRepo) Fetch(ctx context.Context, scholarshipIDs []int64) ([]sa.Payment, error) {
	qSelect := sq.Select("id",
		"scholarship_id",
		"deadline",
		"transfer_date",
		"bank_account_name",
		"image",
		"created_at",
	).From("payment").
		OrderBy("created_at desc")

	if len(scholarshipIDs) > 0 {
		qSelect = qSelect.Where(sq.Eq{"scholarship_id": scholarshipIDs})
	}

	query, args, err := qSelect.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errCloser := rows.Close(); errCloser != nil {
			logrus.Error("error close row payment : ", errCloser)
		}
	}()

	payments := make([]sa.Payment, 0)

	for rows.Next() {
		var (
			payment    sa.Payment
			bytesImage []byte
		)

		if err = rows.Scan(&payment.ID,
			&payment.ScholarshipID,
			&payment.Deadline,
			&payment.TransferDate,
			&payment.BankAccountName,
			&bytesImage,
			&payment.CreatedAt,
		); err != nil {
			return nil, err
		}

		if bytesImage != nil {
			if err = json.Unmarshal(bytesImage, &payment.Image); err != nil {
				return nil, err
			}
		}

		payments = append(payments, payment)
	}

	return payments, nil
}

// NewPaymentRepository ...
func NewPaymentRepository(db *sql.DB) sa.PaymentRepository {
	return paymentRepo{db: db}
}
