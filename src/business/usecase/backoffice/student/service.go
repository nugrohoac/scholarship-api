package student

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type studentService struct {
	studentRepo business.StudentRepository
}

// NewSponsorService ...
func NewStudentService(studentRepo business.StudentRepository) business.StudentService {
	return studentService{studentRepo: studentRepo}
}

// Fetch ...
func (s studentService) FetchStudent(ctx context.Context, filter entity.StudentFilter) (entity.StudentFeed, error) {
	students, cursor, err := s.studentRepo.FetchStudent(ctx, filter)
	if err != nil {
		return entity.StudentFeed{}, err
	}

	studentsFeed := entity.StudentFeed{
		Cursor:   cursor,
		Students: students,
	}

	filter.Cursor = cursor
	filter.Limit = 1

	if len(students) == 0 {
		studentsFeed.Cursor = ""
	}

	return entity.StudentFeed{Cursor: cursor, Students: students}, nil
}
