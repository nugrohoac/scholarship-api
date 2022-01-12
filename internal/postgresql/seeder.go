package postgresql

import (
	"database/sql"
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
