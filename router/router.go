package router

import (
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	auth := e.Group("/auth")

	AuthGroup(auth)

	return e
}
