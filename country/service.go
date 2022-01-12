package country

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

type countryService struct {
	countryRepo sa.CountryRepository
}

// Fetch ...
func (c countryService) Fetch(ctx context.Context, filter sa.CountryFilter) (sa.CountryFeed, error) {
	countries, cursor, err := c.countryRepo.Fetch(ctx, filter)
	if err != nil {
		return sa.CountryFeed{}, err
	}

	countryFeed := sa.CountryFeed{
		Cursor:    cursor,
		Countries: countries,
	}

	return countryFeed, nil
}

// NewCountryService ...
func NewCountryService(countryRepo sa.CountryRepository) sa.CountryService {
	return countryService{countryRepo: countryRepo}
}
