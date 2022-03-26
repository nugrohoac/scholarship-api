package resolver

import (
	"time"

	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// UserSchoolResolver .
type UserSchoolResolver struct {
	UserSchool entity.UserSchool
}

// ID .
func (u UserSchoolResolver) ID() *int32 {
	ID := int32(u.UserSchool.ID)
	return &ID
}

// UserID .
func (u UserSchoolResolver) UserID() *int32 {
	userID := int32(u.UserSchool.UserID)
	return &userID
}

// School .
func (u UserSchoolResolver) School() *SchoolResolver {
	return &SchoolResolver{School: u.UserSchool.School}
}

// Degree .
func (u UserSchoolResolver) Degree() *DegreeResolver {
	return &DegreeResolver{Degree: u.UserSchool.Degree}
}

// Major .
func (u UserSchoolResolver) Major() *MajorResolver {
	return &MajorResolver{Major: u.UserSchool.Major}
}

// EnrollmentDate .
func (u UserSchoolResolver) EnrollmentDate() *string {
	ed := u.UserSchool.EnrollmentDate.Format(time.RFC3339)
	return &ed
}

// GraduationDate .
func (u UserSchoolResolver) GraduationDate() *string {
	gd := u.UserSchool.GraduationDate.Format(time.RFC3339)
	return &gd
}

// Gpa .
func (u UserSchoolResolver) Gpa() *float64 {
	return &u.UserSchool.Gpa
}
