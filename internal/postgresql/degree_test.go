package postgresql_test

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type degreeSuite struct {
	postgresql.TestSuite
}

func TestDegreeRepo(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test degree")
	}

	suite.Run(t, new(degreeSuite))
}

func (d degreeSuite) TestDegreeGet() {
	t := d.T()
	var degrees []entity.Degree
	testdata.GoldenJSONUnmarshal(t, "degrees", &degrees)
	postgresql.SeedDegrees(d.DBConn, t, degrees)

	degreeRepo := postgresql.NewDegreeRepository(d.DBConn)
	response, err := degreeRepo.Fetch(context.Background())
	require.NoError(t, err)
	require.Equal(t, len(degrees), len(response))
	require.Equal(t, degrees[0].Name, response[4].Name)
}
