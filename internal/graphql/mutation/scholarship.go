package mutation

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"time"
)

// ScholarshipMutation ...
type ScholarshipMutation struct {
	scholarshipService sa.ScholarshipService
}

// CreateScholarship ...
func (s ScholarshipMutation) CreateScholarship(ctx context.Context, param sa.InputScholarship) (*resolver.ScholarshipResolver, error) {
	requirements := make([]sa.Requirement, 0)

	scholarship := sa.Scholarship{
		SponsorID:              int64(param.SponsorID),
		Name:                   param.Name,
		Amount:                 int(param.Amount),
		Awardee:                int(param.Awardee),
		EligibilityDescription: param.EligibilityDescription,
		SubsidyDescription:     param.SubsidyDescription,
	}

	if param.Image != nil {
		scholarship.Image = sa.Image{
			URL:    param.Image.URL,
			Width:  param.Image.Width,
			Height: param.Image.Height,
		}
	}

	applicationEnd, err := time.Parse(time.RFC3339, param.ApplicationEnd)
	if err != nil {
		return nil, err
	}
	scholarship.ApplicationEnd = applicationEnd

	applicationStart, err := time.Parse(time.RFC3339, param.ApplicationStart)
	if err != nil {
		return nil, err
	}
	scholarship.ApplicationStart = applicationStart

	fundingStart, err := time.Parse(time.RFC3339, param.FundingStart)
	if err != nil {
		return nil, err
	}
	scholarship.FundingStart = fundingStart

	fundingEnd, err := time.Parse(time.RFC3339, param.FundingEnd)
	if err != nil {
		return nil, err
	}
	scholarship.FundingEnd = fundingEnd

	announcementDate, err := time.Parse(time.RFC3339, param.AnnouncementDate)
	if err != nil {
		return nil, err
	}
	scholarship.AnnouncementDate = announcementDate

	for _, req := range param.Requirements {
		requirements = append(requirements, sa.Requirement{
			Type:  req.Type,
			Name:  req.Name,
			Value: req.Value,
		})
	}

	scholarship.RequirementDescriptions = append(scholarship.RequirementDescriptions, param.RequirementDescriptions...)

	scholarship.Requirements = requirements

	scholarship, err = s.scholarshipService.Create(ctx, scholarship)
	if err != nil {
		return nil, err
	}

	return &resolver.ScholarshipResolver{Scholarship: scholarship}, nil
}

// NewScholarshipMutation ....
func NewScholarshipMutation(scholarshipService sa.ScholarshipService) ScholarshipMutation {
	return ScholarshipMutation{scholarshipService: scholarshipService}
}
