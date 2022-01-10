package postgresql

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"

	sq "github.com/Masterminds/squirrel"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type bankRepo struct {
	db *sql.DB
}

// Fetch ...
func (b *bankRepo) Fetch(ctx context.Context, filter sa.BankFilter) ([]sa.Bank, string, error) {
	qSelect := sq.Select("id",
		"name",
		"code",
		"created_at").
		From("bank").
		OrderBy("created_at DESC")

	if filter.Limit > 0 {
		qSelect = qSelect.Limit(uint64(filter.Limit))
	}

	if filter.Cursor != "" {
		cursorTime, err := decodeCursor(filter.Cursor)
		if err != nil {
			return nil, "", err
		}

		qSelect = qSelect.Where(sq.Lt{"created_at": cursorTime})
	}

	if filter.Name != "" {
		name := "%" + filter.Name + "%"

		qSelect = qSelect.Where(sq.Like{"LOWER(name)": name})
	}

	query, args, err := qSelect.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		banks  = make([]sa.Bank, 0)
		cursor = time.Time{}
	)

	for rows.Next() {
		var bank sa.Bank

		if err = rows.Scan(
			&bank.ID,
			&bank.Name,
			&bank.Code,
			&bank.CreatedAt,
		); err != nil {
			return nil, "", err
		}

		cursor = bank.CreatedAt
		banks = append(banks, bank)
	}

	cursorStr, err := encodeCursor(cursor)
	if err != nil {
		return nil, "", err
	}

	return banks, cursorStr, nil
}

// NewBankRepository ...
func NewBankRepository(db *sql.DB) sa.BankRepository {
	return &bankRepo{db: db}
}
