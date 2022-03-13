package query

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// CountryQuery ...
type CountryQuery struct {
	countryService business.CountryService
}

// FetchCountry .
func (c CountryQuery) FetchCountry(ctx context.Context, param entity.InputCountryFilter) (*resolver.CountryFeedResolver, error) {
	filter := entity.CountryFilter{}

	if param.Limit != nil {
		filter.Limit = int(*param.Limit)
	}

	if param.Cursor != nil {
		filter.Cursor = *param.Cursor
	}

	if param.Name != nil {
		filter.Name = *param.Name
	}

	countryFeed, err := c.countryService.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.CountryFeedResolver{CountryFeed: countryFeed}, nil
}

// NewCountryQuery ...
func NewCountryQuery(countryService business.CountryService) CountryQuery {
	return CountryQuery{countryService: countryService}
}
