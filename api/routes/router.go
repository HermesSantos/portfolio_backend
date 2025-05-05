package routes

import (
	"log"
	"net/http"
	admin_handlers "portfolio_backend/api/handlers/adminHandlers"
	"portfolio_backend/api/middlewares"

	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()

	// route groups
	admin := e.Group("/admin")
	publicAdmin := e.Group("/admin")

	// applying middlewares
	admin.Use(middlewares.BasicAuthMiddleware())

	// routes
	publicAdmin.POST("/login", admin_handlers.AdminLogin())
	admin.GET("/", admin_handlers.AdminDash())

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
