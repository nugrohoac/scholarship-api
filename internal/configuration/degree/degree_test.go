package degree_test

import (
	"testing"

	"github.com/Nusantara-Muda/scholarship-api/internal/configuration/degree"
	"github.com/stretchr/testify/require"
)

func TestDegreeGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		degrees := []string{"D3", "D4", "S1", "S2", "S3"}

		degreeRepo := degree.NewDegreeRepository(degrees)
		degreeResponse := degreeRepo.Get()
		require.Equal(t, degrees, degreeResponse)
	})
}
