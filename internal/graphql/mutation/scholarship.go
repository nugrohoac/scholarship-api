package mutation

import (
	"context"
	"time"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// ScholarshipMutation ...
type ScholarshipMutation struct {
	scholarshipService business.ScholarshipService
}

// CreateScholarship ...
func (s ScholarshipMutation) CreateScholarship(ctx context.Context, param entity.InputScholarship) (*resolver.ScholarshipResolver, error) {
	requirements := make([]entity.Requirement, 0)

	scholarship := entity.Scholarship{
		SponsorID:              int64(param.SponsorID),
		Name:                   param.Name,
		Amount:                 int(param.Amount),
		Awardee:                int(param.Awardee),
		EligibilityDescription: param.EligibilityDescription,
		SubsidyDescription:     param.SubsidyDescription,
	}

	if param.Image != nil {
		scholarship.Image = entity.Image{
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
		requirements = append(requirements, entity.Requirement{
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

// ApplyScholarship .
func (s ScholarshipMutation) ApplyScholarship(ctx context.Context, param entity.InputApplyScholarship) (*string, error) {
	var (
		essay                string
		recommendationLetter entity.Image
	)

	if param.Essay != nil {
		essay = *param.Essay
	}

	if param.RecommendationLetter != nil {
		recommendationLetter = entity.Image{
			URL:    param.RecommendationLetter.URL,
			Width:  param.RecommendationLetter.Width,
			Height: param.RecommendationLetter.Height,
		}

		if param.RecommendationLetter.Mime != nil {
			recommendationLetter.Mime = *param.RecommendationLetter.Mime
		}

		if param.RecommendationLetter.Caption != nil {
			recommendationLetter.Caption = *param.RecommendationLetter.Caption
		}
	}

	message, err := s.scholarshipService.Apply(ctx, int64(param.UserID), int64(param.ScholarshipID), essay, recommendationLetter)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (s ScholarshipMutation) ApprovedScholarship(ctx context.Context, param entity.UpdateScholarshipStatus) (*string, error) {
	message, err := s.scholarshipService.ApprovedScholarship(ctx, int64(param.ID))
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// NewScholarshipMutation ....
func NewScholarshipMutation(scholarshipService business.ScholarshipService) ScholarshipMutation {
	return ScholarshipMutation{scholarshipService: scholarshipService}
}
