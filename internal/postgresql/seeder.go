package postgresql

import (
	"database/sql"
	"encoding/json"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"testing"

	"github.com/stretchr/testify/require"

	sq "github.com/Masterminds/squirrel"
)

// SeedBanks ...
func SeedBanks(db *sql.DB, t *testing.T, banks []entity.Bank) {
	qInsert := sq.Insert("bank").Columns("name", "code", "created_at")

	for _, bank := range banks {
		qInsert = qInsert.Values(bank.Name, bank.Code, bank.CreatedAt)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedCountries ...
func SeedCountries(db *sql.DB, t *testing.T, countries []entity.Country) {
	qInsert := sq.Insert("country").Columns("name", "created_at")

	for _, country := range countries {
		qInsert = qInsert.Values(country.Name, country.CreatedAt)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedUsers ...
func SeedUsers(db *sql.DB, t *testing.T, users []entity.User) {
	qInsert := sq.Insert("\"user\"").
		Columns("id",
			"name",
			"type",
			"email",
			"phone_no",
			"photo",
			"company_name",
			"password",
			"status",
			"country_id",
			"postal_code",
			"address",
			"gender",
			"ethnic_id",
			"birth_date",
			"birth_place",
			"bank_id",
			"bank_account_no",
			"bank_account_name",
			"created_at")

	for _, user := range users {
		bytePhoto, err := json.Marshal(user.Photo)
		require.NoError(t, err)

		qInsert = qInsert.Values(
			user.ID,
			user.Name,
			user.Type,
			user.Email,
			user.PhoneNo,
			bytePhoto,
			user.CompanyName,
			user.Password,
			user.Status,
			user.CountryID,
			user.PostalCode,
			user.Address,
			user.Gender,
			user.Ethnic.ID,
			user.BirthDate,
			user.BirthPlace,
			user.BankID,
			user.BankAccountNo,
			user.BankAccountName,
			user.CreatedAt,
		)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedCardIdentities ...
func SeedCardIdentities(db *sql.DB, t *testing.T, cardIdentities []entity.CardIdentity) {
	var (
		byteImg []byte
		err     error
	)

	qInsert := sq.Insert("card_identity").
		Columns("id",
			"type",
			"no",
			"image",
			"user_id",
			"created_at",
		)

	for _, cardIdentity := range cardIdentities {
		byteImg, err = json.Marshal(cardIdentity.Image)
		require.NoError(t, err)

		qInsert = qInsert.Values(
			cardIdentity.ID,
			cardIdentity.Type,
			cardIdentity.No,
			byteImg,
			cardIdentity.UserID,
			cardIdentity.CreatedAt,
		)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedScholarship ...
func SeedScholarship(db *sql.DB, t *testing.T, scholarships []entity.Scholarship) {
	qInsert := sq.Insert("scholarship").
		Columns("id",
			"sponsor_id",
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
			"created_at",
		)

	qInsertReqDesc := sq.Insert("requirement_description").
		Columns("scholarship_id", "description")

	for _, scholarship := range scholarships {
		byteImage, err := json.Marshal(scholarship.Image)
		require.NoError(t, err)

		qInsert = qInsert.Values(scholarship.ID,
			scholarship.SponsorID,
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
			scholarship.CreatedAt,
		)

		for _, desc := range scholarship.RequirementDescriptions {
			qInsertReqDesc = qInsertReqDesc.Values(scholarship.ID, desc)
		}
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)

	query, args, err = qInsertReqDesc.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)
	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedRequirements ...
func SeedRequirements(db *sql.DB, t *testing.T, requirements []entity.Requirement) {
	qInsert := sq.Insert("requirement").
		Columns("id",
			"scholarship_id",
			"type",
			"name",
			"value",
			"created_at",
		)

	for _, req := range requirements {
		qInsert = qInsert.Values(req.ID,
			req.ScholarshipID,
			req.Type,
			req.Name,
			req.Value,
			req.CreatedAt,
		)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedPayments ...
func SeedPayments(db *sql.DB, t *testing.T, payments []entity.Payment) {
	qInsert := sq.Insert("payment").
		Columns("id",
			"scholarship_id",
			"deadline",
			"transfer_date",
			"bank_account_name",
			"bank_account_no",
			"image",
			"created_at",
		)

	for _, payment := range payments {
		byteImage, err := json.Marshal(payment.Image)
		require.NoError(t, err)

		qInsert = qInsert.Values(payment.ID,
			payment.ScholarshipID,
			payment.Deadline,
			payment.TransferDate,
			payment.BankAccountName,
			payment.BankAccountNo,
			byteImage,
			payment.CreatedAt,
		)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedMajors ...
func SeedMajors(db *sql.DB, t *testing.T, majors []entity.Major) {
	qInsert := sq.Insert("major").Columns("id", "name", "created_at")

	for _, major := range majors {
		qInsert = qInsert.Values(major.ID, major.Name, major.CreatedAt)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedSchools .
func SeedSchools(db *sql.DB, t *testing.T, schools []entity.School) {
	qInsert := sq.Insert("school").
		Columns("id",
			"name",
			"type",
			"address",
			"status",
			"created_at",
			"created_by",
		)

	for _, school := range schools {
		qInsert = qInsert.Values(school.ID,
			school.Name,
			school.Type,
			school.Address,
			school.Status,
			school.CreatedAt,
			school.CreatedBy,
		)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedDegrees .
func SeedDegrees(db *sql.DB, t *testing.T, degrees []entity.Degree) {
	qInsert := sq.Insert("degree").
		Columns("id",
			"name",
			"created_at",
		)

	for _, degree := range degrees {
		qInsert = qInsert.Values(
			degree.ID,
			degree.Name,
			degree.CreatedAt,
		)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedBankTransfer ....
func SeedBankTransfer(db *sql.DB, t *testing.T, bankTransfers ...entity.BankTransfer) {
	qInsert := sq.Insert("bank_transfer").
		Columns("id",
			"name",
			"account_name",
			"account_no",
			"image",
			"created_at",
		)

	for _, bt := range bankTransfers {
		byteImage, err := json.Marshal(bt.Image)
		require.NoError(t, err)

		qInsert = qInsert.Values(bt.ID,
			bt.Name,
			bt.AccountName,
			bt.AccountNo,
			byteImage,
			bt.CreatedAt,
		)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}

// SeedEthnics .
func SeedEthnics(db *sql.DB, t *testing.T, ethnics []entity.Ethnic) {
	qInsert := sq.Insert("ethnic").Columns("id", "name")

	for _, ethnic := range ethnics {
		qInsert = qInsert.Values(ethnic.ID, ethnic.Name)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}
