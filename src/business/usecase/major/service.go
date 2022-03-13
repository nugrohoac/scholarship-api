package major

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type majorService struct {
	majorRepo business.MajorRepository
}

// Fetch .
func (m majorService) Fetch(ctx context.Context, filter entity.MajorFilter) (entity.MajorFeed, error) {
	majors, cursor, err := m.majorRepo.Fetch(ctx, filter)
	if err != nil {
		return entity.MajorFeed{}, err
	}

	response := entity.MajorFeed{
		Cursor: cursor,
		Majors: majors,
	}

	return response, nil
}

// NewMajorService .
func NewMajorService(majorRepo business.MajorRepository) business.MajorService {
	return majorService{majorRepo: majorRepo}
}
