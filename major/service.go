package major

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type majorService struct {
	majorRepo sa.MajorRepository
}

// Fetch .
func (m majorService) Fetch(ctx context.Context, filter sa.MajorFilter) (sa.MajorFeed, error) {
	majors, cursor, err := m.majorRepo.Fetch(ctx, filter)
	if err != nil {
		return sa.MajorFeed{}, err
	}

	response := sa.MajorFeed{
		Cursor: cursor,
		Majors: majors,
	}

	return response, nil
}

// NewMajorService .
func NewMajorService(majorRepo sa.MajorRepository) sa.MajorService {
	return majorService{majorRepo: majorRepo}
}
