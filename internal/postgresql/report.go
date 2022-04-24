package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type reportRepository struct {
	db *sql.DB
}

// Fetch .
func (r reportRepository) Fetch(ctx context.Context, filter entity.ReportFilter) ([]entity.ApplicantReport, string, error) {
	qSelect := sq.Select("id",
		"applicant_id",
		"file",
		"created_at",
	).From("applicant_report").
		OrderBy("created_at desc").
		PlaceholderFormat(sq.Dollar)

	if filter.ApplicantID > 0 {
		qSelect = qSelect.Where(sq.Eq{"applicant_id": filter.ApplicantID})
	}

	if filter.Limit > 0 {
		qSelect = qSelect.Limit(filter.Limit)
	}

	if filter.Cursor != "" {
		timeCursor, err := decodeCursor(filter.Cursor)
		if err != nil {
			return nil, "", err
		}

		qSelect = qSelect.Where(sq.Lt{"created_at": timeCursor})
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	reports := make([]entity.ApplicantReport, 0)
	var (
		cursorTime time.Time
		cursorStr  string
		byteFile   []byte
	)

	for rows.Next() {
		var report entity.ApplicantReport

		if err = rows.Scan(
			&report.ID,
			&report.ApplicantID,
			&byteFile,
			&report.CreatedAt,
		); err != nil {
			return nil, "", err
		}

		if err = json.Unmarshal(byteFile, &report.File); err != nil {
			return nil, "", err
		}

		cursorTime = report.CreatedAt
		reports = append(reports, report)
	}

	cursorStr, err = encodeCursor(cursorTime)
	if err != nil {
		return nil, "", err
	}

	return reports, cursorStr, nil
}

// Store .
func (r reportRepository) Store(ctx context.Context, report entity.ApplicantReport) error {
	byteFile, err := json.Marshal(report.File)
	if err != nil {
		return err
	}

	query, args, err := sq.Insert("applicant_report").
		Columns("applicant_id", "file", "created_at").
		Values(report.ApplicantID, byteFile, time.Now()).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// NewReportRepository .
func NewReportRepository(db *sql.DB) business.ReportRepository {
	return reportRepository{db: db}
}
