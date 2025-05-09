package routes

import (
	admin_handlers "portfolio_backend/api/handlers/adminHandlers"
	"portfolio_backend/api/middlewares"

	"github.com/labstack/echo/v4"
)

func InitiateRoutes(e *echo.Echo) {
	// route groups
	admin := e.Group("/admin")
	publicAdmin := e.Group("/public")

	// applying middlewares
	admin.Use(middlewares.BasicAuthMiddleware())

	// routes
	publicAdmin.POST("/login", admin_handlers.AdminLogin)
	admin.GET("/", admin_handlers.AdminDash)
}
