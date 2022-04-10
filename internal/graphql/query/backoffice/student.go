package backoffice

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// SponsorQuery ...
type StudentQuery struct {
	studentService business.StudentService
}

// NewSponsorQuery ...
func NewStudentQuery(studentService business.StudentService) StudentQuery {
	return StudentQuery{studentService: studentService}
}

// FetchSponsor .
func (c StudentQuery) FetchStudent(ctx context.Context, param entity.InputStudentFilter) (*resolver.StudentFeedResolver, error) {
	filter := entity.StudentFilter{}

	if param.Limit != nil {
		filter.Limit = int(*param.Limit)
	}

	if param.Cursor != nil {
		filter.Cursor = *param.Cursor
	}

	if param.SearchText != nil {
		filter.SearchText = *param.SearchText
	}

	studentFeed, err := c.studentService.FetchStudent(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &resolver.StudentFeedResolver{StudentFeed: studentFeed}, nil
}
