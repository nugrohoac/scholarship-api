package degree

import sa "github.com/Nusantara-Muda/scholarship-api"

type degreeRepo struct {
	degrees []string
}

// Get ...
func (d degreeRepo) Get() []string {
	return d.degrees
}

// NewDegreeRepository .
func NewDegreeRepository(degrees []string) sa.DegreeRepository {
	return degreeRepo{degrees: degrees}
}
