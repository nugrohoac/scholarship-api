package main

import (
	"fmt"
	"net/http"
	"time"

	_graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Nusantara-Muda/scholarship-api/cmd"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql"
	"github.com/Nusantara-Muda/scholarship-api/internal/graphql/schema"
	"github.com/Nusantara-Muda/scholarship-api/src/business/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${remote_ip} ${time_rfc3339_nano} \"${method} ${path}\" ${status} ${bytes_out} \"${referer}\" \"${user_agent}\"\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
	}))

	opts := make([]_graphql.SchemaOpt, 0)
	opts = append(opts, _graphql.SubscribeResolverTimeout(10*time.Second))

	rootResolver := graphql.NewRootGraphql(
		cmd.BankQuery,
		cmd.CountryQuery,
		cmd.UserQuery,
		cmd.ScholarshipQuery,
		cmd.DegreeQuery,
		cmd.MajorQuery,
		cmd.SchoolQuery,
		cmd.EthnicQuery,
		cmd.UserMutation,
		cmd.ScholarshipMutation,
		cmd.PaymentMutation,
		cmd.SchoolMutation,
		cmd.SponsorQuery,
	)

	graphQlSchema := _graphql.MustParseSchema(schema.String(), rootResolver, opts...)
	//e.Use(cmd.Middleware.Auth)
	e.POST("/scholarship/graphql", handler.GraphQLHandler(&relay.Handler{Schema: graphQlSchema}), cmd.Middleware.Auth)
	e.GET("/scholarship/graphiql", handler.GraphQLHandler(&relay.Handler{Schema: graphQlSchema}), cmd.Middleware.Auth)

	e.File("/scholarship/graphiql", "web/scholarship/graphiql.html")

	e.GET("/scholarship/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cmd.PortApp)))
}
