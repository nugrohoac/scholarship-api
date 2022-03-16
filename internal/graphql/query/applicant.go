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

	applicantFeed, err := a.applicantService.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.ApplicantFeedResolver{ApplicantFeed: applicantFeed}, nil
}

// NewApplicantQuery .
func NewApplicantQuery(applicantService business.ApplicantService) ApplicantQuery {
	return ApplicantQuery{applicantService: applicantService}
}
