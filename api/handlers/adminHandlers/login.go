package admin_handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var (
	JWT_TOKEN_SECRET = os.Getenv("JWT_TOKEN_SECRET")
)

func AdminLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Logged in")
	}
}
func AuthWithJWT() (string, error) {
	secret := os.Getenv(JWT_TOKEN_SECRET)
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "Erro ao gerar token", err
	}
	return tokenString, nil
}
