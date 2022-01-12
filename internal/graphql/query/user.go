package query

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
)

// UserQuery ...
type UserQuery struct {
	userService sa.UserService
}

// Login ...
func (u UserQuery) Login(ctx context.Context, email, password string) (*string, error) {
	panic("implement me")
}
