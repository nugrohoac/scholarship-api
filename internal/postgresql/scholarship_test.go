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
		scholarship  entity.Scholarship
		requirements = make([]entity.Requirement, 0)
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
	scholarships := make([]entity.Scholarship, 0)
	testdata.GoldenJSONUnmarshal(t, "scholarships", &scholarships)
	postgresql.SeedScholarship(s.DBConn, t, scholarships)

	scholarshipRepository := postgresql.NewScholarshipRepository(s.DBConn, 72)

	// without filter
	response, cursor, err := scholarshipRepository.Fetch(context.Background(), entity.ScholarshipFilter{})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMS0yOVQxMzozMTowMS4yNjNa", cursor)
	require.Equal(t, 2, len(response))
	require.Equal(t, int64(3), response[0].ID)
	require.Equal(t, int64(2), response[1].ID)

	// filter limit 1
	response, cursor, err = scholarshipRepository.Fetch(context.Background(), entity.ScholarshipFilter{Limit: 1})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMS0yOVQxMzozMTowMy4yNjNa", cursor)
	require.Equal(t, 1, len(response))
	require.Equal(t, int64(3), response[0].ID)

	// filter cursor
	response, cursor, err = scholarshipRepository.Fetch(context.Background(), entity.ScholarshipFilter{Cursor: "MjAyMi0wMS0yOVQxMzozMTowMy4yNjNa"})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMS0yOVQxMzozMTowMS4yNjNa", cursor)
	require.Equal(t, 1, len(response))
	require.Equal(t, int64(2), response[0].ID)

	// filter sponsor id
	response, cursor, err = scholarshipRepository.Fetch(context.Background(), entity.ScholarshipFilter{SponsorID: 1})
	require.NoError(t, err)
	require.Equal(t, "MjAyMi0wMS0yOVQxMzozMTowMS4yNjNa", cursor)
	require.Equal(t, 1, len(response))
	require.Equal(t, int64(2), response[0].ID)
}

func (s scholarshipSuite) TestScholarshipRepoGetByID() {
	var (
		scholarship  entity.Scholarship
		requirements = make([]entity.Requirement, 0)
		user         entity.User
	)

	testdata.GoldenJSONUnmarshal(s.T(), "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(s.T(), "requirements", &requirements)
	testdata.GoldenJSONUnmarshal(s.T(), "user", &user)

	user.Type = entity.Sponsor
	user.ID = scholarship.SponsorID

	requirements[0].ScholarshipID = 1
	requirements[1].ScholarshipID = 1

	postgresql.SeedScholarship(s.DBConn, s.T(), []entity.Scholarship{scholarship})
	postgresql.SeedRequirements(s.DBConn, s.T(), requirements)
	postgresql.SeedUsers(s.DBConn, s.T(), []entity.User{user})

	scholarshipRepo := postgresql.NewScholarshipRepository(s.DBConn, 72)
	response, err := scholarshipRepo.GetByID(context.Background(), scholarship.ID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 2, len(response.Requirements))
}

func (s scholarshipSuite) TestScholarshipRepoApplyScholarship() {
	var (
		user        entity.User
		scholarship entity.Scholarship
		t           = s.T()
	)

	recommendationLetter := entity.Image{
		URL:    "https://recommendation-letter",
		Width:  100,
		Height: 100,
	}

	essay := "essay example"

	currentApplicant := scholarship.CurrentApplicant + 1

	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "user", &user)
	postgresql.SeedScholarship(s.DBConn, t, []entity.Scholarship{scholarship})

	scholarshipRepo := postgresql.NewScholarshipRepository(s.DBConn, 72)
	err := scholarshipRepo.Apply(context.Background(), user.ID, scholarship.ID, currentApplicant, essay, recommendationLetter)
	require.NoError(t, err)
}
