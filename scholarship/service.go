package scholarship

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

type scholarshipService struct {
	scholarshipRepo sa.ScholarshipRepository
}

// Create ...
// Status of sponsor should 2
func (s scholarshipService) Create(ctx context.Context, scholarship sa.Scholarship) (sa.Scholarship, error) {
	sponsor, err := sa.GetUserOnContext(ctx)
	if err != nil {
		return sa.Scholarship{}, err
	}

	if sponsor.ID != scholarship.SponsorID {
		return sa.Scholarship{}, sa.ErrUnAuthorize{Message: "sponsor id is not match"}
	}

	if sponsor.Status != 2 {
		return sa.Scholarship{}, sa.ErrNotAllowed{Message: "sponsor un complete profile"}
	}

	if scholarship.FundingEnd.Before(scholarship.FundingStart) {
		return sa.Scholarship{}, sa.ErrBadRequest{Message: "scholarship funding end before funding start"}
	}

	scholarship.Sponsor = sponsor

	// default at created, look at readme.md to get more status
	scholarship.Status = 0

	scholarship, err = s.scholarshipRepo.Create(ctx, scholarship)
	if err != nil {
		return sa.Scholarship{}, nil
	}

	// crate invoice here, next sprint ( sprint 3 )

	return scholarship, nil
}

// Fetch ...
func (s scholarshipService) Fetch(ctx context.Context, filter sa.ScholarshipFilter) (sa.ScholarshipFeed, error) {
	scholarships, cursor, err := s.scholarshipRepo.Fetch(ctx, filter)
	if err != nil {
		return sa.ScholarshipFeed{}, err
	}

	scholarshipFeed := sa.ScholarshipFeed{
		Cursor:       cursor,
		Scholarships: scholarships,
	}

	filter.Cursor = cursor
	filter.Limit = 1
	scholarships, _, err = s.scholarshipRepo.Fetch(ctx, filter)
	if err != nil {
		return sa.ScholarshipFeed{}, err
	}

	if len(scholarships) == 0 {
		scholarshipFeed.Cursor = ""
	}

	return scholarshipFeed, nil
}

// GetByID ...
func (s scholarshipService) GetByID(ctx context.Context, ID int64) (sa.Scholarship, error) {
	return s.scholarshipRepo.GetByID(ctx, ID)
}

// NewScholarshipService ...
func NewScholarshipService(scholarshipRepo sa.ScholarshipRepository) sa.ScholarshipService {
	return scholarshipService{scholarshipRepo: scholarshipRepo}
}
