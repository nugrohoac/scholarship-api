package degree

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

type degreeService struct {
	degreeRepository sa.DegreeRepository
}

// Fetch ...
func (d degreeService) Fetch(ctx context.Context) ([]sa.Degree, error) {
	return d.degreeRepository.Fetch(ctx)
}

// NewDegreeService ...
func NewDegreeService(degreeRepository sa.DegreeRepository) sa.DegreeService {
	return degreeService{degreeRepository: degreeRepository}
}
