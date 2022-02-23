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

// SchoolFeedResolver .
type SchoolFeedResolver struct {
	SchoolFeed sa.SchoolFeed
}

// Cursor .
func (s SchoolFeedResolver) Cursor() *string {
	return &s.SchoolFeed.Cursor
}

// Schools .
func (s SchoolFeedResolver) Schools() *[]*SchoolResolver {
	schools := make([]*SchoolResolver, 0)

	for _, school := range s.SchoolFeed.Schools {
		school := school
		schools = append(schools, &SchoolResolver{School: school})
	}

	return &schools
}
