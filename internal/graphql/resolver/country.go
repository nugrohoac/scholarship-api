package resolver

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// CountryResolver ...
type CountryResolver struct {
	Country entity.Country
}

// ID ..
func (c CountryResolver) ID() *int32 {
	ID := c.Country.ID
	return &ID
}

// Name ...
func (c CountryResolver) Name() *string {
	return &c.Country.Name
}

// CountryFeedResolver ...
type CountryFeedResolver struct {
	CountryFeed entity.CountryFeed
}

// Cursor ...
func (c CountryFeedResolver) Cursor() *string {
	return &c.CountryFeed.Cursor
}

// Countries ...
func (c CountryFeedResolver) Countries() *[]*CountryResolver {
	countryResolvers := make([]*CountryResolver, 0)

	for _, country := range c.CountryFeed.Countries {
		country := country

		countryResolvers = append(countryResolvers, &CountryResolver{
			Country: country,
		})
	}

	return &countryResolvers
}
