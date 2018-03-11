package middlewares

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var secret = os.Getenv("JWT_SECRET")

// GetAuthenticate get authenticate middleware
func GetAuthenticate() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(secret),
	})
}

// SetAuthenticate set authenticate middleware
func SetAuthenticate(g *echo.Group) {
	g.Use(GetAuthenticate())
}
