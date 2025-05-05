package admin_handlers

import (
	"net/http"
	"portfolio_backend/api/entities"

	"github.com/labstack/echo/v4"
)

func AdminLogin() echo.HandlerFunc {
	return func(c echo.Context) error {

		var admin entities.UserDTO

		if err := c.Bind(&admin); err != nil {
			return err
		}

		if err := c.Validate(&admin); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		return c.String(http.StatusOK, "Logged in")
	}
}

