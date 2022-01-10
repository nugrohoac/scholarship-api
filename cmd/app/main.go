package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	var configPath *string
	configPath = flag.String("config-path", ".", "config path")
	fmt.Println(*configPath)
	flag.Parse()

	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(*configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	name := viper.GetString("aku")
	fmt.Println("NAME : ", name)

	//e := echo.New()
	//
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: "${remote_ip} ${time_rfc3339_nano} \"${method} ${path}\" ${status} ${bytes_out} \"${referer}\" \"${user_agent}\"\n",
	//}))
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
	//}))
	//
	//opts := make([]_graphql.SchemaOpt, 0)
	//opts = append(opts, _graphql.SubscribeResolverTimeout(10*time.Second))
	//
	//rootResolver := graphql.NewRootGraphql(cmd.BankQuery)
	//
	////cmd.BankQuery=query.NewBankQuery()
	//graphQlSchema := _graphql.MustParseSchema(schema.String(), &rootResolver, opts...)
	//
	//e.POST("/schedule/graphql", sa.GraphQLHandler(&relay.Handler{Schema: graphQlSchema}))
	//e.GET("/schedule/graphiql", sa.GraphQLHandler(&relay.Handler{Schema: graphQlSchema}))
	//
	//e.File("/schedule/graphiql", "web/scholarship/graphiql.html")
	//
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "OK!")
	//})
	//
	//e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cmd.PortApp)))
}
