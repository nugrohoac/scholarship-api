package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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

	if len(filter.IDs) > 0 {
		qSelect = qSelect.Where(sq.Eq{"id": filter.IDs})
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

// Login is use just for login
// Different select columns with fetch
func (u userRepo) Login(ctx context.Context, email string) (sa.User, error) {
	query, args, err := sq.Select("id",
		"name",
		"type",
		"email",
		"status",
		"password",
	).From("\"user\"").
		Where(sq.Eq{"email": email}).
		OrderBy("created_at desc").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return sa.User{}, err
	}

	row := u.db.QueryRowContext(ctx, query, args...)
	var user sa.User
	if err = row.Scan(&user.ID,
		&user.Name,
		&user.Type,
		&user.Email,
		&user.Status,
		&user.Password,
	); err != nil {
		return sa.User{}, err
	}

	return user, nil
}

// UpdateByID ...
// Update at table user
// insert into table card identity
// use transaction !!!!!!!!!!!
func (u userRepo) UpdateByID(ctx context.Context, ID int64, user sa.User) (sa.User, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return sa.User{}, err
	}

	var (
		timeNow     = time.Now()
		bytesImg    []byte
		errRollback error
	)

	bytesImg, err = json.Marshal(user.Photo)
	if err != nil {
		return sa.User{}, err
	}

	query, args, err := sq.Update("\"user\"").
		SetMap(sq.Eq{
			"name":              user.Name,
			"photo":             bytesImg,
			"company_name":      user.CompanyName,
			"country_id":        user.CountryID,
			"address":           user.Address,
			"postal_code":       user.PostalCode,
			"bank_id":           user.BankID,
			"bank_account_no":   user.BankAccountNo,
			"bank_account_name": user.BankAccountName,
			"status":            user.Status,
			"updated_at":        timeNow,
		}).Where(sq.Eq{"id": ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return sa.User{}, err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		return sa.User{}, err
	}

	qInsert := sq.Insert("card_identity").
		Columns("type",
			"no",
			"image",
			"user_id",
			"created_at",
		)
	for _, cardIdentity := range user.CardIdentities {
		bytesImg, err = json.Marshal(cardIdentity.Image)
		if err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				fmt.Println("Err rollback update profile at json marshal image : ", errRollback)
			}

			return sa.User{}, err
		}

		qInsert = qInsert.Values(
			cardIdentity.Type,
			cardIdentity.No,
			bytesImg,
			cardIdentity.UserID,
			timeNow,
		)
	}

	query, args, err = qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			fmt.Println("Err rollback update profile generate query insert card identity : ", errRollback)
		}

		return sa.User{}, err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			fmt.Println("Err rollback update profile exec insert card identity : ", errRollback)
		}

		return sa.User{}, err
	}

	if errCommit := tx.Commit(); errCommit != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			fmt.Println("Err rollback update profile at commit : ", errRollback)
		}
	}

	return user, nil
}

// SetStatus ....
// 1 active
// 0 inactive
func (u userRepo) SetStatus(ctx context.Context, ID int64, status int) error {
	query, args, err := sq.Update("\"user\"").SetMap(sq.Eq{
		"status":     status,
		"updated_at": time.Now(),
	}).Where(sq.Eq{"id": ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = u.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// ResetPassword ...
func (u userRepo) ResetPassword(ctx context.Context, email, password string) error {
	query, args, err := sq.Update("\"user\"").
		SetMap(sq.Eq{
			"password":   password,
			"updated_at": time.Now(),
		}).Where(sq.Eq{"email": email}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = u.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// NewUserRepository .
func NewUserRepository(db *sql.DB) sa.UserRepository {
	return userRepo{db: db}
}
