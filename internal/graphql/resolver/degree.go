package resolver

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// DegreeResolver .
type DegreeResolver struct {
	Degree entity.Degree
}

// ID .
func (d DegreeResolver) ID() *int32 {
	return &d.Degree.ID
}

// Name .
func (d DegreeResolver) Name() *string {
	return &d.Degree.Name
}
