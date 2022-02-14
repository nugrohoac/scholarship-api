package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Nusantara-Muda/scholarship-api/payment"

	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/bank"
	"github.com/Nusantara-Muda/scholarship-api/country"
	"github.com/Nusantara-Muda/scholarship-api/internal/bank_transfer"
	"github.com/Nusantara-Muda/scholarship-api/internal/email"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/Nusantara-Muda/scholarship-api/jwt_hash"
	_middleware "github.com/Nusantara-Muda/scholarship-api/middleware"
	"github.com/Nusantara-Muda/scholarship-api/scholarship"
	"github.com/Nusantara-Muda/scholarship-api/user"
)

var (
	// Database
	dsn             string
	deadlinePayment int
	bankTransfer    sa.BankTransfer

	bankRepo         sa.BankRepository
	countryRepo      sa.CountryRepository
	userRepo         sa.UserRepository
	emailRepo        sa.EmailRepository
	scholarshipRepo  sa.ScholarshipRepository
	bankTransferRepo sa.BankTransferRepository
	paymentRepo      sa.PaymentRepository

	bankService        sa.BankService
	countryService     sa.CountryService
	userService        sa.UserService
	scholarshipService sa.ScholarshipService
	paymentService     sa.PaymentService

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

	// UserMutation ...
	UserMutation mutation.UserMutation
	// ScholarshipMutation ...
	ScholarshipMutation mutation.ScholarshipMutation
	//PaymentMutation .
	PaymentMutation mutation.PaymentMutation
	// ScholarshipQuery ...
	ScholarshipQuery query.ScholarshipQuery

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

	// bank transfer load from env
	if bankTransfer.Name = viper.GetString("bank_transfer_name"); bankTransfer.Name == "" {
		log.Fatal("Please provide bank transfer name.......!!!")
	}

	if bankTransfer.AccountName = viper.GetString("bank_transfer_account_name"); bankTransfer.AccountName == "" {
		log.Fatal("Please provide bank transfer account name.......!!!")
	}

	if bankTransfer.AccountNo = viper.GetInt("bank_transfer_account_no"); bankTransfer.AccountNo == 0 {
		log.Fatal("Please provide bank transfer account no.......!!!")
	}

	if bankTransfer.Image.URL = viper.GetString("bank_transfer_image_url"); bankTransfer.Image.URL == "" {
		log.Fatal("Please provide bank transfer image url.......!!!")
	}

	if bankTransfer.Image.Width = viper.GetInt32("bank_transfer_image_width"); bankTransfer.Image.Width == 0 {
		log.Fatal("Please provide bank transfer image width.......!!!")
	}

	if bankTransfer.Image.Height = viper.GetInt32("bank_transfer_image_height"); bankTransfer.Image.Height == 0 {
		log.Fatal("Please provide bank transfer image height.......!!!")
	}

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
	bankTransferRepo = bank_transfer.NewBankTransfer(bankTransfer)
	paymentRepo = postgresql.NewPaymentRepository(db)

	bankService = bank.NewBankService(bankRepo)
	userService = user.NewUserService(userRepo, jwtHash, emailRepo)
	countryService = country.NewCountryService(countryRepo)
	scholarshipService = scholarship.NewScholarshipService(scholarshipRepo, bankTransferRepo, paymentRepo)
	paymentService = payment.NewPaymentService(paymentRepo)

	UserMutation = mutation.NewUserMutation(userService)
	ScholarshipMutation = mutation.NewScholarshipMutation(scholarshipService)
	PaymentMutation = mutation.NewPaymentMutation(paymentService)

	BankQuery = query.NewBankQuery(bankService)
	CountryQuery = query.NewCountryQuery(countryService)
	UserQuery = query.NewUserQuery(userService)
	ScholarshipQuery = query.NewScholarshipQuery(scholarshipService)
}
