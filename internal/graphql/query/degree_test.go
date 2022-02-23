package query_test

import (
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDegreeQuery_GetDegree(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		degrees := []string{"D3", "D4", "S1", "S2", "S3"}
		responsePtr := make([]*string, 0)
		for _, d := range degrees {
			d := d
			responsePtr = append(responsePtr, &d)
		}

		degreeServiceMock := new(mocks.DegreeService)
		degreeServiceMock.On("Get").Return(responsePtr).Once()

		degreeQuery := query.NewDegreeQuery(degreeServiceMock)
		response, err := degreeQuery.FetchDegree()
		require.NoError(t, err)
		require.Equal(t, &responsePtr, response)
	})
}
