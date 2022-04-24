package resolver

import (
	"time"

	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// ReportResolver .
type ReportResolver struct {
	Report entity.ApplicantReport
}

// ID .
func (r ReportResolver) ID() *int32 {
	ID := int32(r.Report.ID)

	return &ID
}

// ApplicantID .
func (r ReportResolver) ApplicantID() *int32 {
	applicantID := int32(r.Report.ApplicantID)

	return &applicantID
}

// File .
func (r ReportResolver) File() *ImageResolver {
	return &ImageResolver{Image: r.Report.File}
}

// CreatedAt .
func (r ReportResolver) CreatedAt() *string {
	createdAt := r.Report.CreatedAt.Format(time.RFC3339)
	return &createdAt
}

// ReportFeedResolver .
type ReportFeedResolver struct {
	reportFeed entity.ReportFeed
}

// Cursor .
func (r ReportFeedResolver) Cursor() *string {
	return &r.reportFeed.Cursor
}

// Reports .
func (r ReportFeedResolver) Reports() *[]*ReportResolver {
	reportResolvers := make([]*ReportResolver, 0)

	for _, report := range r.reportFeed.Reports {
		report := report

		reportResolvers = append(reportResolvers, &ReportResolver{Report: report})
	}

	return &reportResolvers
}
