package postgresql_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

type reqDescSuite struct {
	postgresql.TestSuite
}

func TestRequirementDescriptionRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration requirement description")
	}

	suite.Run(t, new(reqDescSuite))
}

func (r reqDescSuite) TestRequirementDescFetch() {
	scholarships := make([]sa.Scholarship, 0)
	testdata.GoldenJSONUnmarshal(r.T(), "scholarships", &scholarships)

	postgresql.SeedScholarship(r.DBConn, r.T(), scholarships)

	reqDescRepo := postgresql.NewRequirementDescriptionRepository(r.DBConn)
	response, err := reqDescRepo.Fetch(context.Background(), []int64{scholarships[0].ID, scholarships[1].ID})
	require.NoError(r.T(), err)

	require.Equal(r.T(), scholarships[0].RequirementDescriptions, response[scholarships[0].ID])
	require.Equal(r.T(), scholarships[1].RequirementDescriptions, response[scholarships[1].ID])
}
