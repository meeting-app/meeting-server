package router

import (
	handlers "github.com/ezradiniz/meeting-server/handlers"
	"github.com/ezradiniz/meeting-server/middlewares"
	"github.com/labstack/echo"
)

// UserGroup create user group
func UserGroup(g *echo.Group) {
	g.POST("", handlers.CreateUser)
	g.GET("", handlers.FetchCurrentUser, middlewares.GetAuthenticate())
	g.GET("/:username", handlers.FetchUser)
}
