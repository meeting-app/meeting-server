package router

import (
	handlers "github.com/ezradiniz/meeting-server/handlers"
	"github.com/ezradiniz/meeting-server/middlewares"
	"github.com/labstack/echo"
)

// PostGroup create post group
func PostGroup(g *echo.Group) {
	middlewares.SetAuthenticate(g)
	g.POST("", handlers.CreatePost)
	g.GET("", handlers.FetchAllPosts)
}
