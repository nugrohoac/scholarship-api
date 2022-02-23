package mutation

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// SchoolMutation ...
type SchoolMutation struct {
	schoolService sa.SchoolService
}

// CreateSchool ...
func (s SchoolMutation) CreateSchool(ctx context.Context, param sa.InputSchool) (*resolver.SchoolResolver, error) {
	school := sa.School{
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
func NewSchoolMutation(schoolService sa.SchoolService) SchoolMutation {
	return SchoolMutation{schoolService: schoolService}
}
