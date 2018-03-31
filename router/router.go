package router

import (
	"github.com/labstack/echo"
)

// New ...
func New() *echo.Echo {
	e := echo.New()

	api := e.Group("/api/v1")

	me := api.Group("/me")
	auth := api.Group("/auth")
	user := api.Group("/users")
	post := api.Group("/posts")

	MeGroup(me)
	AuthGroup(auth)
	UserGroup(user)
	PostGroup(post)

	return e
}
