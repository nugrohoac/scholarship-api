package resolver

import sa "github.com/Nusantara-Muda/scholarship-api"

// MajorResolver .
type MajorResolver struct {
	Major sa.Major
}

// ID .
func (m MajorResolver) ID() *int32 {
	ID := int32(m.Major.ID)
	return &ID
}

// Name .
func (m MajorResolver) Name() *string {
	return &m.Major.Name
}

// MajorFeedResolver ...
type MajorFeedResolver struct {
	MajorFeed sa.MajorFeed
}

// Cursor .
func (m MajorFeedResolver) Cursor() *string {
	return &m.MajorFeed.Cursor
}

// Majors .
func (m MajorFeedResolver) Majors() *[]*MajorResolver {
	majors := make([]*MajorResolver, 0)

	for _, major := range m.MajorFeed.Majors {
		major := major

		majors = append(majors, &MajorResolver{Major: major})
	}

	return &majors
}
