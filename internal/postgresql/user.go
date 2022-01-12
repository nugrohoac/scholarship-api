package postgresql

import (
	"context"
	"database/sql"
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

// NewUserRepository .
func NewUserRepository(db *sql.DB) sa.UserRepository {
	return userRepo{db: db}
}
