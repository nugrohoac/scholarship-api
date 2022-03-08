package scholarship

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"reflect"
)

type scholarshipService struct {
	scholarshipRepo     sa.ScholarshipRepository
	bankTransferRepo    sa.BankTransferRepository
	paymentRepo         sa.PaymentRepository
	requirementDescRepo sa.RequirementDescriptionRepository
}

// Create ...
// Status of sponsor should 2
func (s scholarshipService) Create(ctx context.Context, scholarship sa.Scholarship) (sa.Scholarship, error) {
	sponsor, err := sa.GetUserOnContext(ctx)
	if err != nil {
		return sa.Scholarship{}, err
	}

	if sponsor.ID != scholarship.SponsorID {
		return sa.Scholarship{}, sa.ErrUnAuthorize{Message: "sponsor id is not match"}
	}

	if sponsor.Status != 2 {
		return sa.Scholarship{}, sa.ErrNotAllowed{Message: "sponsor un complete profile"}
	}

	if scholarship.FundingEnd.Before(scholarship.FundingStart) {
		return sa.Scholarship{}, sa.ErrBadRequest{Message: "scholarship funding end before funding start"}
	}

	scholarship.Sponsor = sponsor

	// default at created, look at readme.md to get more status
	scholarship.Status = 0

	scholarship, err = s.scholarshipRepo.Create(ctx, scholarship)
	if err != nil {
		return sa.Scholarship{}, nil
	}

	scholarship.Payment.BankTransfer, err = s.bankTransferRepo.Get(ctx)
	if err != nil {
		return sa.Scholarship{}, err
	}

	return scholarship, nil
}

// Fetch ...
func (s scholarshipService) Fetch(ctx context.Context, filter sa.ScholarshipFilter) (sa.ScholarshipFeed, error) {
	scholarships, cursor, err := s.scholarshipRepo.Fetch(ctx, filter)
	if err != nil {
		return sa.ScholarshipFeed{}, err
	}

	scholarshipIDs := make([]int64, 0)
	mapScholarshipIndex := map[int64]int{}
	for index, scholarship := range scholarships {
		scholarshipIDs = append(scholarshipIDs, scholarship.ID)

		mapScholarshipIndex[scholarship.ID] = index
	}

	if len(scholarshipIDs) > 0 {
		requirementDesc, err := s.requirementDescRepo.Fetch(ctx, scholarshipIDs)
		if err != nil {
			return sa.ScholarshipFeed{}, err
		}

		for _, scholarship := range scholarships {
			index := mapScholarshipIndex[scholarship.ID]
			scholarships[index].RequirementDescriptions = requirementDesc[scholarship.ID]
		}
	}

	scholarshipFeed := sa.ScholarshipFeed{
		Cursor:       cursor,
		Scholarships: scholarships,
	}

	filter.Cursor = cursor
	filter.Limit = 1
	scholarships, _, err = s.scholarshipRepo.Fetch(ctx, filter)
	if err != nil {
		return sa.ScholarshipFeed{}, err
	}

	if len(scholarships) == 0 {
		scholarshipFeed.Cursor = ""
	}

	return scholarshipFeed, nil
}

// GetByID ...
func (s scholarshipService) GetByID(ctx context.Context, ID int64) (sa.Scholarship, error) {
	scholarship, err := s.scholarshipRepo.GetByID(ctx, ID)
	if err != nil {
		return sa.Scholarship{}, err
	}

	if scholarship.ID == 0 {
		return sa.Scholarship{}, sa.ErrNotFound{Message: "scholarship is not found"}
	}

	requirementDesc, err := s.requirementDescRepo.Fetch(ctx, []int64{ID})
	if err != nil {
		return sa.Scholarship{}, err
	}

	scholarship.RequirementDescriptions = requirementDesc[ID]

	payments, err := s.paymentRepo.Fetch(ctx, []int64{ID})
	if err != nil {
		return sa.Scholarship{}, err
	}

	if len(payments) == 0 {
		payments = append(payments, sa.Payment{})
	}

	scholarship.Payment = payments[0]
	scholarship.Payment.BankTransfer, err = s.bankTransferRepo.Get(ctx)
	if err != nil {
		return sa.Scholarship{}, err
	}

	return scholarship, nil
}

// Apply .
func (s scholarshipService) Apply(ctx context.Context, userID, scholarshipID int64, essay string, recommendationLetter sa.Image) (string, error) {
	userCtx, err := sa.GetUserOnContext(ctx)
	if err != nil {
		return "", err
	}

	if userCtx.ID != userID {
		return "", sa.ErrUnAuthorize{Message: "user is not match"}
	}

	scholarship, err := s.scholarshipRepo.GetByID(ctx, scholarshipID)
	if err != nil {
		return "", err
	}

	// status 3 is open registration
	if scholarship.Status != 3 {
		return "", sa.ErrNotAllowed{Message: "scholarship is not open registration"}
	}

	applicant := scholarship.CurrentApplicant + 1
	if applicant > scholarship.Awardee {
		return "", sa.ErrNotAllowed{Message: "awardee has been maximum"}
	}

	for _, req := range scholarship.Requirements {
		switch req.Name {
		case "essay_request":
			if req.Value == "required" {
				if essay == "" {
					return "", sa.ErrNotAllowed{Message: "essay is required"}
				}
			}
		case "recommendation_letter_request":
			if req.Value == "required" {
				if reflect.DeepEqual(sa.Image{}, recommendationLetter) {
					return "", sa.ErrNotAllowed{Message: "recommendation letter is required"}
				}
			}
		}
	}

	if err = s.scholarshipRepo.Apply(ctx, userID, scholarshipID, applicant, essay, recommendationLetter); err != nil {
		return "", err
	}

	return "success", nil
}

// NewScholarshipService ...
func NewScholarshipService(
	scholarshipRepo sa.ScholarshipRepository,
	bankTransferRepository sa.BankTransferRepository,
	paymentRepo sa.PaymentRepository,
	requirementDescRepo sa.RequirementDescriptionRepository,
) sa.ScholarshipService {
	return scholarshipService{
		scholarshipRepo:     scholarshipRepo,
		bankTransferRepo:    bankTransferRepository,
		paymentRepo:         paymentRepo,
		requirementDescRepo: requirementDescRepo,
	}
}
