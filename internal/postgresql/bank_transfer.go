package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"

	sq "github.com/Masterminds/squirrel"
)

type bankTransferRepo struct {
	db *sql.DB
}

// Get ...
func (b bankTransferRepo) Get(ctx context.Context) (entity.BankTransfer, error) {
	query, args, err := sq.Select("id",
		"name",
		"account_name",
		"account_no",
		"image",
	).From("bank_transfer").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return entity.BankTransfer{}, err
	}

	var (
		bankTransfer entity.BankTransfer
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
			return entity.BankTransfer{}, errors.ErrNotFound{Message: "bank transfer is not found"}
		}

		return entity.BankTransfer{}, err
	}

	if byteImage != nil {
		if err = json.Unmarshal(byteImage, &bankTransfer.Image); err != nil {
			return entity.BankTransfer{}, err
		}
	}

	return bankTransfer, nil
}

// NewBankTransferRepository ...
func NewBankTransferRepository(db *sql.DB) business.BankTransferRepository {
	return bankTransferRepo{
		db: db,
	}
}
