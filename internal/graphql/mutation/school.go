package mutation

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// SchoolMutation ...
type SchoolMutation struct {
	schoolService business.SchoolService
}

// CreateSchool ...
func (s SchoolMutation) CreateSchool(ctx context.Context, param entity.InputSchool) (*resolver.SchoolResolver, error) {
	school := entity.School{
		Name:    param.Name,
		Type:    param.Type,
		Address: param.Address,
	}

	school, err := s.schoolService.Create(ctx, school)
	if err != nil {
		return nil, err
	}

	return &resolver.SchoolResolver{School: school}, nil
}

// NewSchoolMutation ...
func NewSchoolMutation(schoolService business.SchoolService) SchoolMutation {
	return SchoolMutation{schoolService: schoolService}
}
