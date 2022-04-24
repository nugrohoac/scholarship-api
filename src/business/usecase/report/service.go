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

// NewReportService .
func NewReportService(reportRepository business.ReportRepository, applicantRepo business.ApplicantRepository) business.ReportService {
	return reportService{
		reportRepo:    reportRepository,
		applicantRepo: applicantRepo,
	}
}
