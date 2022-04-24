package report

import (
	"context"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"
)

type reportService struct {
	reportRepo    business.ReportRepository
	applicantRepo business.ApplicantRepository
}

// Store .
func (r reportService) Store(ctx context.Context, report entity.ApplicantReport) (string, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return "", err
	}

	applicant, err := r.applicantRepo.GetByID(ctx, report.ApplicantID)
	if err != nil {
		return "", err
	}

	if user.ID != applicant.Scholarship.SponsorID {
		return "", errors.ErrNotAllowed{Message: "user is not own of scholarship"}
	}

	if err = r.reportRepo.Store(ctx, report); err != nil {
		return "", err
	}

	return "success", nil
}

// Fetch .
func (r reportService) Fetch(ctx context.Context, filter entity.ReportFilter) (entity.ReportFeed, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return entity.ReportFeed{}, err
	}

	applicant, err := r.applicantRepo.GetByID(ctx, filter.ApplicantID)
	if err != nil {
		return entity.ReportFeed{}, err
	}

	if user.ID != applicant.Scholarship.SponsorID {
		return entity.ReportFeed{}, errors.ErrNotAllowed{Message: "user is not own of scholarship"}
	}

	reports, cursor, err := r.reportRepo.Fetch(ctx, filter)
	if err != nil {
		return entity.ReportFeed{}, err
	}

	reportFeed := entity.ReportFeed{
		Cursor:  cursor,
		Reports: reports,
	}

	filter.Cursor = cursor
	filter.Limit = 1
	reports, _, err = r.reportRepo.Fetch(ctx, filter)
	if err != nil {
		return entity.ReportFeed{}, err
	}

	if len(reports) == 0 {
		reportFeed.Cursor = ""
	}

	return reportFeed, nil
}

// NewReportService .
func NewReportService(reportRepository business.ReportRepository, applicantRepo business.ApplicantRepository) business.ReportService {
	return reportService{
		reportRepo:    reportRepository,
		applicantRepo: applicantRepo,
	}
}
