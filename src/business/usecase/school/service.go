package school

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type schoolService struct {
	schoolRepo business.SchoolRepository
}

// Create ...
func (s schoolService) Create(ctx context.Context, school entity.School) (entity.School, error) {
	user, err := common.GetUserOnContext(ctx)
	if err != nil {
		return entity.School{}, err
	}

	school.CreatedBy = user.Email

	return s.schoolRepo.Create(ctx, school)
}

// Fetch ...
func (s schoolService) Fetch(ctx context.Context, filter entity.SchoolFilter) (entity.SchoolFeed, error) {
	schools, cursor, err := s.schoolRepo.Fetch(ctx, filter)
	if err != nil {
		return entity.SchoolFeed{}, err
	}

	return entity.SchoolFeed{Cursor: cursor, Schools: schools}, nil
}

// NewSchoolService ...
func NewSchoolService(schoolRepo business.SchoolRepository) business.SchoolService {
	return schoolService{schoolRepo: schoolRepo}
}
