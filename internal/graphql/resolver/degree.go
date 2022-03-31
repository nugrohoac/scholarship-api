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

// Rank .
func (d DegreeResolver) Rank() *int32 {
	rank := int32(d.Degree.Rank)
	return &rank
}
