package router

import (
	handlers "github.com/ezradiniz/meeting-server/handlers"
	"github.com/labstack/echo"
)

func UserGroup(g *echo.Group) {
	g.POST("", handlers.CreateUser)
}
