package postgresql

import (
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"

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
