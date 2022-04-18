package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
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
	//order by aps.value desc nulls last
	sort := "order by us.created_at desc"
	if filter.Sort == "score" {
		sort = "order by aps.value desc nulls last"
	}

	qSelect := sq.Select("us.id applicant_id",
		"us.scholarship_id scholarship_id",
		"us.user_id user_id",
		"us.status applicant_status",
		"us.created_at applicant_created_at",
		"u.id _user_id",
		"u.name user_name",
		"u.type user_type",
		"u.email user_email",
		"u.phone_no user_phone_no",
		"u.photo user_photo",
		"u.company_name user_company_name",
		"u.status user_status",
		"u.country_id user_country_id",
		"u.postal_code user_postal_code",
		"u.address user_address",
		"u.gender user_gender",
		"u.ethnic_id user_ethnic_id",
		"u.birth_date user_birth_date",
		"u.birth_place user_birth_place",
		"u.bank_id user_bank_id",
		"u.bank_account_no user_bank_account_no",
		"u.bank_account_name user_bank_account_name",
		"u.ethnic_id _user_ethnic_id",
		"e.name ethnic_name",
		"aps.name score_name",
		"aps.value score_value",
		"row_number () over(%s) row_id",
	).From("user_scholarship us").
		Join("\"user\" u on us.user_id  = u.id").
		LeftJoin("ethnic e on e.id = u.ethnic_id").
		LeftJoin("(select applicant_id, name, value from applicant_score where name='total') aps on aps.applicant_id = us.id ").
		LeftJoin("scholarship s on s.id = us.scholarship_id").
		Where(sq.Eq{"us.scholarship_id": filter.ScholarshipID}).
		PlaceholderFormat(sq.Dollar)

	if len(filter.Status) > 0 {
		qSelect = qSelect.Where(sq.Eq{"us.status": filter.Status})
	}

	if filter.SponsorID > 0 {
		qSelect = qSelect.Where(sq.Eq{"s.sponsor_id": filter.SponsorID})
	}

	if filter.UserID > 0 {
		qSelect = qSelect.Where(sq.Eq{"us.user_id": filter.UserID})
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, "", err
	}

	query = fmt.Sprintf(query, sort)

	finalQuery := `select 
		res.applicant_id,
		res.scholarship_id,
		res.user_id _user_id,
		res.applicant_status,
		res.applicant_created_at,
		res.user_id,
		res.user_name,
		res.user_type,
		res.user_email,
		res.user_phone_no,
		res.user_photo,
		res.user_company_name,
		res.user_status,
		res.user_country_id,
		res.user_postal_code,
		res.user_address,
		res.user_gender,
		res.user_ethnic_id,
		res.user_birth_date,
		res.user_birth_place,
		res.user_bank_id,
		res.user_bank_account_no,
		res.user_bank_account_name,
		res.user_ethnic_id eth_id,
		res.ethnic_name,
		res.score_name,
		res.score_value,
		res.row_id from (` + query + `) as res`

	if filter.Cursor != "" {
		cursorRow, err := decodeCursorRow(filter.Cursor)
		if err != nil {
			return nil, "", err
		}

		finalQuery = finalQuery + " WHERE res.row_id > " + cursorRow
	}

	if filter.Limit > 0 {
		finalQuery = finalQuery + " LIMIT " + strconv.Itoa(int(filter.Limit))
	}

	rows, err := a.db.QueryContext(ctx, finalQuery, args...)
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
		cursor     string
		rowNumber  int64
	)

	for rows.Next() {
		var (
			applicant  entity.Applicant
			bytePhoto  []byte
			ethnicName sql.NullString
			scoreName  sql.NullString
			scoreValue sql.NullInt32
			score      entity.ApplicantScore
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
			&scoreName,
			&scoreValue,
			&rowNumber,
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

		score.Name = `total`
		if scoreName.Valid {
			score.Name = scoreName.String
		}

		score.Value = 0
		if scoreValue.Valid {
			score.Value = scoreValue.Int32
		}

		applicant.UserID = applicant.User.ID
		applicant.User.Ethnic.ID = applicant.User.EthnicID
		applicant.Scores = []entity.ApplicantScore{score}
		cursor = strconv.Itoa(int(rowNumber))

		applicants = append(applicants, applicant)
	}

	cursor = encodeCursorRow(cursor)

	return applicants, cursor, nil
}

// GetByID .
// ID is reference to user scholarship id
func (a applicantRepository) GetByID(ctx context.Context, ID int64) (entity.Applicant, error) {
	query, args, err := sq.Select(
		"us.id",
		"us.user_id",
		"us.scholarship_id",
		"us.status",
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
		&applicant.ScholarshipID,
		&applicant.Status,
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

// SubmitAssessment .
func (a applicantRepository) SubmitAssessment(ctx context.Context, ApplicantID int64, eligibilities []entity.ApplicantEligibility, scores []entity.ApplicantScore) error {
	timeNow := time.Now()

	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	// there are chance eligibilities is nil
	if len(eligibilities) > 0 {
		qInsert := sq.Insert("applicant_eligibility").
			Columns("applicant_id",
				"requirement_id",
				"value",
				"created_at",
			).PlaceholderFormat(sq.Dollar)

		for _, eligibility := range eligibilities {
			qInsert = qInsert.Values(
				ApplicantID,
				eligibility.RequirementID,
				eligibility.Value,
				timeNow,
			)
		}

		query, args, err := qInsert.ToSql()
		if err != nil {
			return err
		}

		if _, err = tx.ExecContext(ctx, query, args...); err != nil {
			return err
		}
	}

	// insert into applicant score
	qInsert := sq.Insert("applicant_score").
		Columns("applicant_id",
			"name",
			"value",
		).PlaceholderFormat(sq.Dollar)

	for _, score := range scores {
		qInsert = qInsert.Values(ApplicantID,
			score.Name,
			score.Value,
		)
	}

	query, args, err := qInsert.ToSql()
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			logrus.Error(err)
		}

		return err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			logrus.Error(err)
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			logrus.Error(err)
		}

		return err
	}

	return nil
}

// UpdateStatus .
func (a applicantRepository) UpdateStatus(ctx context.Context, ID int64, status int32) error {
	query, args, err := sq.Update("user_scholarship").
		SetMap(sq.Eq{
			"status":     status,
			"updated_at": time.Now(),
		}).
		Where(sq.Eq{"id": ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = a.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// SetStatusWaitForConfirmation .
// please readme.md for detail status user_scholarship
func (a applicantRepository) SetStatusWaitForConfirmation(ctx context.Context, userIDs []int64, scholarshipID int64) error {
	query, args, err := sq.Update("user_scholarship").
		Set("status", 3).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"user_id": userIDs}).
		Where(sq.Eq{"scholarship_id": scholarshipID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = a.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// SetStatusConfirmation .
// please readme.md for detail status user_scholarship
func (a applicantRepository) SetStatusConfirmation(ctx context.Context, userID, scholarshipID int64) error {
	query, args, err := sq.Update("user_scholarship").
		Set("status", 4).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"scholarship_id": scholarshipID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = a.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// NewApplicantRepository .
func NewApplicantRepository(db *sql.DB) business.ApplicantRepository {
	return applicantRepository{db: db}
}
