package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/sirupsen/logrus"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type userRepo struct {
	db *sql.DB
}

// Store .
func (u userRepo) Store(ctx context.Context, user entity.User) (entity.User, error) {
	user.CreatedAt = time.Now()

	query, args, err := sq.Insert("\"user\"").
		Columns("type",
			"email",
			"phone_no",
			"password",
			"status",
			"created_at",
		).Values(user.Type,
		user.Email,
		user.PhoneNo,
		user.Password,
		user.Status,
		user.CreatedAt,
	).PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return entity.User{}, err
	}

	_, err = u.db.ExecContext(ctx, query, args...)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// Fetch ...
func (u userRepo) Fetch(ctx context.Context, filter entity.UserFilter) ([]entity.User, string, error) {
	qSelect := sq.Select("u.id",
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
		"u.created_at",
		"e.name",
	).From("\"user\" u").
		LeftJoin("ethnic e on u.ethnic_id = e.id").
		OrderBy("created_at desc").
		PlaceholderFormat(sq.Dollar)

	if filter.Email != "" {
		qSelect = qSelect.Where(sq.Eq{"u.email": filter.Email})
	}

	if len(filter.IDs) > 0 {
		qSelect = qSelect.Where(sq.Eq{"u.id": filter.IDs})
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
		bytePhoto []byte
	)

	for rows.Next() {
		var (
			user       entity.User
			ethnicName sql.NullString
		)

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Type,
			&user.Email,
			&user.PhoneNo,
			&bytePhoto,
			&user.CompanyName,
			&user.Status,
			&user.CountryID,
			&user.PostalCode,
			&user.Address,
			&user.Gender,
			&user.EthnicID,
			&user.BirthDate,
			&user.BirthPlace,
			&user.BankID,
			&user.BankAccountNo,
			&user.BankAccountName,
			&user.CreatedAt,
			&ethnicName,
		); err != nil {
			return nil, "", err
		}

		cursor = user.CreatedAt
		if bytePhoto != nil {
			if err = json.Unmarshal(bytePhoto, &user.Photo); err != nil {
				return nil, "", err
			}
		}

		if ethnicName.Valid {
			user.Ethnic.Name = ethnicName.String
		}

		user.Ethnic.ID = user.EthnicID

		users = append(users, user)
	}

	cursorStr, err := encodeCursor(cursor)
	if err != nil {
		return nil, "", err
	}

	return users, cursorStr, nil
}

// Login is use just for login
// Different select columns with fetch
func (u userRepo) Login(ctx context.Context, email string) (entity.User, error) {
	query, args, err := sq.Select("id",
		"name",
		"type",
		"email",
		"status",
		"password",
		"photo",
	).From("\"user\"").
		Where(sq.Eq{"email": email}).
		OrderBy("created_at desc").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return entity.User{}, err
	}

	row := u.db.QueryRowContext(ctx, query, args...)
	var (
		user       entity.User
		bytesPhoto []byte
	)
	if err = row.Scan(&user.ID,
		&user.Name,
		&user.Type,
		&user.Email,
		&user.Status,
		&user.Password,
		&bytesPhoto,
	); err != nil {
		return entity.User{}, err
	}

	if bytesPhoto != nil {
		if err = json.Unmarshal(bytesPhoto, &user.Photo); err != nil {
			return entity.User{}, err
		}
	}

	return user, nil
}

// UpdateByID ...
// Update at table user
// insert into table card identity
// use transaction !!!!!!!!!!!
func (u userRepo) UpdateByID(ctx context.Context, ID int64, user entity.User) (entity.User, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return entity.User{}, err
	}

	var (
		timeNow     = time.Now()
		bytesImg    []byte
		errRollback error
	)

	bytesImg, err = json.Marshal(user.Photo)
	if err != nil {
		return entity.User{}, err
	}

	query, args, err := sq.Update("\"user\"").
		SetMap(sq.Eq{
			"name":              user.Name,
			"photo":             bytesImg,
			"company_name":      user.CompanyName,
			"country_id":        user.CountryID,
			"address":           user.Address,
			"postal_code":       user.PostalCode,
			"bank_id":           user.BankID,
			"bank_account_no":   user.BankAccountNo,
			"bank_account_name": user.BankAccountName,
			"status":            user.Status,
			"ethnic_id":         user.Ethnic.ID,
			"birth_date":        user.BirthDate,
			"birth_place":       user.BirthPlace,
			"gender":            user.Gender,
			"updated_at":        timeNow,
		}).Where(sq.Eq{"id": ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return entity.User{}, err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		return entity.User{}, err
	}

	qInsert := sq.Insert("card_identity").
		Columns("type",
			"no",
			"image",
			"user_id",
			"created_at",
		)
	for _, cardIdentity := range user.CardIdentities {
		bytesImg, err = json.Marshal(cardIdentity.Image)
		if err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				fmt.Println("Err rollback update profile at json marshal image : ", errRollback)
			}

			return entity.User{}, err
		}

		qInsert = qInsert.Values(
			cardIdentity.Type,
			cardIdentity.No,
			bytesImg,
			cardIdentity.UserID,
			timeNow,
		)
	}

	query, args, err = qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			fmt.Println("Err rollback update profile generate query insert card identity : ", errRollback)
		}

		return entity.User{}, err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			fmt.Println("Err rollback update profile exec insert card identity : ", errRollback)
		}

		return entity.User{}, err
	}

	if errCommit := tx.Commit(); errCommit != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			fmt.Println("Err rollback update profile at commit : ", errRollback)
		}
	}

	return user, nil
}

// SetStatus ....
// 1 active
// 0 inactive
func (u userRepo) SetStatus(ctx context.Context, ID int64, status int) error {
	query, args, err := sq.Update("\"user\"").SetMap(sq.Eq{
		"status":     status,
		"updated_at": time.Now(),
	}).Where(sq.Eq{"id": ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = u.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// ResetPassword ...
func (u userRepo) ResetPassword(ctx context.Context, email, password string) error {
	query, args, err := sq.Update("\"user\"").
		SetMap(sq.Eq{
			"password":   password,
			"updated_at": time.Now(),
		}).Where(sq.Eq{"email": email}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = u.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// SetupEducation .
func (u userRepo) SetupEducation(ctx context.Context, user entity.User) (entity.User, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return entity.User{}, err
	}

	var (
		timeNow     = time.Now()
		errRollback error
		byteDoc     []byte
	)

	// update user start
	// update status user to 3, status is set from service
	query, args, err := sq.Update("\"user\"").
		SetMap(sq.Eq{"status": user.Status,
			"career_goal":        user.CareerGoal,
			"study_country_goal": user.StudyCountryGoal.ID,
			"study_destination":  user.StudyDestination,
			"gap_year_reason":    user.GapYearReason,
			"updated_at":         timeNow,
		}).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": user.ID}).
		ToSql()
	if err != nil {
		return entity.User{}, err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		fmt.Println(err)
		return entity.User{}, err
	}
	// update user end

	// insert into user school
	qInsertUserSchool := sq.Insert("user_school").
		Columns("user_id",
			"school_id",
			"degree_id",
			"major_id",
			"enrollment_date",
			"graduation_date",
			"gpa",
			"created_at",
		)

	for _, us := range user.UserSchools {
		qInsertUserSchool = qInsertUserSchool.Values(user.ID,
			us.School.ID,
			us.Degree.ID,
			us.Major.ID,
			us.EnrollmentDate,
			us.GraduationDate,
			us.Gpa,
			timeNow,
		)
	}

	query, args, err = qInsertUserSchool.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			fmt.Println("Err rollback setup education insert user school : ", errRollback)
		}

		return entity.User{}, err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			fmt.Println("Err rollback setup education exec insert user school : ", errRollback)
		}

		return entity.User{}, err
	}

	// insert user school end

	// insert user document start
	if len(user.UserDocuments) > 0 {
		qInsertUserDocument := sq.Insert("user_document").Columns("user_id", "document", "created_at")

		for _, ud := range user.UserDocuments {
			byteDoc, err = json.Marshal(ud.Document)
			if err != nil {
				if errRollback = tx.Rollback(); errRollback != nil {
					fmt.Println("Err rollback setup education marshal document : ", errRollback)
				}

				return entity.User{}, err
			}

			qInsertUserDocument = qInsertUserDocument.Values(user.ID, byteDoc, timeNow)
		}

		query, args, err = qInsertUserDocument.PlaceholderFormat(sq.Dollar).ToSql()
		if err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				fmt.Println("Err rollback setup education generate sql user document : ", errRollback)
			}

			return entity.User{}, err
		}

		if _, err = tx.ExecContext(ctx, query, args...); err != nil {
			if errRollback = tx.Rollback(); errRollback != nil {
				fmt.Println("Err rollback setup education exec user document : ", errRollback)
			}

			return entity.User{}, err
		}
	}
	// insert user document end

	if err = tx.Commit(); err != nil {
		if errRollback = tx.Rollback(); errRollback != nil {
			fmt.Println("Err commit setup education : ", errRollback)
		}

		return entity.User{}, err
	}

	return user, nil
}

// GetDocuments .
func (u userRepo) GetDocuments(ctx context.Context, userID int64) ([]entity.UserDocument, error) {
	query, args, err := sq.Select("id", "user_id", "document").
		From("user_document").
		Where(sq.Eq{"user_id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := u.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	userDocs := make([]entity.UserDocument, 0)

	for rows.Next() {
		var (
			ud      entity.UserDocument
			byteDoc []byte
		)

		if err = rows.Scan(
			&ud.ID,
			&ud.UserID,
			&byteDoc,
		); err != nil {
			return nil, err
		}

		if byteDoc != nil {
			if err = json.Unmarshal(byteDoc, &ud.Document); err != nil {
				return nil, err
			}
		}

		userDocs = append(userDocs, ud)
	}

	return userDocs, nil
}

// NewUserRepository .
func NewUserRepository(db *sql.DB) business.UserRepository {
	return userRepo{db: db}
}
