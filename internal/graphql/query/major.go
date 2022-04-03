package query

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// MajorQuery .
type MajorQuery struct {
	majorService business.MajorService
}

// FetchMajor .
func (m MajorQuery) FetchMajor(ctx context.Context, param entity.InputMajorFilter) (*resolver.MajorFeedResolver, error) {
	filter := entity.MajorFilter{}
	if param.Limit != nil {
		filter.Limit = uint64(*param.Limit)
	}

	if param.Cursor != nil {
		filter.Cursor = *param.Cursor
	}

	if param.Name != nil {
		filter.Name = *param.Name
	}

	response, err := m.majorService.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.MajorFeedResolver{MajorFeed: response}, nil
}

// NewMajorQuery .
func NewMajorQuery(majorService business.MajorService) MajorQuery {
	return MajorQuery{majorService: majorService}
}
