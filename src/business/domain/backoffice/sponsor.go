package backoffice

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/sirupsen/logrus"
	"time"
)

type sponsorRepo struct {
	db *sql.DB
}

// NewCountryRepository ...
func NewSponsorRepository(db *sql.DB) business.SponsorRepository {
	return sponsorRepo{db: db}
}

// Fetch sponsor
func (u sponsorRepo) FetchSponsor(ctx context.Context, filter entity.SponsorFilter) ([]entity.User, string, error) {
	qSelect := sq.Select("id", "name", "email", "phone_no", "photo").
		From("public.user").
		Where(sq.Eq{"type": "sponsor"}).
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

	if filter.Email != "" {
		email := "%" + filter.Email + "%"
		qSelect = qSelect.Where(sq.Like{"LOWER(email)": email})
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
		users = make([]entity.User, 0)
		cursor    = time.Time{}
		bytePhoto []byte
	)

	for rows.Next() {
		var user entity.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PhoneNo,
			&bytePhoto,
		); err != nil {
			return nil, "", err
		}

		cursor = user.CreatedAt
		if bytePhoto != nil {
			if err = json.Unmarshal(bytePhoto, &user.Photo); err != nil {
				return nil, "", err
			}
		}
		users = append(users, user)
	}

	cursorStr, err := encodeCursor(cursor)
	if err != nil {
		return nil, "", err
	}

	return users, cursorStr, nil
}

func encodeCursor(cursor time.Time) (string, error) {
	byt, err := cursor.MarshalText()
	if err != nil {
		return "", err
	}

	stringEncode := base64.StdEncoding.EncodeToString(byt)

	return stringEncode, nil
}

func decodeCursor(cursor string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return time.Time{}, err
	}

	cursorTime, err := time.Parse(time.RFC3339, string(byt))
	if err != nil {
		return time.Time{}, err
	}

	return cursorTime, nil
}
