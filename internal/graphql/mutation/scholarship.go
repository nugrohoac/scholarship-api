package mutation

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
)

// ScholarshipMutation ...
type ScholarshipMutation struct {
	scholarshipService sa.ScholarshipService
}

// CreateScholarship ...
func (s ScholarshipMutation) CreateScholarship(ctx context.Context, param sa.InputScholarship) (*resolver.ScholarshipResolver, error) {
	requirements := make([]sa.Requirement, 0)

	scholarship := sa.Scholarship{
		SponsorID: int64(param.SponsorID),
		Name:      param.Name,
		Amount:    param.Amount,
		Image: sa.Image{
			URL:    param.Image.URL,
			Width:  param.Image.Width,
			Height: param.Image.Height,
		},
		Awardee:                param.Awardee,
		Deadline:               param.Deadline,
		EligibilityDescription: param.EligibilityDescription,
		SubsidyDescription:     param.SubsidyDescription,
		FundingStart:           param.FundingStart,
		FundingEnd:             param.FundingEnd,
	}

	for _, req := range param.Requirements {
		requirements = append(requirements, sa.Requirement{
			Type:  req.Type,
			Name:  req.Name,
			Value: req.Value,
		})
	}

	scholarship.Requirements = requirements

	scholarship, err := s.scholarshipService.Create(ctx, scholarship)
	if err != nil {
		return nil, err
	}

	return &resolver.ScholarshipResolver{Scholarship: scholarship}, nil
}

// NewScholarshipMutation ....
func NewScholarshipMutation(scholarshipService sa.ScholarshipService) ScholarshipMutation {
	return ScholarshipMutation{scholarshipService: scholarshipService}
}
