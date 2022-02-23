package school_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/mocks"
	service "github.com/Nusantara-Muda/scholarship-api/school"
	"github.com/Nusantara-Muda/scholarship-api/testdata"
)

var cursor = "cursor"

func TestSchoolServiceCreate(t *testing.T) {
	var (
		school sa.School
		user   sa.User
	)

	testdata.GoldenJSONUnmarshal(t, "school", &school)
	testdata.GoldenJSONUnmarshal(t, "user", &user)

	school.CreatedBy = user.Email
	ctxValid := sa.SetUserOnContext(context.Background(), user)

	tests := map[string]struct {
		paramCtx     context.Context
		paramSchool  sa.School
		createSchool testdata.FuncCaller
		expectedResp sa.School
		expectedErr  error
	}{
		"error get user on context": {
			paramCtx:     context.Background(),
			paramSchool:  school,
			createSchool: testdata.FuncCaller{},
			expectedResp: sa.School{},
			expectedErr:  sa.ErrBadRequest{Message: "context doesn't contain user"},
		},
		"error create school": {
			paramCtx:    ctxValid,
			paramSchool: school,
			createSchool: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctxValid, school},
				Output:   []interface{}{sa.School{}, errors.New("error")},
			},
			expectedResp: sa.School{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramCtx:    ctxValid,
			paramSchool: school,
			createSchool: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{ctxValid, school},
				Output:   []interface{}{school, nil},
			},
			expectedResp: school,
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			schoolRepoMock := new(mocks.SchoolRepository)

			if test.createSchool.IsCalled {
				schoolRepoMock.On("Create", test.createSchool.Input...).
					Return(test.createSchool.Output...).
					Once()
			}

			schoolService := service.NewSchoolService(schoolRepoMock)
			response, err := schoolService.Create(test.paramCtx, test.paramSchool)
			schoolRepoMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, response)
		})
	}
}

func TestSchoolServiceFetch(t *testing.T) {
	var schools []sa.School
	testdata.GoldenJSONUnmarshal(t, "schools", &schools)

	tests := map[string]struct {
		paramFilter  sa.SchoolFilter
		fetchSchool  testdata.FuncCaller
		expectedResp sa.SchoolFeed
		expectedErr  error
	}{
		"error": {
			paramFilter: sa.SchoolFilter{},
			fetchSchool: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.SchoolFilter{}},
				Output:   []interface{}{[]sa.School{}, "", errors.New("error")},
			},
			expectedResp: sa.SchoolFeed{},
			expectedErr:  errors.New("error"),
		},
		"success": {
			paramFilter: sa.SchoolFilter{},
			fetchSchool: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, sa.SchoolFilter{}},
				Output:   []interface{}{schools, cursor, nil},
			},
			expectedResp: sa.SchoolFeed{Cursor: cursor, Schools: schools},
			expectedErr:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			schoolRepoMock := new(mocks.SchoolRepository)

			if test.fetchSchool.IsCalled {
				schoolRepoMock.On("Fetch", test.fetchSchool.Input...).
					Return(test.fetchSchool.Output...).
					Once()
			}

			schoolService := service.NewSchoolService(schoolRepoMock)
			response, err := schoolService.Fetch(context.Background(), test.paramFilter)
			schoolRepoMock.AssertExpectations(t)

			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expectedErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedResp, response)
		})
	}
}
