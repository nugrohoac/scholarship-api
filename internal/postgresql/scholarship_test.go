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

type scholarshipSuite struct {
	postgresql.TestSuite
}

func TestScholarshipRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test scholarship")
	}

	suite.Run(t, new(scholarshipSuite))
}

func (s scholarshipSuite) TestScholarshipRepoCreate() {
	var (
		scholarship  sa.Scholarship
		requirements = make([]sa.Requirement, 0)
	)

	testdata.GoldenJSONUnmarshal(s.T(), "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(s.T(), "requirements", &requirements)

	scholarship.Requirements = requirements

	scholarshipRepo := postgresql.NewScholarshipRepository(s.DBConn)
	scholarshipResp, err := scholarshipRepo.Create(context.Background(), scholarship)
	require.NoError(s.T(), err)
	require.Equal(s.T(), scholarship.Name, scholarshipResp.Name)

	row := s.DBConn.QueryRow("SELECT COUNT(id) from scholarship WHERE id=$1", scholarshipResp.ID)
	var count int
	err = row.Scan(&count)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 1, count)
}
