package query

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// MajorQuery .
type MajorQuery struct {
	majorService sa.MajorService
}

// GetMajor .
func (m MajorQuery) GetMajor(ctx context.Context, param sa.InputMajorFilter) (*resolver.MajorFeedResolver, error) {
	filter := sa.MajorFilter{}
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
func NewMajorQuery(majorService sa.MajorService) MajorQuery {
	return MajorQuery{majorService: majorService}
}
