package mutation_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestScholarshipMutationCreate(t *testing.T) {
	var (
		scholarship  entity.Scholarship
		requirements = make([]entity.Requirement, 0)
		payment      entity.Payment
	)

	testdata.GoldenJSONUnmarshal(t, "scholarship", &scholarship)
	testdata.GoldenJSONUnmarshal(t, "requirements", &requirements)
	testdata.GoldenJSONUnmarshal(t, "payment", &payment)

	inputScholarship := entity.InputScholarship{
		SponsorID: int32(scholarship.SponsorID),
		Name:      scholarship.Name,
		Amount:    int32(scholarship.Amount),
		Image: &entity.InputImage{
			URL:    scholarship.Image.URL,
			Width:  scholarship.Image.Width,
			Height: scholarship.Image.Height,
		},
		Awardee:                 int32(scholarship.Awardee),
		ApplicationStart:        scholarship.ApplicationStart.Format(time.RFC3339Nano),
		ApplicationEnd:          scholarship.ApplicationEnd.Format(time.RFC3339Nano),
		AnnouncementDate:        scholarship.AnnouncementDate.Format(time.RFC3339Nano),
		EligibilityDescription:  scholarship.EligibilityDescription,
		SubsidyDescription:      scholarship.SubsidyDescription,
		RequirementDescriptions: scholarship.RequirementDescriptions,
		FundingStart:            scholarship.FundingStart.Format(time.RFC3339Nano),
		FundingEnd:              scholarship.FundingEnd.Format(time.RFC3339Nano),
	}

	for i, req := range requirements {
		inputScholarship.Requirements = append(inputScholarship.Requirements, entity.InputRequirement{
			Type:  req.Type,
			Name:  req.Name,
			Value: req.Value,
		})

		requirements[i].ScholarshipID = 0
		requirements[i].ID = 0
		requirements[i].CreatedAt = time.Time{}
	}

	scholarship.Requirements = requirements
	scholarship.ID = 0
	scholarship.CreatedAt = time.Time{}
	scholarship.UpdatedAt = time.Time{}

	scholarshipResp := scholarship
	scholarshipResp.Payment.ID = payment.ID
	scholarshipResp.Payment.Deadline = payment.Deadline
	scholarshipResp.Payment.ScholarshipID = scholarship.ID
	scholarshipResolver := resolver.ScholarshipResolver{Scholarship: scholarshipResp}

	tests := map[string]struct {
		paramScholarship  entity.InputScholarship
		createScholarship testdata.FuncCaller
		expectedResp      *resolver.ScholarshipResolver
		expectedErr       error
	}{
		"success": {
			paramScholarship: inputScholarship,
			createScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarship},
				Output:   []interface{}{scholarshipResp, nil},
			},
			expectedResp: &scholarshipResolver,
			expectedErr:  nil,
		},
		"error": {
			paramScholarship: inputScholarship,
			createScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, scholarship},
				Output:   []interface{}{entity.Scholarship{}, errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			scholarshipServiceMock := new(mocks.ScholarshipService)

			if test.createScholarship.IsCalled {
				scholarshipServiceMock.On("Create", test.createScholarship.Input...).
					Return(test.createScholarship.Output...).
					Once()
			}

			scholarshipMutation := mutation.NewScholarshipMutation(scholarshipServiceMock)
			resp, err := scholarshipMutation.CreateScholarship(context.Background(), test.paramScholarship)
			scholarshipServiceMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, resp)
		})
	}
}

func TestApprovedScholarship(t *testing.T) {
	inputScholarship := entity.UpdateScholarshipStatus{
		ID:         1,
		ActionType: 1,
	}

	msgSuccess := "success"
	tests := map[string]struct {
		paramScholarship    entity.UpdateScholarshipStatus
		approvedScholarship testdata.FuncCaller
		expectedResp        *string
		expectedErr         error
	}{
		"success": {
			paramScholarship: inputScholarship,
			approvedScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, int64(1), int32(1)},
				Output:   []interface{}{"success", nil},
			},
			expectedResp: &msgSuccess,
			expectedErr:  nil,
		},
		"error": {
			paramScholarship: inputScholarship,
			approvedScholarship: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, int64(1), int32(1)},
				Output:   []interface{}{"", errors.New("error")},
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			scholarshipServiceMock := new(mocks.ScholarshipService)

			if test.approvedScholarship.IsCalled {
				scholarshipServiceMock.On("ApprovedScholarship", test.approvedScholarship.Input...).
					Return(test.approvedScholarship.Output...).
					Once()
			}

			scholarshipMutation := mutation.NewScholarshipMutation(scholarshipServiceMock)
			resp, err := scholarshipMutation.ApprovedScholarship(context.Background(), test.paramScholarship)
			scholarshipServiceMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, resp)
		})
	}
}
