package postgresql

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

type schoolRepo struct {
	db *sql.DB
}

// Create .
func (s schoolRepo) Create(ctx context.Context, school sa.School) (sa.School, error) {
	school.Status = 1
	school.CreatedAt = time.Now()

	query, args, err := sq.Insert("school").
		Columns("name",
			"type",
			"address",
			"status",
			"created_at",
		).Values(school.Name,
		school.Type,
		school.Address,
		school.Status,
		school.CreatedAt,
	).PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return sa.School{}, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&school.ID); err != nil {
		return sa.School{}, err
	}

	return school, nil
}

// NewSchoolRepository .
func NewSchoolRepository(db *sql.DB) sa.SchoolRepository {
	return schoolRepo{db: db}
}
