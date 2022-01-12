package postgresql_test

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type userSuite struct {
	postgresql.TestSuite
}

func TestUserRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test user")
	}

	suite.Run(t, new(userSuite))
}

func (u userSuite) TestUserRepoFetch() {
	t := u.T()
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	userRepo := postgresql.NewUserRepository(u.DBConn)
	_, err := userRepo.Store(context.Background(), users[0])
	require.NoError(t, err)

	var phoneNo string
	row := u.DBConn.QueryRow("select phone_no from \"user\" where email = $1", users[0].Email)
	err = row.Scan(&phoneNo)
	require.NoError(t, err)
	require.Equal(t, users[0].PhoneNo, phoneNo)
}
