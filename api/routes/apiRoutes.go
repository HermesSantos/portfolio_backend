package routes

import (
	"github.com/labstack/echo/v4"
	admin_handlers "portfolio_backend/api/handlers/adminHandlers"
	"portfolio_backend/api/middlewares"
)

func InitiateRoutes (e *echo.Echo) {
	// route groups
	admin := e.Group("/admin")
	publicAdmin := e.Group("/admin")

	// applying middlewares
	admin.Use(middlewares.BasicAuthMiddleware())

	// routes
	publicAdmin.POST("/login", admin_handlers.AdminLogin())
	admin.GET("/", admin_handlers.AdminDash())
}
