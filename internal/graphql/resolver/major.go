package resolver

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// MajorResolver .
type MajorResolver struct {
	Major entity.Major
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
	MajorFeed entity.MajorFeed
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
