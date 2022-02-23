package query

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// DegreeQuery ...
type DegreeQuery struct {
	degreeService sa.DegreeService
}

// FetchDegree ...
func (d DegreeQuery) FetchDegree(ctx context.Context) (*[]*resolver.DegreeResolver, error) {
	degrees, err := d.degreeService.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	
	response := make([]*resolver.DegreeResolver, 0)

	for _, degree := range degrees {
		degree := degree
		response = append(response, &resolver.DegreeResolver{Degree: degree})
	}

	return &response, nil
}

// NewDegreeQuery ...
func NewDegreeQuery(degreeService sa.DegreeService) DegreeQuery {
	return DegreeQuery{degreeService: degreeService}
}
