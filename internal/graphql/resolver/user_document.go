package resolver

import "github.com/Nusantara-Muda/scholarship-api/src/business/entity"

// UserDocumentResolver .
type UserDocumentResolver struct {
	userDocument entity.UserDocument
}

// ID .
func (u UserDocumentResolver) ID() *int32 {
	ID := int32(u.userDocument.ID)
	return &ID
}

// UserID .
func (u UserDocumentResolver) UserID() *int32 {
	userID := int32(u.userDocument.UserID)
	return &userID
}

// Document .
func (u UserDocumentResolver) Document() *ImageResolver {
	return &ImageResolver{Image: u.userDocument.Document}
}
