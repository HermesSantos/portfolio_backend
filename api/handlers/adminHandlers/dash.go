package admin_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminDash(c echo.Context) error {
	return c.String(http.StatusOK, "Hi, you have access!")
}
