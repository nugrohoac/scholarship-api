package postgresql

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type assessmentRepository struct {
	db *sql.DB
}

// Submit .
func (a assessmentRepository) Submit(ctx context.Context, ApplicantID int64, eligibilities []entity.ApplicantEligibility, scores []entity.ApplicantScore) error {
	timeNow := time.Now()

	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	// there are chance eligibilities is nil
	if len(eligibilities) > 0 {
		qInsert := sq.Insert("applicant_eligibility").
			Columns("applicant_id",
				"requirement_id",
				"value",
				"created_at",
			).PlaceholderFormat(sq.Dollar)

		for _, eligibility := range eligibilities {
			qInsert = qInsert.Values(
				ApplicantID,
				eligibility.RequirementID,
				eligibility.Value,
				timeNow,
			)
		}

		query, args, err := qInsert.ToSql()
		if err != nil {
			return err
		}

		if _, err = tx.ExecContext(ctx, query, args...); err != nil {
			return err
		}
	}

	// insert into applicant score
	qInsert := sq.Insert("applicant_score").
		Columns("applicant_id",
			"name",
			"value",
		).PlaceholderFormat(sq.Dollar)

	for _, score := range scores {
		qInsert = qInsert.Values(ApplicantID,
			score.Name,
			score.Value,
		)
	}

	query, args, err := qInsert.ToSql()
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			logrus.Error(err)
		}

		return err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			logrus.Error(err)
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			logrus.Error(err)
		}

		return err
	}

	return nil
}

// NewAssessmentRepository ,
func NewAssessmentRepository(db *sql.DB) business.AssessmentRepository {
	return assessmentRepository{db: db}
}
