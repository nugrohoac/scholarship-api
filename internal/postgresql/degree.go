package postgresql

import (
	"context"
	"database/sql"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/sirupsen/logrus"

	sq "github.com/Masterminds/squirrel"
)

type degreeRepo struct {
	db *sql.DB
}

// Fetch ...
func (d degreeRepo) Fetch(ctx context.Context) ([]entity.Degree, error) {
	query, args, err := sq.Select("id", "name").
		From("degree").
		OrderBy("created_at desc").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var degrees []entity.Degree

	for rows.Next() {
		var degree entity.Degree

		if err = rows.Scan(&degree.ID, &degree.Name); err != nil {
			return nil, err
		}

		degrees = append(degrees, degree)
	}

	return degrees, nil
}

// NewDegreeRepository .
func NewDegreeRepository(db *sql.DB) business.DegreeRepository {
	return degreeRepo{db: db}
}
