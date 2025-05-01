package middlewares

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuthMiddleware() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("hermes")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("123123")) == 1 {
			return true, nil
		}
		return false, nil
	})
}

func AuthWithJWT(){}
