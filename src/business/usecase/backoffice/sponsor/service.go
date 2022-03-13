package sponsor

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type sponsorService struct {
	sponsorRepo business.SponsorRepository
}

// NewSponsorService ...
func NewSponsorService(sponsorRepo business.SponsorRepository) business.SponsorService {
	return sponsorService{sponsorRepo: sponsorRepo}
}


// Fetch ...
func (s sponsorService) FetchSponsor(ctx context.Context, filter entity.SponsorFilter) (entity.SponsorFeed, error) {
	sponsors, cursor, err := s.sponsorRepo.FetchSponsor(ctx, filter)
	if err != nil {
		return entity.SponsorFeed{}, err
	}

	return entity.SponsorFeed{Cursor: cursor, Sponsors: sponsors}, nil
}