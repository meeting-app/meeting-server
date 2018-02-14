package middlewares

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAuthenticate(g *echo.Group) {
	secret := os.Getenv("JWT_SECRET")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(secret),
	}))
}
