package resolver

import "github.com/Nusantara-Muda/scholarship-api/src/business/entity"

// UserFeedResolver ...
type StudentFeedResolver struct {
	StudentFeed entity.StudentFeed
}

// Cursor .
func (m StudentFeedResolver) Cursor() *string {
	return &m.StudentFeed.Cursor
}

// User .
func (m StudentFeedResolver) Students() *[]*UserResolver {
	students := make([]*UserResolver, 0)

	for _, student := range m.StudentFeed.Students {
		students = append(students, &UserResolver{User: student})
	}

	return &students
}
