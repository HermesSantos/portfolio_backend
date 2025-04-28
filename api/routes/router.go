package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Router () {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/get-user-data", func(c echo.Context) error {
		id := c.FormValue("id")
		return c.String(http.StatusOK, "User data for ID: "+id)
	})

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
  }
}
