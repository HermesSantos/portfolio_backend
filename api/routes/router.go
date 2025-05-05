package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()

	InitiateRoutes(e)
	AddValidators(e)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
