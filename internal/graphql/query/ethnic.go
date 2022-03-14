package query

import (
	"context"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
)

// EthnicQuery .
type EthnicQuery struct {
	ethnicService business.EthnicService
}

// FetchEthnic .
func (e EthnicQuery) FetchEthnic(ctx context.Context) (*[]*resolver.EthnicResolver, error) {
	ethnics, err := e.ethnicService.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]*resolver.EthnicResolver, 0)

	for _, ethnic := range ethnics {
		ethnic := ethnic

		response = append(response, &resolver.EthnicResolver{Ethnic: ethnic})
	}

	return &response, nil
}

// NewEthnicQuery .
func NewEthnicQuery(ethnicService business.EthnicService) EthnicQuery {
	return EthnicQuery{ethnicService: ethnicService}
}
