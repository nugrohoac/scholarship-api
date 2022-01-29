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

// GetScholarshipBySponsor ...
func (s ScholarshipQuery) GetScholarshipBySponsor(ctx context.Context, param struct{ ID int32 }) (*resolver.ScholarshipFeedResolver, error) {
	scholarshipFeed, err := s.scholarshipService.GetBySponsor(ctx, int64(param.ID))
	if err != nil {
		return nil, err
	}

	return &resolver.ScholarshipFeedResolver{ScholarshipFeed: scholarshipFeed}, nil
}

// NewScholarshipQuery ...
func NewScholarshipQuery(scholarshipService sa.ScholarshipService) ScholarshipQuery {
	return ScholarshipQuery{scholarshipService: scholarshipService}
}
