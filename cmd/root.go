package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Nusantara-Muda/scholarship-api/internal/email"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"time"

	"github.com/spf13/viper"

	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/bank"
	"github.com/Nusantara-Muda/scholarship-api/country"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
	"github.com/Nusantara-Muda/scholarship-api/jwt_hash"
	_middleware "github.com/Nusantara-Muda/scholarship-api/middleware"
	"github.com/Nusantara-Muda/scholarship-api/user"
)

var (
	// Database
	dsn string

	bankRepo    sa.BankRepository
	countryRepo sa.CountryRepository
	userRepo    sa.UserRepository
	emailRepo   sa.EmailRepository

	bankService    sa.BankService
	countryService sa.CountryService
	userService    sa.UserService

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

	bankService = bank.NewBankService(bankRepo)
	userService = user.NewUserService(userRepo, jwtHash, emailRepo)
	countryService = country.NewCountryService(countryRepo)

	UserMutation = mutation.NewUserMutation(userService)

	BankQuery = query.NewBankQuery(bankService)
	CountryQuery = query.NewCountryQuery(countryService)
	UserQuery = query.NewUserQuery(userService)
}
