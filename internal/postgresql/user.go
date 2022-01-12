package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

type userRepo struct {
	db *sql.DB
}

// Store .
func (u userRepo) Store(ctx context.Context, user sa.User) (sa.User, error) {
	user.CreatedAt = time.Now()
	bytePhoto, err := json.Marshal(user.Photo)
	if err != nil {
		return sa.User{}, err
	}

	query, args, err := sq.Insert("\"user\"").
		Columns("name",
			"type",
			"email",
			"phone_no",
			"photo",
			"company_name",
			"password",
			"status",
			"country_id",
			"postal_code",
			"address",
			"gender",
			"ethnic",
			"bank_id",
			"bank_account_no",
			"bank_account_name",
			"created_at",
		).Values(user.Name,
		user.Type,
		user.Email,
		user.PhoneNo,
		bytePhoto,
		user.CompanyName,
		user.Password,
		user.Status,
		user.CountryID,
		user.PostalCode,
		user.Address,
		user.Gender,
		user.Ethnic,
		user.BankID,
		user.BankAccountNo,
		user.BankAccountName,
		user.CreatedAt,
	).PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return sa.User{}, err
	}

	_, err = u.db.ExecContext(ctx, query, args...)
	if err != nil {
		return sa.User{}, err
	}

	return user, nil
}

// NewUserRepository .
func NewUserRepository(db *sql.DB) sa.UserRepository {
	return userRepo{db: db}
}
