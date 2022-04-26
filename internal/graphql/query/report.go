package query

import (
	"context"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// ReportQuery .
type ReportQuery struct {
	reportService business.ReportService
}

// FetchReport .
func (r ReportQuery) FetchReport(ctx context.Context, param struct {
	ApplicantID *int32
	SponsorID   *int32
	Limit       *int32
	Cursor      *string
}) (*resolver.ReportFeedResolver, error) {
	filter := entity.ReportFilter{
		Limit: 20,
	}

	if param.ApplicantID != nil {
		filter.ApplicantID = int64(*param.ApplicantID)
	}

	if param.SponsorID != nil {
		filter.SponsorID = int64(*param.SponsorID)
	}

	if param.Limit != nil {
		filter.Limit = uint64(*param.Limit)
	}

	if param.Cursor != nil {
		filter.Cursor = *param.Cursor
	}

	reportFeed, err := r.reportService.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.ReportFeedResolver{ReportFeed: reportFeed}, err
}

// NewReportQuery .
func NewReportQuery(reportService business.ReportService) ReportQuery {
	return ReportQuery{reportService: reportService}
}
