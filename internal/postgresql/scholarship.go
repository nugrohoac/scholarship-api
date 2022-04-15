package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/util"

	"github.com/sirupsen/logrus"

	sq "github.com/Masterminds/squirrel"
)

type scholarshipRepo struct {
	db              *sql.DB
	deadlinePayment int
}

func (s scholarshipRepo) FetchScholarshipBackoffice(ctx context.Context, filter entity.ScholarshipFilterBackoffice) ([]entity.Scholarship, string, error) {
	qSelect := sq.Select("id",
		"sponsor_id",
		"name",
		"amount",
		"status",
		"image",
		"awardee",
		"current_applicant",
		"application_start",
		"application_end",
		"announcement_date",
		"eligibility_description",
		"subsidy_description",
		"funding_start",
		"funding_end",
		"created_at",
	).From("scholarship").
		PlaceholderFormat(sq.Dollar).
		OrderBy("created_at desc")

	if filter.Limit > 0 {
		qSelect = qSelect.Limit(uint64(filter.Limit))
	}

	if filter.Cursor != "" {
		cursor, err := decodeCursor(filter.Cursor)
		if err != nil {
			return nil, "", err
		}

		qSelect = qSelect.Where(sq.Lt{"created_at": cursor})
	}

	if filter.SearchText != "" {
		searchText := "%" + filter.SearchText + "%"
		qSelect = qSelect.Where(sq.Or{
			sq.Like{"name": searchText},
		})
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		scholarships = make([]entity.Scholarship, 0)
		cursor       time.Time
		cursorStr    string
		byteImg      []byte
	)

	for rows.Next() {
		var scholarship entity.Scholarship

		if err = rows.Scan(
			&scholarship.ID,
			&scholarship.SponsorID,
			&scholarship.Name,
			&scholarship.Amount,
			&scholarship.Status,
			&byteImg,
			&scholarship.Awardee,
			&scholarship.CurrentApplicant,
			&scholarship.ApplicationStart,
			&scholarship.ApplicationEnd,
			&scholarship.AnnouncementDate,
			&scholarship.EligibilityDescription,
			&scholarship.SubsidyDescription,
			&scholarship.FundingStart,
			&scholarship.FundingEnd,
			&scholarship.CreatedAt,
		); err != nil {
			return nil, "", err
		}

		if byteImg != nil {
			if err = json.Unmarshal(byteImg, &scholarship.Image); err != nil {
				return nil, "", err
			}
		}

		cursor = scholarship.CreatedAt
		scholarship.TextStatus = util.GetNameStatus(scholarship.Status)
		scholarships = append(scholarships, scholarship)
	}

	cursorStr, err = encodeCursor(cursor)
	if err != nil {
		return nil, "", err
	}

	return scholarships, cursorStr, nil
}

// Create ...
func (s scholarshipRepo) Create(ctx context.Context, scholarship entity.Scholarship) (entity.Scholarship, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return entity.Scholarship{}, err
	}

	byteImage, err := json.Marshal(scholarship.Image)
	if err != nil {
		return entity.Scholarship{}, err
	}

	var (
		timeNow     = time.Now()
		errRollback error
	)

	query, args, err := sq.Insert("scholarship").
		Columns("sponsor_id",
			"name",
			"amount",
			"status",
			"image",
			"awardee",
			"application_start",
			"application_end",
			"announcement_date",
			"eligibility_description",
			"subsidy_description",
			"funding_start",
			"funding_end",
			"created_at").
		Values(scholarship.SponsorID,
			scholarship.Name,
			scholarship.Amount,
			scholarship.Status,
			byteImage,
			scholarship.Awardee,
			scholarship.ApplicationStart,
			scholarship.ApplicationEnd,
			scholarship.AnnouncementDate,
			scholarship.EligibilityDescription,
			scholarship.SubsidyDescription,
			scholarship.FundingStart,
			scholarship.FundingEnd,
			timeNow,
		).Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return entity.Scholarship{}, err
	}

	row := tx.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&scholarship.ID); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return entity.Scholarship{}, err
	}

	// Insert requirements start
	if len(scholarship.Requirements) > 0 {
		qInsertRequirement := sq.Insert("requirement").
			Columns("scholarship_id",
				"type",
				"name",
				"value",
				"created_at",
			)

		for _, req := range scholarship.Requirements {
			qInsertRequirement = qInsertRequirement.Values(scholarship.ID,
				req.Type,
				req.Name,
				req.Value,
				timeNow,
			)
		}

		query, args, err = qInsertRequirement.PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return entity.Scholarship{}, err
		}

		if _, err = tx.ExecContext(ctx, query, args...); err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return entity.Scholarship{}, err
		}
	}
	// Insert requirements end

	// insert requirement description start
	if len(scholarship.RequirementDescriptions) > 0 {
		qInsertRequirementDescription := sq.Insert("requirement_description").
			Columns("scholarship_id", "description")

		for _, reqDescription := range scholarship.RequirementDescriptions {
			qInsertRequirementDescription = qInsertRequirementDescription.
				Values(scholarship.ID, reqDescription)
		}

		query, args, err = qInsertRequirementDescription.PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return entity.Scholarship{}, err
		}

		if _, err = tx.ExecContext(ctx, query, args...); err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return entity.Scholarship{}, err
		}
	}
	// insert requirements description end

	// Create payment start
	scholarship.Payment.Deadline = timeNow.Add(time.Duration(s.deadlinePayment) * time.Hour)
	scholarship.Payment.CreatedAt = timeNow
	scholarship.Payment.ScholarshipID = scholarship.ID

	query, args, err = sq.Insert("payment").
		Columns("scholarship_id",
			"deadline",
			"transfer_date",
			"created_at",
		).Values(
		scholarship.ID,
		scholarship.Payment.Deadline,
		scholarship.Payment.TransferDate,
		scholarship.Payment.CreatedAt,
	).Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return entity.Scholarship{}, err
	}

	row = tx.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&scholarship.Payment.ID); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return entity.Scholarship{}, err
	}
	// Create payment end

	if errCommit := tx.Commit(); errCommit != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return entity.Scholarship{}, errCommit
	}

	scholarship.CreatedAt = timeNow
	return scholarship, nil
}

// Fetch ....
func (s scholarshipRepo) Fetch(ctx context.Context, filter entity.ScholarshipFilter) ([]entity.Scholarship, string, error) {
	qSelect := sq.Select("id",
		"sponsor_id",
		"name",
		"amount",
		"status",
		"image",
		"awardee",
		"current_applicant",
		"application_start",
		"application_end",
		"announcement_date",
		"eligibility_description",
		"subsidy_description",
		"funding_start",
		"funding_end",
		"created_at",
	).From("scholarship").
		PlaceholderFormat(sq.Dollar).
		OrderBy("created_at desc")

	if filter.Limit > 0 {
		qSelect = qSelect.Limit(filter.Limit)
	}

	if filter.Cursor != "" {
		cursor, err := decodeCursor(filter.Cursor)
		if err != nil {
			return nil, "", err
		}

		qSelect = qSelect.Where(sq.Lt{"created_at": cursor})
	}

	if filter.SponsorID > 0 {
		qSelect = qSelect.Where(sq.Eq{"sponsor_id": filter.SponsorID})
	}

	if len(filter.Status) > 0 {
		qSelect = qSelect.Where(sq.Eq{"status": filter.Status})
	}

	if filter.Name != "" {
		likeName := "%" + strings.ToLower(filter.Name) + "%"
		qSelect = qSelect.Where(sq.Like{"LOWER(name)": likeName})
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		scholarships = make([]entity.Scholarship, 0)
		cursor       time.Time
		cursorStr    string
		byteImg      []byte
	)

	for rows.Next() {
		var scholarship entity.Scholarship

		if err = rows.Scan(
			&scholarship.ID,
			&scholarship.SponsorID,
			&scholarship.Name,
			&scholarship.Amount,
			&scholarship.Status,
			&byteImg,
			&scholarship.Awardee,
			&scholarship.CurrentApplicant,
			&scholarship.ApplicationStart,
			&scholarship.ApplicationEnd,
			&scholarship.AnnouncementDate,
			&scholarship.EligibilityDescription,
			&scholarship.SubsidyDescription,
			&scholarship.FundingStart,
			&scholarship.FundingEnd,
			&scholarship.CreatedAt,
		); err != nil {
			return nil, "", err
		}

		if byteImg != nil {
			if err = json.Unmarshal(byteImg, &scholarship.Image); err != nil {
				return nil, "", err
			}
		}

		cursor = scholarship.CreatedAt
		scholarships = append(scholarships, scholarship)
	}

	cursorStr, err = encodeCursor(cursor)
	if err != nil {
		return nil, "", err
	}

	return scholarships, cursorStr, nil
}

// GetByID ...
func (s scholarshipRepo) GetByID(ctx context.Context, ID int64) (entity.Scholarship, error) {
	query, args, err := sq.Select("s.id",
		"s.sponsor_id",
		"s.name",
		"s.amount",
		"s.status",
		"s.image",
		"s.awardee",
		"s.current_applicant",
		"s.application_start",
		"s.application_end",
		"s.announcement_date",
		"s.eligibility_description",
		"s.subsidy_description",
		"s.funding_start",
		"s.funding_end",
		"s.created_at",
		"r.id",
		"r.name",
		"r.type",
		"r.value",
		"u.id as sponsor_id",
		"u.name as sponsor_name",
		"u.email as sponsor_email",
		"u.phone_no as sponsor_phone_number",
		"u.company_name as sponsor_company_name",
		"u.photo as sponsor_image",
	).From("scholarship s").
		LeftJoin("requirement r on s.id = r.scholarship_id").
		LeftJoin("\"user\" u on u.id = s.sponsor_id").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"s.id": ID}).
		ToSql()
	if err != nil {
		return entity.Scholarship{}, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return entity.Scholarship{}, err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		scholarship  entity.Scholarship
		byteImage    []byte
		sponsorImage []byte
		reqID        sql.NullInt64
		name         sql.NullString
		_type        sql.NullString
		value        sql.NullString
	)

	for rows.Next() {
		var requirement entity.Requirement
		var sponsor entity.User

		if err = rows.Scan(
			&scholarship.ID,
			&scholarship.SponsorID,
			&scholarship.Name,
			&scholarship.Amount,
			&scholarship.Status,
			&byteImage,
			&scholarship.Awardee,
			&scholarship.CurrentApplicant,
			&scholarship.ApplicationStart,
			&scholarship.ApplicationEnd,
			&scholarship.AnnouncementDate,
			&scholarship.EligibilityDescription,
			&scholarship.SubsidyDescription,
			&scholarship.FundingStart,
			&scholarship.FundingEnd,
			&scholarship.CreatedAt,
			&reqID,
			&name,
			&_type,
			&value,
			&sponsor.ID,
			&sponsor.Name,
			&sponsor.Email,
			&sponsor.PhoneNo,
			&sponsor.CompanyName,
			&sponsorImage,
		); err != nil {
			return entity.Scholarship{}, err
		}

		if byteImage != nil {
			if err = json.Unmarshal(byteImage, &scholarship.Image); err != nil {
				return entity.Scholarship{}, err
			}
		}

		if sponsorImage != nil {
			if err = json.Unmarshal(sponsorImage, &sponsor.Photo); err != nil {
				return entity.Scholarship{}, err
			}
		}

		if name.Valid {
			requirement.Name = name.String
		}

		if _type.Valid {
			requirement.Type = _type.String
		}

		if value.Valid {
			requirement.Value = value.String
		}

		if reqID.Valid {
			requirement.ID = reqID.Int64
		}

		scholarship.Requirements = append(scholarship.Requirements, requirement)
		scholarship.Sponsor = sponsor
	}

	return scholarship, nil
}

// Apply .
func (s scholarshipRepo) Apply(ctx context.Context, userID, scholarshipID int64, applicant int, essay string, recommendationLetter entity.Image) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var (
		timeNow           = time.Now()
		errRollback       error
		userScholarshipID int64
		byteRecLetter     []byte
	)

	byteRecLetter, err = json.Marshal(recommendationLetter)
	if err != nil {
		return err
	}

	// update current applicant
	query, args, err := sq.Update("scholarship").
		SetMap(sq.Eq{"current_applicant": applicant, "updated_at": timeNow}).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": scholarshipID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	// insert into user scholarship
	query, args, err = sq.Insert("user_scholarship").
		Columns("scholarship_id",
			"user_id",
			"status",
			"essay",
			"recommendation_letter",
			"created_at",
		).Values(scholarshipID,
		userID,
		0,
		essay,
		byteRecLetter,
		timeNow,
	).Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return err
	}

	row := tx.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&userScholarshipID); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return err
	}

	if errCommit := tx.Commit(); errCommit != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return errCommit
	}

	return nil
}

// CheckApply .
func (s scholarshipRepo) CheckApply(ctx context.Context, userID, scholarshipID int64) (bool, int, error) {
	query, args, err := sq.Select("status").
		From("user_scholarship").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"scholarship_id": scholarshipID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, 0, err
	}

	var status int
	row := s.db.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&status); err != nil {
		if err == sql.ErrNoRows {
			return false, 0, nil
		}

		return false, 0, err
	}

	return true, status, nil
}

// MyScholarship .
/* scholarship was applied by specific student
 */
func (s scholarshipRepo) MyScholarship(ctx context.Context, userID int64, filter entity.ScholarshipFilter) ([]entity.Applicant, string, error) {
	qSelect := sq.Select("us.id",
		"us.scholarship_id",
		"us.user_id",
		"us.essay",
		"us.recommendation_letter",
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
		"u.bank_id",
		"u.bank_account_no",
		"u.bank_account_name",
		"s.id",
		"s.sponsor_id",
		"s.name",
		"s.amount",
		"s.awardee",
		"s.current_applicant",
		"s.eligibility_description",
		"s.subsidy_description",
		"s.application_start",
		"s.application_end",
		"s.announcement_date",
		"s.funding_start",
		"s.funding_end",
		"s.status",
	).From("user_scholarship us").
		Join("\"user\" u on u.id = us.user_id").
		Join("scholarship s on s.id = us.scholarship_id").
		Where(sq.Eq{"us.user_id": userID}).
		OrderBy("us.created_at desc").
		PlaceholderFormat(sq.Dollar)

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

	if filter.Name != "" {
		name := "%" + strings.ToLower(filter.Name) + "%"
		qSelect = qSelect.Where(sq.Like{"LOWER(s.name)": name})
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, "", err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		cursorTime time.Time
		cursorStr  string
		applicants = make([]entity.Applicant, 0)
	)

	for rows.Next() {
		var (
			applicant     entity.Applicant
			byteRecLetter []byte
			bytePhoto     []byte
		)

		if err = rows.Scan(
			&applicant.ID,
			&applicant.ScholarshipID,
			&applicant.UserID,
			&applicant.Essay,
			&byteRecLetter,
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
			&applicant.User.Bank.ID,
			&applicant.User.BankAccountNo,
			&applicant.User.BankAccountName,
			&applicant.Scholarship.ID,
			&applicant.Scholarship.SponsorID,
			&applicant.Scholarship.Name,
			&applicant.Scholarship.Amount,
			&applicant.Scholarship.Awardee,
			&applicant.Scholarship.CurrentApplicant,
			&applicant.Scholarship.EligibilityDescription,
			&applicant.Scholarship.SubsidyDescription,
			&applicant.Scholarship.ApplicationStart,
			&applicant.Scholarship.ApplicationEnd,
			&applicant.Scholarship.AnnouncementDate,
			&applicant.Scholarship.FundingStart,
			&applicant.Scholarship.FundingEnd,
			&applicant.Scholarship.Status,
		); err != nil {
			return nil, "", err
		}

		if bytePhoto != nil {
			if err = json.Unmarshal(bytePhoto, &applicant.User.Photo); err != nil {
				return nil, "", err
			}
		}

		if byteRecLetter != nil {
			if err = json.Unmarshal(byteRecLetter, &applicant.RecommendationLetter); err != nil {
				return nil, "", err
			}
		}

		cursorTime = applicant.ApplyDate
		applicants = append(applicants, applicant)
	}

	cursorStr, err = encodeCursor(cursorTime)
	if err != nil {
		return nil, "", err
	}

	return applicants, cursorStr, nil
}

// ChangeStatus .
func (s scholarshipRepo) ChangeStatus(ctx context.Context, ID int64, status int) error {
	query, args, err := sq.Update("scholarship").
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

	if _, err = s.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (s scholarshipRepo) ApprovedScholarship(ctx context.Context, scholarshipID int64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var (
		timeNow     = time.Now()
		errRollback error
	)

	query, args, err := sq.Update("scholarship").
		SetMap(sq.Eq{"status": util.APPROVE, "updated_at": timeNow}).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": scholarshipID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	if errCommit := tx.Commit(); errCommit != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return errCommit
	}

	return nil
}

// NewScholarshipRepository ...
func NewScholarshipRepository(db *sql.DB, deadlinePayment int) business.ScholarshipRepository {
	return scholarshipRepo{
		db:              db,
		deadlinePayment: deadlinePayment,
	}
}
