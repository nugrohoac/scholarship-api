package jwt_hash

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// JwtToken ...
type jwtHash struct {
	secretKey []byte
	duration  time.Duration
}

// Encode ...
func (j jwtHash) Encode(user entity.User) (string, error) {
	expireTime := time.Now().Add(j.duration)

	c := entity.Claim{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Type:   user.Type,
		Status: user.Status,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Decode ...
func (j jwtHash) Decode(tokenString string, c *entity.Claim) error {
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return errors.ErrUnAuthorize{Message: err.Error()}
	}

	if !token.Valid {
		return errors.ErrUnAuthorize{Message: "token is invalid"}
	}

	return nil
}

// NewJwtHash ...
func NewJwtHash(secretKey []byte, duration time.Duration) business.JwtHash {
	return jwtHash{
		secretKey: secretKey,
		duration:  duration,
	}
}
