package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	backoffice2 "github.com/Nusantara-Muda/scholarship-api/internal/graphql/query/backoffice"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/domain/backoffice"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/backoffice/sponsor"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/ethnic"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/major"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/school"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"

	"github.com/Nusantara-Muda/scholarship-api/internal/email"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	_middleware "github.com/Nusantara-Muda/scholarship-api/middleware"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/bank"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/country"
	_degree "github.com/Nusantara-Muda/scholarship-api/src/business/usecase/degree"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/jwt_hash"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/payment"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/scholarship"
	"github.com/Nusantara-Muda/scholarship-api/src/business/usecase/user"
)

var (
	// Database
	dsn             string
	deadlinePayment int

	bankRepo            business.BankRepository
	countryRepo         business.CountryRepository
	userRepo            business.UserRepository
	emailRepo           business.EmailRepository
	scholarshipRepo     business.ScholarshipRepository
	bankTransferRepo    business.BankTransferRepository
	paymentRepo         business.PaymentRepository
	degreeRepo          business.DegreeRepository
	requirementDescRepo business.RequirementDescriptionRepository
	majorRepo           business.MajorRepository
	schoolRepo          business.SchoolRepository
	sponsorRepo         business.SponsorRepository
	ethnicRepo          business.EthnicRepository

	bankService        business.BankService
	countryService     business.CountryService
	userService        business.UserService
	scholarshipService business.ScholarshipService
	paymentService     business.PaymentService
	degreeService      business.DegreeService
	majorService       business.MajorService
	schoolService      business.SchoolService
	sponsorService     business.SponsorService
	ethnicService      business.EthnicService

	// email
	emailDomain        string
	emailApiKey        string
	pathActivateUser   string
	pathForgotPassword string
	emailSender        string

	// BankQuery ...
	BankQuery query.BankQuery
	// CountryQuery ...
	CountryQuery query.CountryQuery
	// UserQuery ...
	UserQuery query.UserQuery
	// MajorQuery .
	MajorQuery query.MajorQuery
	// SchoolQuery .
	SchoolQuery query.SchoolQuery
	// EthnicQuery .
	EthnicQuery query.EthnicQuery

	// UserMutation ...
	UserMutation mutation.UserMutation
	// ScholarshipMutation ...
	ScholarshipMutation mutation.ScholarshipMutation
	//PaymentMutation .
	PaymentMutation mutation.PaymentMutation
	// SchoolMutation .
	SchoolMutation mutation.SchoolMutation
	// ScholarshipQuery ...
	ScholarshipQuery query.ScholarshipQuery
	// DegreeQuery ...
	DegreeQuery query.DegreeQuery
	// SponsorQuery ...
	SponsorQuery backoffice2.SponsorQuery

	// PortApp apps
	PortApp = 7070

	configPath *string

	secretKey    string
	tokeDuration time.Duration

	// Middleware ...
	Middleware _middleware.Middleware
)

func init() {
	initEnv()
	initApp()
}

func initEnv() {
	configPath = flag.String("config-path", ".", "config path")

	flag.Parse()

	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(*configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	timezone := "UTC"
	if tz := viper.GetString("timezone"); tz != "" {
		timezone = tz
	}

	dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s TimeZone=%s sslmode=%s search_path=%s",
		viper.GetString("psql_host"),
		viper.GetString("psql_port"),
		viper.GetString("psql_user"),
		viper.GetString("psql_database"),
		viper.GetString("psql_password"),
		timezone,
		viper.GetString("psql_ssl_mode"),
		viper.GetString("psql_schema"),
	)

	if _portApp := viper.GetInt("port_app"); _portApp != 0 {
		PortApp = _portApp
	}

	secretKey = viper.GetString("secret_key")
	if secretKey == "" {
		log.Fatal("Please provide secret key.......!!!")
	}

	duration := viper.GetInt("token_duration")
	if duration == 0 {
		log.Fatal("Please provide token duration.......!!!")
	}
	tokeDuration = time.Duration(duration) * time.Second

	deadlinePayment = viper.GetInt("deadline_payment")
	if deadlinePayment == 0 {
		log.Fatal("Please provide deadline payment.......!!!")
	}

	emailDomain = viper.GetString("email_domain")
	emailApiKey = viper.GetString("email_api_key")
	pathActivateUser = viper.GetString("email_path_activate_user")
	pathForgotPassword = viper.GetString("email_path_forgot_password")
	emailSender = viper.GetString("email_sender")

	viper.WatchConfig()
}

func initApp() {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("Error init database connection : ", err)
	}

	// email
	mg := mailgun.NewMailgun(emailDomain, emailApiKey)
	emailRepo = email.NewEmailRepository(mg, emailSender, pathActivateUser, pathForgotPassword)

	jwtHash := jwt_hash.NewJwtHash([]byte(secretKey), tokeDuration)
	Middleware = _middleware.New(jwtHash)

	bankRepo = postgresql.NewBankRepository(db)
	userRepo = postgresql.NewUserRepository(db)
	countryRepo = postgresql.NewCountryRepository(db)
	scholarshipRepo = postgresql.NewScholarshipRepository(db, deadlinePayment)
	bankTransferRepo = postgresql.NewBankTransferRepository(db)
	paymentRepo = postgresql.NewPaymentRepository(db)
	degreeRepo = postgresql.NewDegreeRepository(db)
	requirementDescRepo = postgresql.NewRequirementDescriptionRepository(db)
	majorRepo = postgresql.NewMajorRepository(db)
	schoolRepo = postgresql.NewSchoolRepository(db)
	sponsorRepo = backoffice.NewSponsorRepository(db)
	ethnicRepo = postgresql.NewEthnicRepository(db)

	bankService = bank.NewBankService(bankRepo)
	userService = user.NewUserService(userRepo, jwtHash, emailRepo)
	countryService = country.NewCountryService(countryRepo)
	scholarshipService = scholarship.NewScholarshipService(scholarshipRepo, bankTransferRepo, paymentRepo, requirementDescRepo)
	paymentService = payment.NewPaymentService(paymentRepo, scholarshipRepo)
	degreeService = _degree.NewDegreeService(degreeRepo)
	majorService = major.NewMajorService(majorRepo)
	schoolService = school.NewSchoolService(schoolRepo)
	sponsorService = sponsor.NewSponsorService(sponsorRepo)
	ethnicService = ethnic.NewEthnicService(ethnicRepo)

	UserMutation = mutation.NewUserMutation(userService)
	ScholarshipMutation = mutation.NewScholarshipMutation(scholarshipService)
	PaymentMutation = mutation.NewPaymentMutation(paymentService)
	SchoolMutation = mutation.NewSchoolMutation(schoolService)

	BankQuery = query.NewBankQuery(bankService)
	CountryQuery = query.NewCountryQuery(countryService)
	UserQuery = query.NewUserQuery(userService)
	ScholarshipQuery = query.NewScholarshipQuery(scholarshipService)
	DegreeQuery = query.NewDegreeQuery(degreeService)
	MajorQuery = query.NewMajorQuery(majorService)
	SchoolQuery = query.NewSchoolQuery(schoolService)
	SponsorQuery = backoffice2.NewSponsorQuery(sponsorService)
	EthnicQuery = query.NewEthnicQuery(ethnicService)
}
