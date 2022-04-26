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
	userRepo      business.UserRepository
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

	if user.ID != applicant.UserID {
		return "", errors.ErrNotAllowed{Message: "student not apply scholarship"}
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

	var (
		applicant entity.Applicant
		users     []entity.User
	)

	if user.Type == entity.Sponsor {
		users, _, err = r.userRepo.Fetch(ctx, entity.UserFilter{Email: user.Email})
		if err != nil {
			return entity.ReportFeed{}, err
		}

		if user.ID != users[0].ID {
			return entity.ReportFeed{}, errors.ErrNotAllowed{Message: "sponsor not own of scholarship"}
		}
	}

	if user.Type == entity.Student {
		applicant, err = r.applicantRepo.GetByID(ctx, filter.ApplicantID)
		if err != nil {
			return entity.ReportFeed{}, err
		}

		if user.ID != applicant.UserID {
			return entity.ReportFeed{}, errors.ErrNotAllowed{Message: "student not apply scholarship"}
		}
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
func NewReportService(reportRepository business.ReportRepository, applicantRepo business.ApplicantRepository, userRepo business.UserRepository) business.ReportService {
	return reportService{
		reportRepo:    reportRepository,
		applicantRepo: applicantRepo,
		userRepo:      userRepo,
	}
}
