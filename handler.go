package scholarship_api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GraphQLHandler handle handler wrapper between go-graphql relay with echo
func GraphQLHandler(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())

		return nil
	}
}
