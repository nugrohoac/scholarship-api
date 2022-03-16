package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type applicantRepository struct {
	db *sql.DB
}

// Fetch .
func (a applicantRepository) Fetch(ctx context.Context, filter entity.FilterApplicant) ([]entity.Applicant, string, error) {
	qSelect := sq.Select("us.id",
		"us.scholarship_id",
		"us.user_id",
		"us.status",
		"us.created_at",
		"u.id",
		"u.name",
		"u.type",
		"u.email",
		"u.phone_no",
		"u.photo",
		"u.company_name",
		"u.status",
		"u.country_id",
		"u.postal_code",
		"u.address",
		"u.gender",
		"u.ethnic_id",
		"u.birth_date",
		"u.birth_place",
		"u.bank_id",
		"u.bank_account_no",
		"u.bank_account_name",
		"u.ethnic_id",
		"e.name",
	).From("user_scholarship us").
		Join("\"user\" u on us.user_id  = u.id").
		LeftJoin("ethnic e on e.id = u.ethnic_id").
		Where(sq.Eq{"us.scholarship_id": filter.ScholarshipID}).
		PlaceholderFormat(sq.Dollar).
		OrderBy("us.created_at desc")

	if filter.Limit > 0 {
		qSelect = qSelect.Limit(filter.Limit)
	}

	if filter.Cursor != "" {
		cursorTime, err := decodeCursor(filter.Cursor)
		if err != nil {
			return nil, "", err
		}

		qSelect = qSelect.Where(sq.Lt{"us.created_at": cursorTime})
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := a.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		applicants = make([]entity.Applicant, 0)
		cursorStr  string
		cursor     time.Time
	)

	for rows.Next() {
		var (
			applicant  entity.Applicant
			bytePhoto  []byte
			ethnicName sql.NullString
		)

		if err = rows.Scan(
			&applicant.ID,
			&applicant.ScholarshipID,
			&applicant.UserID,
			&applicant.Status,
			&applicant.ApplyDate,
			&applicant.User.ID,
			&applicant.User.Name,
			&applicant.User.Type,
			&applicant.User.Email,
			&applicant.User.PhoneNo,
			&bytePhoto,
			&applicant.User.CompanyName,
			&applicant.User.Status,
			&applicant.User.CountryID,
			&applicant.User.PostalCode,
			&applicant.User.Address,
			&applicant.User.Gender,
			&applicant.User.EthnicID,
			&applicant.User.BirthDate,
			&applicant.User.BirthPlace,
			&applicant.User.BankID,
			&applicant.User.BankAccountNo,
			&applicant.User.BankAccountName,
			&applicant.User.EthnicID,
			&ethnicName,
		); err != nil {
			return nil, "", err
		}

		if ethnicName.Valid {
			applicant.User.Ethnic.Name = ethnicName.String
		}

		if bytePhoto != nil {
			if err = json.Unmarshal(bytePhoto, &applicant.User.Photo); err != nil {
				return nil, "", err
			}
		}

		applicant.UserID = applicant.User.ID
		applicant.User.Ethnic.ID = applicant.User.EthnicID
		cursor = applicant.ApplyDate

		applicants = append(applicants, applicant)
	}

	cursorStr, err = encodeCursor(cursor)
	if err != nil {
		return nil, "", err
	}

	return applicants, cursorStr, nil
}

// NewApplicantRepository .
func NewApplicantRepository(db *sql.DB) business.ApplicantRepository {
	return applicantRepository{db: db}
}
