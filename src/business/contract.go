package business

import (
	"context"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

// BankRepository .
type BankRepository interface {
	Fetch(ctx context.Context, filter entity.BankFilter) ([]entity.Bank, string, error)
}

// BankService .
type BankService interface {
	Fetch(ctx context.Context, filter entity.BankFilter) (entity.BankFeed, error)
}

// UserRepository ....
type UserRepository interface {
	Store(ctx context.Context, user entity.User) (entity.User, error)
	Fetch(ctx context.Context, filter entity.UserFilter) ([]entity.User, string, error)
	Login(ctx context.Context, email string) (entity.User, error)
	UpdateByID(ctx context.Context, ID int64, user entity.User) (entity.User, error)
	SetStatus(ctx context.Context, ID int64, status int) error
	ResetPassword(ctx context.Context, email, password string) error
	SetupEducation(ctx context.Context, user entity.User) (entity.User, error)
}

// UserService ....
type UserService interface {
	Store(ctx context.Context, user entity.User) (entity.User, error)
	Login(ctx context.Context, email, password string) (entity.LoginResponse, error)
	UpdateByID(ctx context.Context, ID int64, user entity.User) (entity.User, error)
	ActivateStatus(ctx context.Context, token string) (entity.User, error)
	ResendEmailVerification(ctx context.Context, email string) (string, error)
	ResetPassword(ctx context.Context, password string) (entity.User, error)
	ForgotPassword(ctx context.Context, email string) (string, error)
	SetupEducation(ctx context.Context, user entity.User) (entity.User, error)
}

// CountryRepository .
type CountryRepository interface {
	Fetch(ctx context.Context, filter entity.CountryFilter) ([]entity.Country, string, error)
}

// CountryService .
type CountryService interface {
	Fetch(ctx context.Context, filter entity.CountryFilter) (entity.CountryFeed, error)
}

// JwtHash ...
type JwtHash interface {
	Encode(user entity.User) (string, error)
	Decode(tokenString string, claim *entity.Claim) error
}

// EmailRepository ...
type EmailRepository interface {
	SendActivateUser(ctx context.Context, email, token string) error
	SendForgotPassword(ctx context.Context, email, token string) error
}

// ScholarshipRepository ...
type ScholarshipRepository interface {
	Create(ctx context.Context, scholarship entity.Scholarship) (entity.Scholarship, error)
	Fetch(ctx context.Context, filter entity.ScholarshipFilter) ([]entity.Scholarship, string, error)
	GetByID(ctx context.Context, ID int64) (entity.Scholarship, error)
	Apply(ctx context.Context, userID, scholarshipID int64, applicant int, essay string, recommendationLetter entity.Image) error
	CheckApply(ctx context.Context, userID, scholarshipID int64) (bool, int, error)
}

// RequirementDescriptionRepository .
type RequirementDescriptionRepository interface {
	Fetch(ctx context.Context, scholarshipIDs []int64) (map[int64][]string, error)
}

// ScholarshipService ...
type ScholarshipService interface {
	Create(ctx context.Context, scholarship entity.Scholarship) (entity.Scholarship, error)
	Fetch(ctx context.Context, filter entity.ScholarshipFilter) (entity.ScholarshipFeed, error)
	GetByID(ctx context.Context, ID int64) (entity.Scholarship, error)
	Apply(ctx context.Context, userID, scholarshipID int64, essay string, recommendationLetter entity.Image) (string, error)
}

// BankTransferRepository ...
type BankTransferRepository interface {
	Get(ctx context.Context) (entity.BankTransfer, error)
}

// PaymentRepository .
type PaymentRepository interface {
	Fetch(ctx context.Context, scholarshipIDs []int64) ([]entity.Payment, error)
	SubmitTransfer(ctx context.Context, payment entity.Payment) (entity.Payment, error)
}

// PaymentService .
type PaymentService interface {
	SubmitTransfer(ctx context.Context, payment entity.Payment) (entity.Payment, error)
}

// DegreeRepository .
type DegreeRepository interface {
	Fetch(ctx context.Context) ([]entity.Degree, error)
}

// DegreeService .
type DegreeService interface {
	Fetch(ctx context.Context) ([]entity.Degree, error)
}

// MajorRepository ...
type MajorRepository interface {
	Fetch(ctx context.Context, filter entity.MajorFilter) ([]entity.Major, string, error)
}

// MajorService ...
type MajorService interface {
	Fetch(ctx context.Context, filter entity.MajorFilter) (entity.MajorFeed, error)
}

// SchoolRepository .
type SchoolRepository interface {
	Create(ctx context.Context, school entity.School) (entity.School, error)
	Fetch(ctx context.Context, filter entity.SchoolFilter) ([]entity.School, string, error)
}

// SchoolService .
type SchoolService interface {
	Create(ctx context.Context, school entity.School) (entity.School, error)
	Fetch(ctx context.Context, filter entity.SchoolFilter) (entity.SchoolFeed, error)
}

// SponsorRepository .
type SponsorRepository interface {
	FetchSponsor(ctx context.Context, filter entity.SponsorFilter) ([]entity.User, string, error)
}

// SponsorService .
type SponsorService interface {
	FetchSponsor(ctx context.Context, filter entity.SponsorFilter) (entity.SponsorFeed, error)
}

// EthnicRepository .
type EthnicRepository interface {
	Fetch(ctx context.Context) ([]entity.Ethnic, error)
}

// EthnicService .
type EthnicService interface {
	Fetch(ctx context.Context) ([]entity.Ethnic, error)
}

// ApplicantRepository .
type ApplicantRepository interface {
	Fetch(ctx context.Context, filter entity.FilterApplicant) ([]entity.Applicant, string, error)
}

// ApplicantService .
type ApplicantService interface {
	Fetch(ctx context.Context, filter entity.FilterApplicant) (entity.ApplicantFeed, error)
}
