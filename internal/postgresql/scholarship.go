package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"

	sq "github.com/Masterminds/squirrel"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

type scholarshipRepo struct {
	db *sql.DB
}

// Create ...
func (s scholarshipRepo) Create(ctx context.Context, scholarship sa.Scholarship) (sa.Scholarship, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return sa.Scholarship{}, err
	}

	byteImage, err := json.Marshal(scholarship.Image)
	if err != nil {
		return sa.Scholarship{}, err
	}

	var (
		timeNow     = time.Now()
		errRollback error
	)

	query, args, err := sq.Insert("scholarship").
		Columns("sponsor_id",
			"name",
			"amount",
			"status",
			"image",
			"awardee",
			"deadline",
			"eligibility_description",
			"subsidy_description",
			"funding_start",
			"funding_end",
			"created_at").
		Values(scholarship.SponsorID,
			scholarship.Name,
			scholarship.Amount,
			scholarship.Status,
			byteImage,
			scholarship.Awardee,
			scholarship.Deadline,
			scholarship.EligibilityDescription,
			scholarship.SubsidyDescription,
			scholarship.FundingStart,
			scholarship.FundingEnd,
			timeNow,
		).Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return sa.Scholarship{}, err
	}

	row := tx.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&scholarship.ID); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return sa.Scholarship{}, err
	}

	if len(scholarship.Requirements) > 0 {
		qInsertRequirement := sq.Insert("requirement").
			Columns("scholarship_id",
				"type",
				"name",
				"value",
				"created_at",
			)

		for _, req := range scholarship.Requirements {
			qInsertRequirement = qInsertRequirement.Values(scholarship.ID,
				req.Type,
				req.Name,
				req.Value,
				timeNow,
			)
		}

		query, args, err = qInsertRequirement.PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return sa.Scholarship{}, err
		}

		if _, err = tx.ExecContext(ctx, query, args...); err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return sa.Scholarship{}, err
		}
	}

	if errCommit := tx.Commit(); errCommit != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return sa.Scholarship{}, errCommit
	}

	return scholarship, nil
}

// NewScholarshipRepository ...
func NewScholarshipRepository(db *sql.DB) sa.ScholarshipRepository {
	return scholarshipRepo{db: db}
}
