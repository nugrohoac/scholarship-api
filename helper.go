package scholarship_api

import (
	"context"
)

type keyContext string

func (k keyContext) String() string {
	return string(k)
}

const keyUser = keyContext("user")

// SetUserOnContext ...
func SetUserOnContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, keyUser, user)
}

// GetUserOnContext .
func GetUserOnContext(ctx context.Context) (User, error) {
	user, ok := ctx.Value(keyUser).(User)
	if !ok {
		return User{}, ErrBadRequest{Message: "context doesn't contain user"}
	}

	return user, nil
}
