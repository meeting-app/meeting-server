package router

import (
	handlers "github.com/ezradiniz/meeting-server/handlers"
	"github.com/labstack/echo"
)

// AuthGroup create auth group
func AuthGroup(g *echo.Group) {
	g.POST("", handlers.LoginUser)
}
