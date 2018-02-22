package router

import (
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	api := e.Group("/api/v1")

	auth := api.Group("/auth")
	user := api.Group("/users")

	AuthGroup(auth)
	UserGroup(user)

	return e
}
