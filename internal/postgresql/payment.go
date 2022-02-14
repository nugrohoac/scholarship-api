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

// SubmitTransfer .
func (p paymentRepo) SubmitTransfer(ctx context.Context, payment sa.Payment) (sa.Payment, error) {
	var (
		errRollback, errCommit error
	)

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return sa.Payment{}, err
	}

	byteImage, err := json.Marshal(payment.Image)
	if err != nil {
		return sa.Payment{}, err
	}

	query, args, err := sq.Update("payment").
		SetMap(sq.Eq{
			"bank_account_name": payment.BankAccountName,
			"transfer_date":     payment.TransferDate,
			"image":             byteImage,
		}).Where(sq.Eq{"scholarship_id": payment.ScholarshipID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return sa.Payment{}, err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return sa.Payment{}, err
	}

	query, args, err = sq.Update("scholarship").
		SetMap(sq.Eq{"status": 1}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return sa.Payment{}, err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return sa.Payment{}, err
	}

	if errCommit = tx.Commit(); errCommit != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return sa.Payment{}, errCommit
	}

	return payment, nil
}

// NewPaymentRepository ...
func NewPaymentRepository(db *sql.DB) sa.PaymentRepository {
	return paymentRepo{db: db}
}
