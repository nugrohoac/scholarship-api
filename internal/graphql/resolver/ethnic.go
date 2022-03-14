package resolver

import "github.com/Nusantara-Muda/scholarship-api/src/business/entity"

// EthnicResolver .
type EthnicResolver struct {
	Ethnic entity.Ethnic
}

// ID .
func (e EthnicResolver) ID() *int32 {
	return &e.Ethnic.ID
}

// Name .
func (e EthnicResolver) Name() *string {
	return &e.Ethnic.Name
}
