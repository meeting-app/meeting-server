package router

import (
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	auth := e.Group("/auth")
	user := e.Group("/users")

	AuthGroup(auth)
	UserGroup(user)

	return e
}
