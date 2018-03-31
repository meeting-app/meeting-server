package router

import (
	handlers "github.com/ezradiniz/meeting-server/handlers"
	"github.com/ezradiniz/meeting-server/middlewares"
	"github.com/labstack/echo"
)

// MeGroup ...
func MeGroup(g *echo.Group) {
	g.GET("", handlers.FetchCurrentUser, middlewares.GetAuthenticate())
}
