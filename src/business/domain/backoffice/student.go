package backoffice

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/sirupsen/logrus"
	"time"
)

type studentRepo struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) business.StudentRepository {
	return studentRepo{db: db}
}

func (u studentRepo) FetchStudent(ctx context.Context, filter entity.StudentFilter) ([]entity.User, string, error) {
	qSelect := sq.Select("id", "lower(name)", "lower(email)", "phone_no", "lower(address)", "lower(birth_place)", "career_goal", "created_at").
		From("public.user").
		Where(sq.Eq{"type": "student"}).
		OrderBy("created_at desc").
		PlaceholderFormat(sq.Dollar)

	if filter.Limit > 0 {
		qSelect = qSelect.Limit(uint64(filter.Limit))
	}

	if filter.Cursor != "" {
		cursorTime, err := decodeCursor(filter.Cursor)
		if err != nil {
			return nil, "", err
		}
		qSelect = qSelect.Where(sq.Lt{"created_at": cursorTime})
	}

	if filter.SearchText != "" {
		searchText := "%" + filter.SearchText + "%"
		qSelect = qSelect.Where(sq.Or{
			sq.Like{"name": searchText},
			sq.Like{"email": searchText},
			sq.Like{"address": searchText},
			sq.Like{"birth_place": searchText},
		})
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := u.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		users     = make([]entity.User, 0)
		cursor    = time.Time{}
		cursorStr string
	)

	for rows.Next() {
		var user entity.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PhoneNo,
			&user.Address,
			&user.BirthPlace,
			&user.CareerGoal,
			&user.CreatedAt,
		); err != nil {
			return nil, "", err
		}

		cursor = user.CreatedAt
		users = append(users, user)
	}

	cursorStr, err = encodeCursor(cursor)
	if err != nil {
		return nil, "", err
	}

	return users, cursorStr, nil
}
