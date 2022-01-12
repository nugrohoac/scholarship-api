package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/mutation"
	"github.com/Nusantara-Muda/scholarship-api/user"
	"github.com/Nusantara-Muda/scholarship-api/country"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/Nusantara-Muda/scholarship-api/bank"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/query"
	"github.com/Nusantara-Muda/scholarship-api/internal/postgresql"
)

var (
	// Database
	dsn string

	bankRepo    sa.BankRepository
	countryRepo sa.CountryRepository
	userRepo sa.UserRepository

	bankService    sa.BankService
	countryService sa.CountryService
	userService sa.UserService

	// BankQuery ...
	BankQuery query.BankQuery
	// CountryQuery ...
	CountryQuery query.CountryQuery
	// UserMutation ...
	UserMutation mutation.UserMutation

	// PortApp apps
	PortApp = 7070

	configPath *string
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

	viper.WatchConfig()
}

func initApp() {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("Error init database connection : ", err)
	}

	bankRepo = postgresql.NewBankRepository(db)
	userRepo = postgresql.NewUserRepository(db)
	countryRepo = postgresql.NewCountryRepository(db)

	bankService = bank.NewBankService(bankRepo)
	userService = user.NewUserService(userRepo)
	countryService = country.NewCountryService(countryRepo)

	BankQuery = query.NewBankQuery(bankService)
	UserMutation = mutation.NewUserMutation(userService)
	CountryQuery = query.NewCountryQuery(countryService)
}
