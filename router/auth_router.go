package router

import (
	handlers "github.com/ezradiniz/meeting-server/handlers/auth"
	"github.com/labstack/echo"
)

func AuthGroup(g *echo.Group) {
	g.GET("", handlers.GetUser)
	g.POST("", handlers.CreateUser)
}
