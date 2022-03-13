package resolver

import "github.com/Nusantara-Muda/scholarship-api/src/business/entity"

// UserFeedResolver ...
type SponsorFeedResolver struct {
	SponsorFeed entity.SponsorFeed
}

// Cursor .
func (m SponsorFeedResolver) Cursor() *string {
	return &m.SponsorFeed.Cursor
}

// User .
func (m SponsorFeedResolver) Sponsors() *[]*UserResolver {
	sponsors := make([]*UserResolver, 0)

	for _, sponsor := range m.SponsorFeed.Sponsors {
		sponsors = append(sponsors, &UserResolver{User: sponsor})
	}

	return &sponsors
}
