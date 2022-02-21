package resolver

import sa "github.com/Nusantara-Muda/scholarship-api"

// SchoolResolver .
type SchoolResolver struct {
	School sa.School
}

// ID .
func (s SchoolResolver) ID() *int32 {
	ID := int32(s.School.ID)
	return &ID
}

// Name .
func (s SchoolResolver) Name() *string {
	return &s.School.Name
}

// Type .
func (s SchoolResolver) Type() *string {
	return &s.School.Type
}

// Address .
func (s SchoolResolver) Address() *string {
	return &s.School.Address
}

// Status .
func (s SchoolResolver) Status() *int32 {
	status := int32(s.School.Status)
	return &status
}
