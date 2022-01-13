package jwt_hash_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/jwt_hash"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

func TestJwtHash(t *testing.T) {
	users := make([]sa.User, 0)
	testdata.GoldenJSONUnmarshal(t, "users", &users)

	secretKey := "this is secret key"
	duration := time.Duration(100) * time.Second

	jwtHash := jwt_hash.NewJwtHash([]byte(secretKey), duration)

	t.Run("success", func(t *testing.T) {
		tokenJwt, err := jwtHash.Encode(users[0])
		require.NoError(t, err)
		require.NotEmpty(t, tokenJwt)

		var c sa.Claim

		err = jwtHash.Decode(tokenJwt, &c)
		require.NoError(t, err)

		require.Equal(t, users[0].Name, c.Name)
		require.Equal(t, users[0].Type, c.Type)
		require.Equal(t, users[0].Email, c.Email)
	})

	t.Run("error", func(t *testing.T) {
		var c sa.Claim

		err := jwtHash.Decode("random-string", &c)
		require.Error(t, err)
		require.Equal(t, "token contains an invalid number of segments", err.Error())
	})
}
