package postgresql

import (
	"context"
	"database/sql"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/sirupsen/logrus"

	sq "github.com/Masterminds/squirrel"
)

type requirementDesc struct {
	db *sql.DB
}

// Fetch ...
func (r requirementDesc) Fetch(ctx context.Context, scholarshipIDs []int64) (map[int64][]string, error) {
	response := map[int64][]string{}

	qSelect := sq.Select("scholarship_id", "description").
		From("requirement_description")

	if len(scholarshipIDs) > 0 {
		qSelect = qSelect.Where(sq.Eq{"scholarship_id": scholarshipIDs})
	}

	query, args, err := qSelect.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	for rows.Next() {
		var (
			scholarshipID int64
			description   string
		)

		if err = rows.Scan(&scholarshipID, &description); err != nil {
			return nil, err
		}

		response[scholarshipID] = append(response[scholarshipID], description)
	}

	return response, nil
}

// NewRequirementDescriptionRepository ...
func NewRequirementDescriptionRepository(db *sql.DB) business.RequirementDescriptionRepository {
	return requirementDesc{db: db}
}
