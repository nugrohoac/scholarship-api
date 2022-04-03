package backoffice

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// SponsorQuery ...
type SponsorQuery struct {
	sponsorService business.SponsorService
}

// NewSponsorQuery ...
func NewSponsorQuery(sponsorService business.SponsorService) SponsorQuery {
	return SponsorQuery{sponsorService: sponsorService}
}

// FetchSponsor .
func (c SponsorQuery) FetchSponsor(ctx context.Context, param entity.InputSponsorFilter) (*resolver.SponsorFeedResolver, error) {
	filter := entity.SponsorFilter{}

	if param.Limit != nil {
		filter.Limit = int(*param.Limit)
	}

	if param.Cursor != nil {
		filter.Cursor = *param.Cursor
	}

	if param.SearchText != nil {
		filter.SearchText = *param.SearchText
	}

	sponsorFeed, err := c.sponsorService.FetchSponsor(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.SponsorFeedResolver{SponsorFeed: sponsorFeed}, nil
}
