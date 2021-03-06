package resolver

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// CardIdentityResolver ...
type CardIdentityResolver struct {
	CardIdentity entity.CardIdentity
}

// ID ...
func (c CardIdentityResolver) ID() *int32 {
	ID := int32(c.CardIdentity.ID)
	return &ID
}

// Type ...
func (c CardIdentityResolver) Type() *string {
	return &c.CardIdentity.Type
}

// No ...
func (c CardIdentityResolver) No() *string {
	return &c.CardIdentity.No
}

// Image ...
func (c CardIdentityResolver) Image() *ImageResolver {
	return &ImageResolver{Image: c.CardIdentity.Image}
}

// UserID ...
func (c CardIdentityResolver) UserID() *int32 {
	UserID := int32(c.CardIdentity.UserID)
	return &UserID
}
