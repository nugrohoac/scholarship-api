package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

type bankTransferRepo struct {
	db *sql.DB
}

// Get ...
func (b bankTransferRepo) Get(ctx context.Context) (sa.BankTransfer, error) {
	query, args, err := sq.Select("id",
		"name",
		"account_name",
		"account_no",
		"image",
	).From("bank_transfer").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return sa.BankTransfer{}, err
	}

	var (
		bankTransfer sa.BankTransfer
		byteImage    []byte
	)

	row := b.db.QueryRowContext(ctx, query, args...)
	if err = row.Scan(
		&bankTransfer.ID,
		&bankTransfer.Name,
		&bankTransfer.AccountName,
		&bankTransfer.AccountNo,
		&byteImage,
	); err != nil {
		if err == sql.ErrNoRows {
			return sa.BankTransfer{}, sa.ErrNotFound{Message: "bank transfer is not found"}
		}

		return sa.BankTransfer{}, err
	}

	if byteImage != nil {
		if err = json.Unmarshal(byteImage, &bankTransfer.Image); err != nil {
			return sa.BankTransfer{}, err
		}
	}

	return bankTransfer, nil
}

// NewBankTransferRepository ...
func NewBankTransferRepository(db *sql.DB) sa.BankTransferRepository {
	return bankTransferRepo{
		db: db,
	}
}
