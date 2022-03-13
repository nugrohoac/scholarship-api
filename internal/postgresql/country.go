package postgresql

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/sirupsen/logrus"
	"time"
)

type countryRepo struct {
	db *sql.DB
}

// Fetch ...
func (c countryRepo) Fetch(ctx context.Context, filter entity.CountryFilter) ([]entity.Country, string, error) {
	qSelect := sq.Select("id", "name", "created_at").
		From("country").
		OrderBy("created_at desc").
		PlaceholderFormat(sq.Dollar)

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

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		countries = make([]entity.Country, 0)
		cursor    = time.Time{}
	)

	for rows.Next() {
		var country entity.Country

		if err = rows.Scan(
			&country.ID,
			&country.Name,
			&country.CreatedAt,
		); err != nil {
			return nil, "", err
		}

		cursor = country.CreatedAt
		countries = append(countries, country)
	}

	cursorStr, err := encodeCursor(cursor)
	if err != nil {
		return nil, "", err
	}

	return countries, cursorStr, nil
}

// Fetch .

// NewCountryRepository ...
func NewCountryRepository(db *sql.DB) business.CountryRepository {
	return countryRepo{db: db}
}
