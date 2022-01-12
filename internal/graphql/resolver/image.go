package resolver

import sa "github.com/Nusantara-Muda/scholarship-api"

// ImageResolver ...
type ImageResolver struct {
	Image sa.Image
}

// URL ...
func (i ImageResolver) URL() *string {
	return &i.Image.URL
}

// Width ...
func (i ImageResolver) Width() *int32 {
	width := int32(i.Image.Width)
	return &width
}

// Height ...
func (i ImageResolver) Height() *int32 {
	height := int32(i.Image.Height)
	return &height
}

// Mime ...
func (i ImageResolver) Mime() *string {
	return &i.Image.Mime
}

// Caption ...
func (i ImageResolver) Caption() *string {
	return &i.Image.Caption
}
