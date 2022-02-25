package query

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// SchoolQuery .
type SchoolQuery struct {
	schoolService sa.SchoolService
}

// FetchSchool .
func (s SchoolQuery) FetchSchool(ctx context.Context, param sa.InputSchoolFilter) (*resolver.SchoolFeedResolver, error) {
	var filter sa.SchoolFilter

	if param.Limit != nil {
		filter.Limit = uint64(*param.Limit)
	}

	if param.Name != nil {
		filter.Name = *param.Name
	}

	if param.Cursor != nil {
		filter.Cursor = *param.Cursor
	}

	if param.Type != nil {
		filter.Type = *param.Type
	}

	schoolFeed, err := s.schoolService.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.SchoolFeedResolver{SchoolFeed: schoolFeed}, nil
}

// NewSchoolQuery .
func NewSchoolQuery(schoolService sa.SchoolService) SchoolQuery {
	return SchoolQuery{schoolService: schoolService}
}
