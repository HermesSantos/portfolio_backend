package admin_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Logged in")
	}
}
