package common

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"
)

type keyContext string

func (k keyContext) String() string {
	return string(k)
}

const keyUser = keyContext("user")

// SetUserOnContext ...
func SetUserOnContext(ctx context.Context, user entity.User) context.Context {
	return context.WithValue(ctx, keyUser, user)
}

// GetUserOnContext .
func GetUserOnContext(ctx context.Context) (entity.User, error) {
	user, ok := ctx.Value(keyUser).(entity.User)
	if !ok {
		return entity.User{}, errors.ErrBadRequest{Message: "context doesn't contain user"}
	}

	return user, nil
}
