package postgresql

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type ethnicRepo struct {
	db *sql.DB
}

// Fetch .
func (e ethnicRepo) Fetch(ctx context.Context) ([]entity.Ethnic, error) {
	query, args, err := sq.Select("id", "name").From("ethnic").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error(errClose)
		}
	}()

	ethnics := make([]entity.Ethnic, 0)

	for rows.Next() {
		var ethnic entity.Ethnic

		if err = rows.Scan(&ethnic.ID, &ethnic.Name); err != nil {
			return nil, err
		}

		ethnics = append(ethnics, ethnic)
	}

	return ethnics, nil
}

// NewEthnicRepository .
func NewEthnicRepository(db *sql.DB) business.EthnicRepository {
	return ethnicRepo{db: db}
}
