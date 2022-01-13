package query

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// CountryQuery ...
type CountryQuery struct {
	countryService sa.CountryService
}

// FetchCountry .
func (c CountryQuery) FetchCountry(ctx context.Context, param sa.InputCountryFilter) (*resolver.CountryFeedResolver, error) {
	filter := sa.CountryFilter{}

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
func NewCountryQuery(countryService sa.CountryService) CountryQuery {
	return CountryQuery{countryService: countryService}
}
