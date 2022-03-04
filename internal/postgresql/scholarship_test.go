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

	scholarshipRepo := postgresql.NewScholarshipRepository(s.DBConn, 72)
	scholarshipResp, err := scholarshipRepo.Create(context.Background(), scholarship)
	require.NoError(s.T(), err)
	require.Equal(s.T(), scholarship.Name, scholarshipResp.Name)

	row := s.DBConn.QueryRow("SELECT COUNT(id) from scholarship WHERE id=$1", scholarshipResp.ID)
	var count int
	err = row.Scan(&count)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 1, count)

	row = s.DBConn.QueryRow("SELECT COUNT(id) from payment WHERE scholarship_id=$1", scholarshipResp.ID)
	err = row.Scan(&count)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 1, count)
}

func (s scholarshipSuite) TestScholarshipRepoFetch() {
	t := s.T()
	scholarships := make([]sa.Scholarship, 0)
	testdata.GoldenJSONUnmarshal(t, "scholarships", &scholarships)
	postgresql.SeedScholarship(s.DBConn, t, scholarships)

	scholarshipRepository := postgresql.NewScholarshipRepository(s.DBConn, 72)

	// without filter
	response, cursor, err := scholarshipRepository.Fetch(context.Background(), sa.ScholarshipFilter{})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMS0yOVQxMzozMTowMS4yNjNa", cursor)
	require.Equal(t, 2, len(response))
	require.Equal(t, int64(3), response[0].ID)
	require.Equal(t, int64(2), response[1].ID)

	// filter limit 1
	response, cursor, err = scholarshipRepository.Fetch(context.Background(), sa.ScholarshipFilter{Limit: 1})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMS0yOVQxMzozMTowMy4yNjNa", cursor)
	require.Equal(t, 1, len(response))
	require.Equal(t, int64(3), response[0].ID)

	// filter cursor
	response, cursor, err = scholarshipRepository.Fetch(context.Background(), sa.ScholarshipFilter{Cursor: "MjAyMi0wMS0yOVQxMzozMTowMy4yNjNa"})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMS0yOVQxMzozMTowMS4yNjNa", cursor)
	require.Equal(t, 1, len(response))
	require.Equal(t, int64(2), response[0].ID)

	// filter sponsor id
	response, cursor, err = scholarshipRepository.Fetch(context.Background(), sa.ScholarshipFilter{SponsorID: 1})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMS0yOVQxMzozMTowMS4yNjNa", cursor)
	require.Equal(t, 1, len(response))
	require.Equal(t, int64(2), response[0].ID)
}

func (s scholarshipSuite) TestScholarshipRepoGetByID() {
	var (
		scholarship  sa.Scholarship
		requirements = make([]sa.Requirement, 0)
	)

	testdata.GoldenJSONUnmarshal(s.T(), "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(s.T(), "requirements", &requirements)

	requirements[0].ScholarshipID = 1
	requirements[1].ScholarshipID = 1

	postgresql.SeedScholarship(s.DBConn, s.T(), []sa.Scholarship{scholarship})
	postgresql.SeedRequirements(s.DBConn, s.T(), requirements)

	scholarshipRepo := postgresql.NewScholarshipRepository(s.DBConn, 72)
	response, err := scholarshipRepo.GetByID(context.Background(), scholarship.ID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 2, len(response.Requirements))
}

func (s scholarshipSuite) TestScholarshipRepoApplyScholarship() {
	var (
		user        sa.User
		scholarship sa.Scholarship
		t           = s.T()
	)

	documents := []sa.Document{
		{
			Name: "essay_request",
			Value: sa.Image{
				URL:    "https://docs1",
				Width:  100,
				Height: 100,
			},
		},
		{
			Name: "recommendation_letter_request",
			Value: sa.Image{
				URL:    "https://docs2",
				Width:  100,
				Height: 100,
			},
		},
	}

	currentApplicant := scholarship.CurrentApplicant + 1

	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "user", &user)
	postgresql.SeedScholarship(s.DBConn, t, []sa.Scholarship{scholarship})

	scholarshipRepo := postgresql.NewScholarshipRepository(s.DBConn, 72)
	err := scholarshipRepo.Apply(context.Background(), user.ID, scholarship.ID, currentApplicant, documents)
	require.NoError(t, err)
}
