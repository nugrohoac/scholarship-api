package query

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// SchoolQuery .
type SchoolQuery struct {
	schoolService business.SchoolService
}

// FetchSchool .
func (s SchoolQuery) FetchSchool(ctx context.Context, param entity.InputSchoolFilter) (*resolver.SchoolFeedResolver, error) {
	var filter entity.SchoolFilter

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
func NewSchoolQuery(schoolService business.SchoolService) SchoolQuery {
	return SchoolQuery{schoolService: schoolService}
}
