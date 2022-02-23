package school

import (
	"context"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

type schoolService struct {
	schoolRepo sa.SchoolRepository
}

// Create ...
func (s schoolService) Create(ctx context.Context, school sa.School) (sa.School, error) {
	user, err := sa.GetUserOnContext(ctx)
	if err != nil {
		return sa.School{}, err
	}

	school.CreatedBy = user.Email

	return s.schoolRepo.Create(ctx, school)
}

// Fetch ...
func (s schoolService) Fetch(ctx context.Context, filter sa.SchoolFilter) (sa.SchoolFeed, error) {
	schools, cursor, err := s.schoolRepo.Fetch(ctx, filter)
	if err != nil {
		return sa.SchoolFeed{}, err
	}

	return sa.SchoolFeed{Cursor: cursor, Schools: schools}, nil
}

// NewSchoolService ...
func NewSchoolService(schoolRepo sa.SchoolRepository) sa.SchoolService {
	return schoolService{schoolRepo: schoolRepo}
}
