package router

import (
	"github.com/labstack/echo"
)

// New ...
func New() *echo.Echo {
	e := echo.New()

	api := e.Group("/api/v1")

	auth := api.Group("/auth")
	user := api.Group("/users")
	post := api.Group("/posts")

	AuthGroup(auth)
	UserGroup(user)
	PostGroup(post)

	return e
}
