package degree

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type degreeService struct {
	degreeRepository business.DegreeRepository
}

// Fetch ...
func (d degreeService) Fetch(ctx context.Context) ([]entity.Degree, error) {
	return d.degreeRepository.Fetch(ctx)
}

// NewDegreeService ...
func NewDegreeService(degreeRepository business.DegreeRepository) business.DegreeService {
	return degreeService{degreeRepository: degreeRepository}
}
