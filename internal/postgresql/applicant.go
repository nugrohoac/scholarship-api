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

	if len(filter.Status) > 0 {
		qSelect = qSelect.Where(sq.Eq{"us.status": filter.Status})
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

	if !cursor.IsZero() {
		cursorStr, err = encodeCursor(cursor)
		if err != nil {
			return nil, "", err
		}
	}

	return applicants, cursorStr, nil
}

// GetByID .
// ID is reference to user scholarship id
func (a applicantRepository) GetByID(ctx context.Context, ID int64) (entity.Applicant, error) {
	query, args, err := sq.Select(
		"us.id",
		"us.user_id",
		"u.name",
		"u.photo",
		"u.address",
		"u.birth_date",
		"u.birth_place",
		"u.phone_no",
		"u.gender",
		"u.ethnic_id",
		"e.name",
		"u.country_id",
		"c.name",
		"u.bank_id",
		"b.name",
		"u.bank_account_no",
		"u.bank_account_name",
		"u.career_goal",
		"us.essay",
		"us.recommendation_letter",
	).From("user_scholarship us").
		Join("\"user\" u on us.user_id = u.id").
		Join("ethnic e on u.ethnic_id = e.id").
		Join("country c on u.country_id = c.id").
		Join("bank b on u.bank_id = b.id").
		Where(sq.Eq{"us.id": ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return entity.Applicant{}, err
	}

	var (
		applicant     entity.Applicant
		bytePhoto     []byte
		byteRecLetter []byte
	)
	row := a.db.QueryRowContext(ctx, query, args...)
	if err = row.Scan(
		&applicant.ID,
		&applicant.UserID,
		&applicant.User.Name,
		&bytePhoto,
		&applicant.User.Address,
		&applicant.User.BirthDate,
		&applicant.User.BirthPlace,
		&applicant.User.PhoneNo,
		&applicant.User.Gender,
		&applicant.User.EthnicID,
		&applicant.User.Ethnic.Name,
		&applicant.User.CountryID,
		&applicant.User.Country.Name,
		&applicant.User.BankID,
		&applicant.User.Bank.Name,
		&applicant.User.BankAccountNo,
		&applicant.User.BankAccountName,
		&applicant.User.CareerGoal,
		&applicant.Essay,
		&byteRecLetter,
	); err != nil {
		return entity.Applicant{}, err
	}

	if bytePhoto != nil {
		if err = json.Unmarshal(bytePhoto, &applicant.User.Photo); err != nil {
			return entity.Applicant{}, err
		}
	}

	if byteRecLetter != nil {
		if err = json.Unmarshal(byteRecLetter, &applicant.RecommendationLetter); err != nil {
			return entity.Applicant{}, err
		}
	}

	applicant.User.ID = applicant.UserID
	applicant.User.Ethnic.ID = applicant.User.EthnicID
	applicant.User.Country.ID = applicant.User.CountryID
	applicant.User.Bank.ID = int64(applicant.User.BankID)

	return applicant, nil
}

// NewApplicantRepository .
func NewApplicantRepository(db *sql.DB) business.ApplicantRepository {
	return applicantRepository{db: db}
}
