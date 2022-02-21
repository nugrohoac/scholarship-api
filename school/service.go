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

// NewSchoolService ...
func NewSchoolService(schoolRepo sa.SchoolRepository) sa.SchoolService {
	return schoolService{schoolRepo: schoolRepo}
}
