package routes

import (
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()

	InitiateRoutes(e)
	e.Validator = &RequestValidator{validator: validator.New()}

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
