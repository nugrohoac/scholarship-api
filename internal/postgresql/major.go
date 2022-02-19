package postgresql

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	sq "github.com/Masterminds/squirrel"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

type majorRepo struct {
	db *sql.DB
}

// Fetch .
func (m majorRepo) Fetch(ctx context.Context, filter sa.MajorFilter) ([]sa.Major, string, error) {
	qSelect := sq.Select("id",
		"name",
		"created_at",
	).From("major").
		OrderBy("created_at desc")

	if filter.Limit > 0 {
		qSelect = qSelect.Limit(filter.Limit)
	}

	if filter.Name != "" {
		name := "%" + strings.ToLower(filter.Name) + "%"
		qSelect = qSelect.Where(sq.Like{"LOWER(name)": name})
	}

	if filter.Cursor != "" {
		cursorTime, err := decodeCursor(filter.Cursor)
		if err != nil {
			return nil, "", err
		}

		qSelect = qSelect.Where(sq.Lt{"created_at": cursorTime})
	}

	query, args, err := qSelect.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		majors    = make([]sa.Major, 0)
		cursor    time.Time
		cursorStr string
	)

	for rows.Next() {
		var major sa.Major
		if err = rows.Scan(
			&major.ID,
			&major.Name,
			&major.CreatedAt,
		); err != nil {
			return nil, "", err
		}

		cursor = major.CreatedAt
		majors = append(majors, major)
	}

	if !cursor.IsZero() {
		cursorStr, err = encodeCursor(cursor)
		if err != nil {
			return nil, "", err
		}
	}

	return majors, cursorStr, nil
}

// NewMajorRepository .
func NewMajorRepository(db *sql.DB) sa.MajorRepository {
	return majorRepo{db: db}
}
