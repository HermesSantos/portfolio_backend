package routes

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	RequestValidator struct {
		validator *validator.Validate
	}
)

func (cv *RequestValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    // Optionally, you could return the error to give each route more control over the status code
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return nil
}

func AddValidators(e *echo.Echo) {
  e.Validator = &RequestValidator{validator: validator.New()}
}
