package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	sq "github.com/Masterminds/squirrel"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

type scholarshipRepo struct {
	db              *sql.DB
	deadlinePayment int
}

// Create ...
func (s scholarshipRepo) Create(ctx context.Context, scholarship sa.Scholarship) (sa.Scholarship, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return sa.Scholarship{}, err
	}

	byteImage, err := json.Marshal(scholarship.Image)
	if err != nil {
		return sa.Scholarship{}, err
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
		return sa.Scholarship{}, err
	}

	row := tx.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&scholarship.ID); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return sa.Scholarship{}, err
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

			return sa.Scholarship{}, err
		}

		if _, err = tx.ExecContext(ctx, query, args...); err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return sa.Scholarship{}, err
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

			return sa.Scholarship{}, err
		}

		if _, err = tx.ExecContext(ctx, query, args...); err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return sa.Scholarship{}, err
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

		return sa.Scholarship{}, err
	}

	row = tx.QueryRowContext(ctx, query, args...)
	if err = row.Scan(&scholarship.Payment.ID); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return sa.Scholarship{}, err
	}
	// Create payment end

	if errCommit := tx.Commit(); errCommit != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			logrus.Error(errRollback)
		}

		return sa.Scholarship{}, errCommit
	}

	scholarship.CreatedAt = timeNow
	return scholarship, nil
}

// Fetch ....
func (s scholarshipRepo) Fetch(ctx context.Context, filter sa.ScholarshipFilter) ([]sa.Scholarship, string, error) {
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
		scholarships = make([]sa.Scholarship, 0)
		cursor       time.Time
		cursorStr    string
		byteImg      []byte
	)

	for rows.Next() {
		var scholarship sa.Scholarship

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
func (s scholarshipRepo) GetByID(ctx context.Context, ID int64) (sa.Scholarship, error) {
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
		"r.name",
		"r.type",
		"r.value",
	).From("scholarship s").
		Join("requirement r on s.id = r.scholarship_id").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"s.id": ID}).
		ToSql()
	if err != nil {
		return sa.Scholarship{}, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return sa.Scholarship{}, err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	var (
		scholarship sa.Scholarship
		byteImage   []byte
	)

	for rows.Next() {
		var requirement sa.Requirement

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
			&requirement.Name,
			&requirement.Type,
			&requirement.Value,
		); err != nil {
			return sa.Scholarship{}, err
		}

		if byteImage != nil {
			if err = json.Unmarshal(byteImage, &scholarship.Image); err != nil {
				return sa.Scholarship{}, err
			}
		}
		scholarship.Requirements = append(scholarship.Requirements, requirement)
	}

	return scholarship, nil
}

// Apply .
func (s scholarshipRepo) Apply(ctx context.Context, userID, scholarshipID int64, applicant int, documents []sa.Document) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var (
		timeNow           = time.Now()
		errRollback       error
		userScholarshipID int64
		byteDoc           []byte
	)

	// update current applicant
	query, args, err := sq.Update("scholarship").
		SetMap(sq.Eq{"current_applicant": applicant, "updated_at": timeNow}).
		PlaceholderFormat(sq.Dollar).
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
			"created_at",
		).Values(scholarshipID,
		userID,
		0,
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

	// insert user scholarship document
	if len(documents) > 0 {
		qInsert := sq.Insert("user_scholarship_document").
			Columns("user_scholarship_id", "name", "value", "created_at")

		for _, doc := range documents {
			byteDoc, err = json.Marshal(doc.Value)
			if err != nil {
				if errRollback = tx.Rollback(); errRollback != nil {
					logrus.Error(errRollback)
				}

				return err
			}

			qInsert = qInsert.Values(userScholarshipID, doc.Name, byteDoc, timeNow)
		}

		query, args, err = qInsert.PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return err
		}

		if _, err = tx.ExecContext(ctx, query, args...); err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				logrus.Error(errRollback)
			}

			return err
		}
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
func NewScholarshipRepository(db *sql.DB, deadlinePayment int) sa.ScholarshipRepository {
	return scholarshipRepo{
		db:              db,
		deadlinePayment: deadlinePayment,
	}
}
