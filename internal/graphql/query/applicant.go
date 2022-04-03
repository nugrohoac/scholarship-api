package query

import (
	"context"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// ApplicantQuery .
type ApplicantQuery struct {
	applicantService business.ApplicantService
}

// FetchApplicant .
func (a ApplicantQuery) FetchApplicant(ctx context.Context, param entity.InputApplicantFilter) (*resolver.ApplicantFeedResolver, error) {
	filter := entity.FilterApplicant{
		SponsorID:     int64(param.SponsorID),
		ScholarshipID: int64(param.ScholarshipID),
	}

	if param.Limit != nil {
		filter.Limit = uint64(*param.Limit)
	}

	if param.Cursor != nil {
		filter.Cursor = *param.Cursor
	}

	if param.Status != nil {
		if len(*param.Status) > 0 {
			for _, status := range *param.Status {
				filter.Status = append(filter.Status, *status)
			}
		}
	}

	applicantFeed, err := a.applicantService.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.ApplicantFeedResolver{ApplicantFeed: applicantFeed}, nil
}

// GetApplicantByID .
func (a ApplicantQuery) GetApplicantByID(ctx context.Context, param struct{ ID int32 }) (*resolver.ApplicantResolver, error) {
	applicant, err := a.applicantService.GetByID(ctx, int64(param.ID))
	if err != nil {
		return nil, err
	}

	return &resolver.ApplicantResolver{Applicant: applicant}, nil
}

// NewApplicantQuery .
func NewApplicantQuery(applicantService business.ApplicantService) ApplicantQuery {
	return ApplicantQuery{applicantService: applicantService}
}
