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

func (u userSuite) TestUserRepoStore() {
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

func (u userSuite) TestUserRepoFetch() {
	t := u.T()
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)
	postgresql.SeedUsers(u.DBConn, t, users)

	userRepo := postgresql.NewUserRepository(u.DBConn)
	usersResp, cursor, err := userRepo.Fetch(context.Background(), sa.UserFilter{Email: users[0].Email})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMS0xMVQxNzozMzo1OC40MDNa", cursor)
	require.Equal(t, 1, len(usersResp))
	require.Equal(t, users[0].Name, usersResp[0].Name)
}

func (u userSuite) TestUserRepoLogin() {
	t := u.T()
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)
	postgresql.SeedUsers(u.DBConn, t, users)

	userRepo := postgresql.NewUserRepository(u.DBConn)
	user, err := userRepo.Login(context.Background(), users[0].Email)
	require.NoError(t, err)
	require.Equal(t, users[0].Name, user.Name)
	require.Equal(t, users[0].Type, user.Type)
	require.Equal(t, users[0].Email, user.Email)
}

func (u userSuite) TestUserRepoUpdateByID() {
	t := u.T()

	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)
	user := users[0]

	for i := range users {
		users[i].Name = ""
		users[i].Photo = sa.Image{}
		users[i].CompanyName = ""
		users[i].CountryID = 0
		users[i].PostalCode = ""
		users[i].Address = ""
		users[i].BankID = 0
		users[i].BankAccountNo = ""
		users[i].BankAccountName = ""
	}

	postgresql.SeedUsers(u.DBConn, t, users)

	cardIdentities := make([]sa.CardIdentity, 0)
	testdata.GoldenJSONUnmarshal(t, "card_identities", &cardIdentities)

	user.CardIdentities = cardIdentities
	user.Status = 2

	userRepo := postgresql.NewUserRepository(u.DBConn)
	userResp, err := userRepo.UpdateByID(context.Background(), user.ID, user)

	require.NoError(t, err)
	require.NotEqual(t, sa.User{}, userResp)

	var count int
	row := u.DBConn.QueryRow("SELECT COUNT(id) FROM card_identity WHERE user_id = $1", user.ID)
	err = row.Scan(&count)
	require.NoError(t, err)
	require.Equal(t, 2, count)

	var status int
	row = u.DBConn.QueryRow("SELECT status FROM \"user\" WHERE id = $1", users[0].ID)
	err = row.Scan(&status)
	require.NoError(t, err)
	require.Equal(t, 2, status)
}

func (u userSuite) TestUserSetStatus() {
	t := u.T()
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	postgresql.SeedUsers(u.DBConn, t, users)

	userRepo := postgresql.NewUserRepository(u.DBConn)
	err := userRepo.SetStatus(context.Background(), users[0].ID, 1)
	require.NoError(t, err)

	var status int
	row := u.DBConn.QueryRow("SELECT status FROM \"user\" WHERE id = $1", users[0].ID)
	err = row.Scan(&status)
	require.NoError(t, err)
	require.Equal(t, 1, status)
}

func (u userSuite) TestUserUpdatePassword() {
	t := u.T()
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	postgresql.SeedUsers(u.DBConn, t, users)

	newPassword := "new password"
	email := users[0].Email

	userRepo := postgresql.NewUserRepository(u.DBConn)
	err := userRepo.ResetPassword(context.Background(), email, newPassword)
	require.NoError(t, err)

	row := u.DBConn.QueryRow("SELECT password FROM \"user\" WHERE email = $1", email)
	var password string
	err = row.Scan(&password)
	require.NoError(t, err)

	require.Equal(t, newPassword, password)
}

