package mutation

import (
	"context"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// ReportMutation .
type ReportMutation struct {
	reportService business.ReportService
}

// SubmitReport .
func (r ReportMutation) SubmitReport(ctx context.Context, param struct {
	ApplicantID int32
	File        entity.InputImage
}) (*string, error) {
	image := entity.Image{
		URL:    param.File.URL,
		Width:  param.File.Width,
		Height: param.File.Height,
	}

	if param.File.Mime != nil {
		image.Mime = *param.File.Mime
	}

	if param.File.Caption != nil {
		image.Caption = *param.File.Caption
	}

	message, err := r.reportService.Store(ctx, entity.ApplicantReport{ApplicantID: int64(param.ApplicantID), File: image})
	if err != nil {
		return nil, err
	}

	return &message, err
}

// NewReportMutation .
func NewReportMutation(reportService business.ReportService) ReportMutation {
	return ReportMutation{reportService: reportService}
}
