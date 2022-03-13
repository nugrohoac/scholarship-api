package postgresql

import (
	"context"
	"database/sql"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/sirupsen/logrus"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type schoolRepo struct {
	db *sql.DB
}

// Create .
func (s schoolRepo) Create(ctx context.Context, school entity.School) (entity.School, error) {
	school.Status = 1
	school.CreatedAt = time.Now()

	query, args, err := sq.Insert("school").
		Columns("name",
			"type",
			"address",
			"status",
			"created_at",
			"created_by",
		).Values(school.Name,
		school.Type,
		school.Address,
		school.Status,
		school.CreatedAt,
		school.CreatedBy,
	).PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return entity.School{}, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&school.ID); err != nil {
		return entity.School{}, err
	}

	return school, nil
}

// Fetch .
func (s schoolRepo) Fetch(ctx context.Context, filter entity.SchoolFilter) ([]entity.School, string, error) {
	qSelect := sq.Select("id",
		"type",
		"name",
		"address",
		"status",
		"created_at",
	).From("school").
		OrderBy("created_at desc")

	if filter.Limit > 0 {
		qSelect = qSelect.Limit(filter.Limit)
	}

	if filter.Type != "" {
		qSelect = qSelect.Where(sq.Eq{"type": filter.Type})
	}

	if filter.Cursor != "" {
		cursor, err := decodeCursor(filter.Cursor)
		if err != nil {
			return nil, "", err
		}

		qSelect = qSelect.Where(sq.Lt{"created_at": cursor})
	}

	if filter.Name != "" {
		name := "%" + strings.ToLower(filter.Name) + "%"
		qSelect = qSelect.Where(sq.Like{"LOWER(name)": name})
	}

	query, args, err := qSelect.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		schools   = make([]entity.School, 0)
		cursor    = time.Time{}
		cursorStr = ""
	)

	for rows.Next() {
		var school entity.School

		if err = rows.Scan(
			&school.ID,
			&school.Type,
			&school.Name,
			&school.Address,
			&school.Status,
			&school.CreatedAt,
		); err != nil {
			return nil, "", err
		}

		cursor = school.CreatedAt
		schools = append(schools, school)
	}

	if !cursor.IsZero() {
		cursorStr, err = encodeCursor(cursor)
		if err != nil {
			return nil, "", err
		}
	}

	return schools, cursorStr, nil
}

// NewSchoolRepository .
func NewSchoolRepository(db *sql.DB) business.SchoolRepository {
	return schoolRepo{db: db}
}
