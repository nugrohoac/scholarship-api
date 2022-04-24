package mutation

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
)

// ApplicantMutation .
type ApplicantMutation struct {
	applicantService business.ApplicantService
}

// UpdateApplicantStatus .
func (a ApplicantMutation) UpdateApplicantStatus(ctx context.Context, param struct {
	ID     int32
	Status int32
}) (*string, error) {
	message, err := a.applicantService.UpdateStatus(ctx, int64(param.ID), param.Status)
	if err != nil {
		return nil, err
	}

	return &message, err
}

// SubmitRating .
func (a ApplicantMutation) SubmitRating(ctx context.Context, param struct {
	ApplicantID int32
	Rating      int32
}) (*string, error) {
	message, err := a.applicantService.StoreRating(ctx, int64(param.ApplicantID), param.Rating)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// NewApplicantMutation .
func NewApplicantMutation(applicantService business.ApplicantService) ApplicantMutation {
	return ApplicantMutation{applicantService: applicantService}
}
