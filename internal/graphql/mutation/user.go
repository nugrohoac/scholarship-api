package mutation

import (
	"context"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/resolver"
	"time"
)

// UserMutation ...
type UserMutation struct {
	userService sa.UserService
}

// RegisterUser ...
func (u UserMutation) RegisterUser(ctx context.Context, param sa.InputRegisterUser) (*resolver.UserResolver, error) {
	user := sa.User{
		Type:     param.Type,
		Email:    param.Email,
		PhoneNo:  param.PhoneNo,
		Password: param.Password,
	}

	user, err := u.userService.Store(ctx, user)
	if err != nil {
		return nil, err
	}

	return &resolver.UserResolver{User: user}, err
}

// UpdateUser ...
func (u UserMutation) UpdateUser(ctx context.Context, param sa.InputUpdateUser) (*resolver.UserResolver, error) {
	cardIdentities := make([]sa.CardIdentity, 0)
	for _, c := range param.CardIdentities {
		cardIdentities = append(cardIdentities, sa.CardIdentity{
			Type: c.Type,
			No:   c.No,
			Image: sa.Image{
				URL:    c.Image.URL,
				Width:  c.Image.Width,
				Height: c.Image.Height,
			},
			UserID: int64(param.ID),
		})
	}

	user := sa.User{
		ID:              int64(param.ID),
		Name:            param.Name,
		CompanyName:     *param.CompanyName,
		CountryID:       param.CountryID,
		PostalCode:      param.PostalCode,
		Address:         param.Address,
		CardIdentities:  cardIdentities,
		BankID:          param.BankID,
		BankAccountNo:   param.BankAccountNo,
		BankAccountName: param.BankAccountName,
	}

	if param.Photo != nil {
		user.Photo = sa.Image{
			URL:    param.Photo.URL,
			Width:  param.Photo.Width,
			Height: param.Photo.Height,
		}
	}

	userUpdated, err := u.userService.UpdateByID(ctx, user.ID, user)
	if err != nil {
		return nil, err
	}

	return &resolver.UserResolver{User: userUpdated}, nil
}

// ActivateUser ...
func (u UserMutation) ActivateUser(ctx context.Context, param struct{ Token string }) (*resolver.UserResolver, error) {
	user, err := u.userService.ActivateStatus(ctx, param.Token)
	if err != nil {
		return nil, err
	}

	return &resolver.UserResolver{User: user}, nil
}

// ResetPassword ...
func (u UserMutation) ResetPassword(ctx context.Context, param struct{ Password string }) (*resolver.UserResolver, error) {
	user, err := u.userService.ResetPassword(ctx, param.Password)
	if err != nil {
		return nil, err
	}

	return &resolver.UserResolver{User: user}, nil
}

// SetupEducation ...
func (u UserMutation) SetupEducation(ctx context.Context, param sa.InputSetupEducation) (*resolver.UserResolver, error) {
	user := sa.User{
		ID:               int64(param.UserID),
		CareerGoal:       param.CareerGoal,
		StudyCountryGoal: sa.Country{ID: param.StudyCountryGoal.ID},
		StudyDestination: param.StudyDestination,
	}

	if param.GapYearReason != nil {
		user.GapYearReason = *param.GapYearReason
	}

	for _, us := range param.UserSchools {
		var userSchool sa.UserSchool

		userSchool.School.ID = int64(us.School.ID)

		if us.Degree != nil {
			userSchool.Degree.ID = us.Degree.ID
		}

		if us.Major != nil {
			userSchool.Major.ID = int64(us.Major.ID)
		}

		if us.EnrollmentDate != nil {
			enrollmentDate, err := time.Parse(time.RFC3339, *us.EnrollmentDate)
			if err != nil {
				return nil, err
			}

			userSchool.EnrollmentDate = enrollmentDate
		}

		graduationDate, err := time.Parse(time.RFC3339, us.GraduationDate)
		if err != nil {
			return nil, err
		}

		userSchool.GraduationDate = graduationDate

		if us.Gpa != nil {
			userSchool.Gpa = *us.Gpa
		}

		user.UserSchools = append(user.UserSchools, userSchool)
	}

	for _, ud := range param.UserDocuments {
		userDoc := sa.UserDocument{
			Document: sa.Image{
				URL:    ud.URL,
				Width:  ud.Width,
				Height: ud.Height,
			},
		}

		if ud.Mime != nil {
			userDoc.Document.Mime = *ud.Mime
		}

		if ud.Caption != nil {
			userDoc.Document.Caption = *ud.Caption
		}

		user.UserDocuments = append(user.UserDocuments, userDoc)
	}

	user, err := u.userService.SetupEducation(ctx, user)
	if err != nil {
		return nil, err
	}

	return &resolver.UserResolver{User: user}, nil
}

// NewUserMutation ...
func NewUserMutation(userService sa.UserService) UserMutation {
	return UserMutation{userService: userService}
}
