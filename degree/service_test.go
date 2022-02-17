package degree_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Nusantara-Muda/scholarship-api/degree"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
)

func TestDegreeServiceGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		degrees := []string{"D3", "D4", "S1", "S2", "S3"}
		degreeRepoMock := new(mocks.DegreeRepository)
		degreeRepoMock.On("Get").Return(degrees).Once()

		degreeService := degree.NewDegreeService(degreeRepoMock)
		degreeResponse := degreeService.Get()
		degreeRepoMock.AssertExpectations(t)

		for i, _degree := range degreeResponse {
			require.Equal(t, degrees[i], *_degree)
		}
	})
}
