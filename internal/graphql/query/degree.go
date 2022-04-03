package query

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
)

// DegreeQuery ...
type DegreeQuery struct {
	degreeService business.DegreeService
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
func NewDegreeQuery(degreeService business.DegreeService) DegreeQuery {
	return DegreeQuery{degreeService: degreeService}
}
