package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type reportRepository struct {
	db *sql.DB
}

// Store .
func (r reportRepository) Store(ctx context.Context, report entity.ApplicantReport) error {
	byteReport, err := json.Marshal(report.Report)
	if err != nil {
		return err
	}

	query, args, err := sq.Insert("applicant_report").
		Columns("applicant_id", "report", "created_at").
		Values(report.ApplicantID, byteReport, time.Now()).
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
