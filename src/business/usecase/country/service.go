package country

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type countryService struct {
	countryRepo business.CountryRepository
}

// Fetch ...
func (c countryService) Fetch(ctx context.Context, filter entity.CountryFilter) (entity.CountryFeed, error) {
	countries, cursor, err := c.countryRepo.Fetch(ctx, filter)
	if err != nil {
		return entity.CountryFeed{}, err
	}

	countryFeed := entity.CountryFeed{
		Cursor:    cursor,
		Countries: countries,
	}

	return countryFeed, nil
}

// NewCountryService ...
func NewCountryService(countryRepo business.CountryRepository) business.CountryService {
	return countryService{countryRepo: countryRepo}
}
