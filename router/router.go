package router

import (
	"github.com/ezradiniz/meeting-server/middlewares"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	userGroup := e.Group("/users")
	middlewares.SetAuthenticate(userGroup)

	return e
}
