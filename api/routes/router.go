package routes

import (
	"crypto/subtle"
	"log"
	"net/http"
	"portfolio_backend/api/handlers/adminHandlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router () {
	e := echo.New()

	admin := e.Group("/admin")
	publicAdmin := e.Group("/admin")

	admin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("hermes")) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte("123123")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	publicAdmin.GET("/login", admin_handlers.AdminLogin())

	admin.GET("/", admin_handlers.AdminDash())

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
