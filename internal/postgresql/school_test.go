package postgresql_test

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
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
	school := sa.School{}
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
