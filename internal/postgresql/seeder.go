package postgresql

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	sq "github.com/Masterminds/squirrel"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

// SeedBanks ...
func SeedBanks(db *sql.DB, t *testing.T, banks []sa.Bank) {
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
func SeedCountries(db *sql.DB, t *testing.T, countries []sa.Country) {
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
func SeedUsers(db *sql.DB, t *testing.T, users []sa.User) {
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
			"ethnic",
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
			user.Ethnic,
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
func SeedCardIdentities(db *sql.DB, t *testing.T, cardIdentities []sa.CardIdentity) {
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
func SeedScholarship(db *sql.DB, t *testing.T, scholarships []sa.Scholarship) {
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
func SeedRequirements(db *sql.DB, t *testing.T, requirements []sa.Requirement) {
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
func SeedPayments(db *sql.DB, t *testing.T, payments []sa.Payment) {
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
func SeedMajors(db *sql.DB, t *testing.T, majors []sa.Major) {
	qInsert := sq.Insert("major").Columns("id", "name", "created_at")

	for _, major := range majors {
		qInsert = qInsert.Values(major.ID, major.Name, major.CreatedAt)
	}

	query, args, err := qInsert.PlaceholderFormat(sq.Dollar).ToSql()
	require.NoError(t, err)

	_, err = db.Exec(query, args...)
	require.NoError(t, err)
}
