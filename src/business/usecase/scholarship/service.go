package scholarship

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/common"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/src/business/errors"
	"reflect"
)

type scholarshipService struct {
	scholarshipRepo     business.ScholarshipRepository
	bankTransferRepo    business.BankTransferRepository
	paymentRepo         business.PaymentRepository
	requirementDescRepo business.RequirementDescriptionRepository
}

// Create ...
// Status of sponsor should 2
func (s scholarshipService) Create(ctx context.Context, scholarship entity.Scholarship) (entity.Scholarship, error) {
	sponsor, err := common.GetUserOnContext(ctx)
	if err != nil {
		return entity.Scholarship{}, err
	}

	if sponsor.ID != scholarship.SponsorID {
		return entity.Scholarship{}, errors.ErrUnAuthorize{Message: "sponsor id is not match"}
	}

	if sponsor.Status != 2 {
		return entity.Scholarship{}, errors.ErrNotAllowed{Message: "sponsor un complete profile"}
	}

	if scholarship.FundingEnd.Before(scholarship.FundingStart) {
		return entity.Scholarship{}, errors.ErrBadRequest{Message: "scholarship funding end before funding start"}
	}

	scholarship.Sponsor = sponsor

	// default at created, look at readme.md to get more status
	scholarship.Status = 0

	scholarship, err = s.scholarshipRepo.Create(ctx, scholarship)
	if err != nil {
		return entity.Scholarship{}, nil
	}

	scholarship.Payment.BankTransfer, err = s.bankTransferRepo.Get(ctx)
	if err != nil {
		return entity.Scholarship{}, err
	}

	return scholarship, nil
}

// Fetch ...
func (s scholarshipService) Fetch(ctx context.Context, filter entity.ScholarshipFilter) (entity.ScholarshipFeed, error) {
	scholarships, cursor, err := s.scholarshipRepo.Fetch(ctx, filter)
	if err != nil {
		return entity.ScholarshipFeed{}, err
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
			return entity.ScholarshipFeed{}, err
		}

		for _, scholarship := range scholarships {
			index := mapScholarshipIndex[scholarship.ID]
			scholarships[index].RequirementDescriptions = requirementDesc[scholarship.ID]
		}
	}

	scholarshipFeed := entity.ScholarshipFeed{
		Cursor:       cursor,
		Scholarships: scholarships,
	}

	filter.Cursor = cursor
	filter.Limit = 1
	scholarships, _, err = s.scholarshipRepo.Fetch(ctx, filter)
	if err != nil {
		return entity.ScholarshipFeed{}, err
	}

	if len(scholarships) == 0 {
		scholarshipFeed.Cursor = ""
	}

	return scholarshipFeed, nil
}

// GetByID ...
func (s scholarshipService) GetByID(ctx context.Context, ID int64) (entity.Scholarship, error) {
	scholarship, err := s.scholarshipRepo.GetByID(ctx, ID)
	if err != nil {
		return entity.Scholarship{}, err
	}

	if scholarship.ID == 0 {
		return entity.Scholarship{}, errors.ErrNotFound{Message: "scholarship is not found"}
	}

	requirementDesc, err := s.requirementDescRepo.Fetch(ctx, []int64{ID})
	if err != nil {
		return entity.Scholarship{}, err
	}

	scholarship.RequirementDescriptions = requirementDesc[ID]

	payments, err := s.paymentRepo.Fetch(ctx, []int64{ID})
	if err != nil {
		return entity.Scholarship{}, err
	}

	if len(payments) == 0 {
		payments = append(payments, entity.Payment{})
	}

	scholarship.Payment = payments[0]
	scholarship.Payment.BankTransfer, err = s.bankTransferRepo.Get(ctx)
	if err != nil {
		return entity.Scholarship{}, err
	}

	return scholarship, nil
}

// Apply .
func (s scholarshipService) Apply(ctx context.Context, userID, scholarshipID int64, essay string, recommendationLetter entity.Image) (string, error) {
	userCtx, err := common.GetUserOnContext(ctx)
	if err != nil {
		return "", err
	}

	if userCtx.ID != userID {
		return "", errors.ErrUnAuthorize{Message: "user is not match"}
	}

	scholarship, err := s.scholarshipRepo.GetByID(ctx, scholarshipID)
	if err != nil {
		return "", err
	}

	// status 3 is open registration
	if scholarship.Status != 3 {
		return "", errors.ErrNotAllowed{Message: "scholarship is not open registration"}
	}

	applicant := scholarship.CurrentApplicant + 1
	if applicant > scholarship.Awardee {
		return "", errors.ErrNotAllowed{Message: "awardee has been maximum"}
	}

	for _, req := range scholarship.Requirements {
		switch req.Name {
		case "essay_request":
			if req.Value == "required" {
				if essay == "" {
					return "", errors.ErrNotAllowed{Message: "essay is required"}
				}
			}
		case "recommendation_letter_request":
			if req.Value == "required" {
				if reflect.DeepEqual(entity.Image{}, recommendationLetter) {
					return "", errors.ErrNotAllowed{Message: "recommendation letter is required"}
				}
			}
		}
	}

	isApply, _, err := s.scholarshipRepo.CheckApply(ctx, userID, scholarshipID)
	if err != nil {
		return "", err
	}

	if isApply {
		return "", errors.ErrNotAllowed{Message: "you have been applied scholarship"}
	}

	if err = s.scholarshipRepo.Apply(ctx, userID, scholarshipID, applicant, essay, recommendationLetter); err != nil {
		return "", err
	}

	return "success", nil
}

// NewScholarshipService ...
func NewScholarshipService(
	scholarshipRepo business.ScholarshipRepository,
	bankTransferRepository business.BankTransferRepository,
	paymentRepo business.PaymentRepository,
	requirementDescRepo business.RequirementDescriptionRepository,
) business.ScholarshipService {
	return scholarshipService{
		scholarshipRepo:     scholarshipRepo,
		bankTransferRepo:    bankTransferRepository,
		paymentRepo:         paymentRepo,
		requirementDescRepo: requirementDescRepo,
	}
}
