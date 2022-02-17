package query

import sa "github.com/Nusantara-Muda/scholarship-api"

// DegreeQuery ...
type DegreeQuery struct {
	degreeService sa.DegreeService
}

// GetDegree ...
func (d DegreeQuery) GetDegree() (*[]*string, error) {
	degrees := d.degreeService.Get()
	return &degrees, nil
}

// NewDegreeQuery ...
func NewDegreeQuery(degreeService sa.DegreeService) DegreeQuery {
	return DegreeQuery{degreeService: degreeService}
}
