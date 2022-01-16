package middleware

import (
	"github.com/graph-gophers/graphql-go/errors"
	"net/http"
	"strings"

	_graphql "github.com/graph-gophers/graphql-go"
	"github.com/labstack/echo/v4"

	sa "github.com/Nusantara-Muda/scholarship-api"
)

// Middleware ...
type Middleware struct {
	jwtHash sa.JwtHash
}

// Auth .
func (m Middleware) Auth(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		token := c.Request().Header.Get(echo.HeaderAuthorization)
		tokens := strings.Split(token, " ")
		if tokens[0] != "Bearer" {
			return handler(c)
		}

		token = tokens[1]

		var claim sa.Claim
		err := m.jwtHash.Decode(token, &claim)
		if err != nil {
			res := _graphql.Response{
				Errors: []*errors.QueryError{
					{
						Message: err.Error(),
					},
				},
			}

			return c.JSON(http.StatusOK, res)
		}

		ctx = sa.SetUserOnContext(ctx, sa.User{
			Name:   claim.Name,
			Email:  claim.Email,
			Type:   claim.Type,
			Status: claim.Status,
		})

		c.SetRequest(c.Request().WithContext(ctx))

		return handler(c)
	}
}

// New ...
func New(jtwHash sa.JwtHash) Middleware {
	return Middleware{jwtHash: jtwHash}
}
