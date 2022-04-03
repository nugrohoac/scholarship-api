package postgresql_test

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
)

type schoolSuite struct {
	postgresql.TestSuite
}

func TestSchoolRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test school")
	}

	suite.Run(t, new(schoolSuite))
}

func (s schoolSuite) TestSchoolRepo_Create() {
	t := s.T()
	school := entity.School{}
	testdata.GoldenJSONUnmarshal(t, "school", &school)

	schoolRepo := postgresql.NewSchoolRepository(s.DBConn)
	schoolResp, err := schoolRepo.Create(context.Background(), school)
	require.NoError(t, err)
	require.NotEqual(t, 0, schoolResp.ID)
	require.Equal(t, school.Name, schoolResp.Name)
	require.Equal(t, school.Type, schoolResp.Type)
	require.Equal(t, school.Address, schoolResp.Address)
	require.Equal(t, school.Status, schoolResp.Status)
}

func (s schoolSuite) TestSchoolRepo_Fetch() {
	t := s.T()
	var schools []entity.School
	testdata.GoldenJSONUnmarshal(t, "schools", &schools)

	postgresql.SeedSchools(s.DBConn, t, schools)
	schoolRepo := postgresql.NewSchoolRepository(s.DBConn)

	// without filter
	schoolsResp, cursor, err := schoolRepo.Fetch(context.Background(), entity.SchoolFilter{})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMi0yMFQxNzo1OToyNS44MjJa", cursor)
	require.Len(t, schoolsResp, 6)

	// filter limit 2
	schoolsResp, cursor, err = schoolRepo.Fetch(context.Background(), entity.SchoolFilter{Limit: 2})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMi0yMFQxNzo1OTozNy44MjJa", cursor)
	require.Len(t, schoolsResp, 2)

	// limit and cursor
	schoolsResp, cursor, err = schoolRepo.Fetch(context.Background(), entity.SchoolFilter{Limit: 3, Cursor: cursor})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMi0yMFQxNzo1OToyOC44MjJa", cursor)
	require.Len(t, schoolsResp, 3)
	require.Equal(t, "universitas gajah mada", schoolsResp[0].Name)

	// filter name
	schoolsResp, cursor, err = schoolRepo.Fetch(context.Background(), entity.SchoolFilter{Name: "indonesia"})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMi0yMFQxNzo1OTozNy44MjJa", cursor)
	require.Len(t, schoolsResp, 1)
	require.Equal(t, "universitas indonesia", schoolsResp[0].Name)
}
