package resolver

import sa "github.com/Nusantara-Muda/scholarship-api"

// DegreeResolver .
type DegreeResolver struct {
	Degree sa.Degree
}

// ID .
func (d DegreeResolver) ID() *int32 {
	return &d.Degree.ID
}

// Name .
func (d DegreeResolver) Name() *string {
	return &d.Degree.Name
}
