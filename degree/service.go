package degree

import sa "github.com/Nusantara-Muda/scholarship-api"

type degreeService struct {
	degreeRepository sa.DegreeRepository
}

// Get ...
func (d degreeService) Get() []*string {
	response := d.degreeRepository.Get()

	degrees := make([]*string, 0)
	for _, r := range response {
		r := r
		degrees = append(degrees, &r)
	}

	return degrees
}

// NewDegreeService ...
func NewDegreeService(degreeRepository sa.DegreeRepository) sa.DegreeService {
	return degreeService{degreeRepository: degreeRepository}
}
