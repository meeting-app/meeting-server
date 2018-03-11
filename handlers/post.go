package handlers

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ezradiniz/meeting-server/models"
	"github.com/ezradiniz/meeting-server/serializers"
	"github.com/labstack/echo"
)

// CreatePost create a post
func CreatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	res := new(serializers.PostResponse)
	if err := c.Bind(res); err != nil {
		return err
	}

	postModel := &models.PostModel{}
	postModel.Text = res.Text
	postModel.UserID = uint(claims["id"].(float64))

	if err := models.AddPost(postModel); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, postModel)
}

// FetchAllPosts fetch all post
func FetchAllPosts(c echo.Context) error {

	res, err := models.FindAllPosts()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
