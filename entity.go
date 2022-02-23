package scholarship_api

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	// Sponsor ...
	Sponsor = "sponsor"
	// KTP // Student ...
	//Student = "student"
	// KTP ...
	KTP = "ktp"
	// NPWP ...
	NPWP = "npwp"
)

// Bank ...
type Bank struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

// BankTransfer  is used to hold information transfer of payment
type BankTransfer struct {
	Name        string
	AccountName string
	AccountNo   int
	Image       Image
}

// InputBankFilter ...
type InputBankFilter struct {
	Limit  *int32
	Cursor *string
	Name   *string
}

// BankFilter ...
type BankFilter struct {
	Limit  int
	Cursor string
	Name   string
}

// BankFeed ...
type BankFeed struct {
	Cursor string `json:"cursor"`
	Banks  []Bank `json:"banks"`
}

// Image ...
type Image struct {
	URL     string `json:"url"`
	Width   int32  `json:"width"`
	Height  int32  `json:"height"`
	Mime    string `json:"mime"`
	Caption string `json:"caption"`
}

// InputImage ...
type InputImage struct {
	URL     string
	Width   int32
	Height  int32
	Mime    *string
	Caption *string
}

// User ....
type User struct {
	ID              int64          `json:"id"`
	Name            string         `json:"name"`
	Type            string         `json:"type"`
	Email           string         `json:"email"`
	PhoneNo         string         `json:"phone_no"`
	Photo           Image          `json:"photo"`
	CompanyName     string         `json:"company_name"`
	Password        string         `json:"-"`
	Status          int            `json:"status"`
	CountryID       int32          `json:"country_id"`
	PostalCode      string         `json:"postal_code"`
	Address         string         `json:"address"`
	Gender          string         `json:"gender"`
	Ethnic          string         `json:"ethnic"`
	CardIdentities  []CardIdentity `json:"card_identities"`
	BankID          int32          `json:"bank_id"`
	BankAccountNo   string         `json:"bank_account_no"`
	BankAccountName string         `json:"bank_account_name"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"-"`
}

// UserFilter ...
type UserFilter struct {
	Email string
	IDs   []int64
}

// InputRegisterUser .
type InputRegisterUser struct {
	Type     string
	Email    string
	PhoneNo  string
	Password string
}

// Country ...
type Country struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// CountryFeed ...
type CountryFeed struct {
	Cursor    string
	Countries []Country `json:"countries"`
}

// CountryFilter ...
type CountryFilter struct {
	Limit  int
	Cursor string
	Name   string
}

// InputCountryFilter ...
type InputCountryFilter struct {
	Limit  *int32
	Cursor *string
	Name   *string
}

// Claim ...
type Claim struct {
	ID     int64
	Name   string
	Email  string
	Type   string
	Status int
	jwt.StandardClaims
}

// InputLogin ...
type InputLogin struct {
	Email    string
	Password string
}

// LoginResponse ...
type LoginResponse struct {
	Token string
	User  User
}

// CardIdentity ...
type CardIdentity struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	No        string    `json:"no"`
	Image     Image     `json:"image"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// InputCardIdentity ...
type InputCardIdentity struct {
	Type  string
	No    string
	Image InputImage
}

// InputUpdateUser ...
type InputUpdateUser struct {
	ID              int32
	Name            string
	Photo           InputImage
	CompanyName     *string
	CountryID       int32
	Address         string
	PostalCode      string
	CardIdentities  []InputCardIdentity
	BankID          int32
	BankAccountNo   string
	BankAccountName string
}

// Scholarship ...
type Scholarship struct {
	ID                      int64         `json:"id"`
	SponsorID               int64         `json:"sponsor_id"`
	Sponsor                 User          `json:"sponsor"`
	Name                    string        `json:"name"`
	Amount                  int           `json:"amount"`
	Status                  int           `json:"status"`
	Image                   Image         `json:"image"`
	Awardee                 int           `json:"awardee"`
	CurrentApplicant        int           `json:"current_applicant"`
	ApplicationStart        time.Time     `json:"application_start"`
	ApplicationEnd          time.Time     `json:"application_end"`
	AnnouncementDate        time.Time     `json:"announcement_date"`
	EligibilityDescription  string        `json:"eligibility_description"`
	SubsidyDescription      string        `json:"subsidy_description"`
	RequirementDescriptions []string      `json:"requirement_descriptions"`
	FundingStart            time.Time     `json:"funding_start"`
	FundingEnd              time.Time     `json:"funding_end"`
	Requirements            []Requirement `json:"requirements"`
	Payment                 Payment       `json:"payment"`
	CreatedAt               time.Time     `json:"created_at"`
	UpdatedAt               time.Time     `json:"-"`
}

// ScholarshipFilter ...
type ScholarshipFilter struct {
	Limit     uint64
	Cursor    string
	SponsorID int64
	Status    []int32
	Name      string
}

// InputScholarshipFilter ...
type InputScholarshipFilter struct {
	Limit     *int32
	Cursor    *string
	SponsorID *int32
	Status    *[]*int32
	Name      *string
}

// ScholarshipFeed ...
type ScholarshipFeed struct {
	Cursor       string
	Scholarships []Scholarship
}

// Requirement ...
type Requirement struct {
	ID            int64     `json:"id"`
	ScholarshipID int64     `json:"scholarship_id"`
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	Value         string    `json:"value"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"-"`
}

// InputScholarship ...
type InputScholarship struct {
	SponsorID               int32
	Name                    string
	Amount                  int32
	Image                   *InputImage
	Awardee                 int32
	ApplicationStart        string
	ApplicationEnd          string
	AnnouncementDate        string
	EligibilityDescription  string
	SubsidyDescription      string
	RequirementDescriptions []string
	FundingStart            string
	FundingEnd              string
	Requirements            []InputRequirement
}

// InputRequirement ...
type InputRequirement struct {
	Type  string
	Name  string
	Value string
}

// Payment of scholarship
type Payment struct {
	ID              int64        `json:"id"`
	ScholarshipID   int64        `json:"scholarship_id"`
	BankTransfer    BankTransfer `json:"bank_transfer"`
	Deadline        time.Time    `json:"deadline"`
	TransferDate    time.Time    `json:"transfer_date"`
	BankAccountName string       `json:"bank_account_name"`
	BankAccountNo   string       `json:"bank_account_no"` // string karena di bank permata syariah dimulai dengan 00
	Image           Image        `json:"image"`
	CreatedAt       time.Time    `json:"created_at"`
}

// InputSubmitTransfer ...
type InputSubmitTransfer struct {
	ScholarshipID   int32
	TransferDate    string
	BankAccountName string
	BankAccountNo   string
	Image           InputImage
}

// Major .
type Major struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// MajorFeed .
type MajorFeed struct {
	Cursor string  `json:"cursor"`
	Majors []Major `json:"majors"`
}

// MajorFilter ...
type MajorFilter struct {
	Limit  uint64
	Cursor string
	Name   string
}

// InputMajorFilter .
type InputMajorFilter struct {
	Limit  *int32
	Cursor *string
	Name   *string
}

// School .
type School struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Address   string    `json:"address"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"-"`
}

// SchoolFilter .
type SchoolFilter struct {
	Limit  uint64
	Cursor string
	Name   string
}

// InputSchoolFilter .
type InputSchoolFilter struct {
	Limit  *int32
	Cursor *string
	Name   *string
}

// InputSchool ...
type InputSchool struct {
	Name    string
	Type    string
	Address string
}
