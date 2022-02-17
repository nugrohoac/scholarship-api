package query_test

import (
	"fmt"
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
		response, err := degreeQuery.GetDegree()
		require.NoError(t, err)
		require.Equal(t, &responsePtr, response)
	})

	t.Run("cek aja", func(t *testing.T) {
		degrees := []string{"D3", "D4", "S1", "S2", "S3"}
		//fmt.Println(&degrees[0])
		//fmt.Println(&degrees[1])
		//fmt.Println(&degrees[2])
		//fmt.Println(&degrees[3])
		//fmt.Println(&degrees[4])

		responsePtr := make([]*string, 0)
		for _, d := range degrees {
			d := d
			responsePtr = append(responsePtr, &d)
		}

		for _, d := range responsePtr {
			fmt.Println(&d, d, *d)
		}

		//a := 100
		//var b, c *int
		//
		//b = &a
		//c = b
		//
		//fmt.Println(c)
		//fmt.Println(&c)
		//fmt.Println(*c)
	})
}
