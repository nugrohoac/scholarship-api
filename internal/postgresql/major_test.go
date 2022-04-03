package postgresql_test

import (
	"context"
	"fmt"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type majorSuite struct {
	postgresql.TestSuite
}

func TestMajorRepo(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test major")
	}

	suite.Run(t, new(majorSuite))
}

func (m majorSuite) TestMajorRepoFetch() {
	t := m.T()
	majors := make([]entity.Major, 0)
	testdata.GoldenJSONUnmarshal(t, "majors", &majors)

	postgresql.SeedMajors(m.DBConn, t, majors)

	majorRepo := postgresql.NewMajorRepository(m.DBConn)
	// without filter
	response, cursor, err := majorRepo.Fetch(context.Background(), entity.MajorFilter{})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMi0xOVQxNToyODo0NS40NTla", cursor)
	require.Len(t, response, 3)

	// limit 1
	response, cursor, err = majorRepo.Fetch(context.Background(), entity.MajorFilter{Limit: 1})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMi0xOVQxNToyODo1NS40NTla", cursor)
	require.Len(t, response, 1)

	// cursor
	response, cursor, err = majorRepo.Fetch(context.Background(), entity.MajorFilter{Limit: 2, Cursor: cursor})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMi0xOVQxNToyODo0NS40NTla", cursor)
	require.Len(t, response, 2)

	// name
	response, cursor, err = majorRepo.Fetch(context.Background(), entity.MajorFilter{Name: "theMAti"})
	require.NoError(t, err)
	fmt.Println(cursor)
	require.Equal(t, "MjAyMi0wMi0xOVQxNToyODo0NS40NTla", cursor)
	require.Len(t, response, 1)
}
