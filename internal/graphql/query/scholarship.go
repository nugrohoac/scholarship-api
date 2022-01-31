package query

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// ScholarshipQuery ...
type ScholarshipQuery struct {
	scholarshipService sa.ScholarshipService
}

// FetchScholarship ...
func (s ScholarshipQuery) FetchScholarship(ctx context.Context, param sa.InputScholarshipFilter) (*resolver.ScholarshipFeedResolver, error) {
	filter := sa.ScholarshipFilter{}

	if param.Limit != nil {
		filter.Limit = uint64(*param.Limit)
	}

	if param.Cursor != nil {
		filter.Cursor = *param.Cursor
	}

	if param.SponsorID != nil {
		filter.SponsorID = int64(*param.SponsorID)
	}

	if param.Status != nil {
		for _, status := range *param.Status {
			filter.Status = append(filter.Status, *status)
		}
	}

	if param.Name != nil {
		filter.Name = *param.Name
	}

	scholarshipFeed, err := s.scholarshipService.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.ScholarshipFeedResolver{ScholarshipFeed: scholarshipFeed}, nil
}

// GetScholarshipByID ...
func (s ScholarshipQuery) GetScholarshipByID(ctx context.Context, param struct{ ID int32 }) (*resolver.ScholarshipResolver, error) {
	scholarship, err := s.scholarshipService.GetByID(ctx, int64(param.ID))
	if err != nil {
		return nil, err
	}

	return &resolver.ScholarshipResolver{Scholarship: scholarship}, nil
}

// NewScholarshipQuery ...
func NewScholarshipQuery(scholarshipService sa.ScholarshipService) ScholarshipQuery {
	return ScholarshipQuery{scholarshipService: scholarshipService}
}
