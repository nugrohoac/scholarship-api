package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/sirupsen/logrus"
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

	query, args, err := sq.Insert("\"user\"").
		Columns("type",
			"email",
			"phone_no",
			"password",
			"status",
			"created_at",
		).Values(user.Type,
		user.Email,
		user.PhoneNo,
		user.Password,
		user.Status,
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

// Fetch ...
func (u userRepo) Fetch(ctx context.Context, filter sa.UserFilter) ([]sa.User, string, error) {
	qSelect := sq.Select("id",
		"name",
		"type",
		"email",
		"phone_no",
		"photo",
		"company_name",
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
	).From("\"user\"").
		OrderBy("created_at desc").
		PlaceholderFormat(sq.Dollar)

	if filter.Email != "" {
		qSelect = qSelect.Where(sq.Eq{"email": filter.Email})
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := u.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		users     = make([]sa.User, 0)
		cursor    = time.Time{}
		bytePhoto []byte
	)

	for rows.Next() {
		var user sa.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Type,
			&user.Email,
			&user.PhoneNo,
			&bytePhoto,
			&user.CompanyName,
			&user.Status,
			&user.CountryID,
			&user.PostalCode,
			&user.Address,
			&user.Gender,
			&user.Ethnic,
			&user.BankID,
			&user.BankAccountNo,
			&user.BankAccountName,
			&user.CreatedAt,
		); err != nil {
			return nil, "", err
		}

		cursor = user.CreatedAt
		if bytePhoto != nil {
			if err = json.Unmarshal(bytePhoto, &user.Photo); err != nil {
				return nil, "", err
			}
		}

		users = append(users, user)
	}

	cursorStr, err := encodeCursor(cursor)
	if err != nil {
		return nil, "", err
	}

	return users, cursorStr, nil
}

// NewUserRepository .
func NewUserRepository(db *sql.DB) sa.UserRepository {
	return userRepo{db: db}
}
