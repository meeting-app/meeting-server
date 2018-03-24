package handlers

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ezradiniz/meeting-server/models"
	"github.com/ezradiniz/meeting-server/serializers"
	"github.com/labstack/echo"
)

// CreateUser create a user
func CreateUser(c echo.Context) error {
	user := new(serializers.UserResponse)

	if err := c.Bind(user); err != nil {
		return err
	}

	userModel := &models.UserModel{}
	userModel.Name = user.Name
	userModel.Username = user.Username
	userModel.Email = user.Email
	if err := userModel.SetPassword(user.Password); err != nil {
		return err
	}

	if err := models.AddUser(userModel); err != nil {
		return err
	}

	tokenString, err := createToken(userModel)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"name":     user.Name,
		"username": user.Username,
		"token":    tokenString,
	})
}

// FetchCurrentUser fetch current user from middleware
func FetchCurrentUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	userRes, err := models.FindUserByEmail(email)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"name":     userRes.Name,
		"username": userRes.Username,
		"token":    user.Raw,
	})
}

// FetchUser fetch user by username
func FetchUser(c echo.Context) error {
	username := c.Param("username")

	user, err := models.FindUserByUsername(username)

	if err != nil {
		return err
	}

	// TODO: Change user.Username to name
	return c.JSON(http.StatusOK, map[string]string{
		"name":     user.Username, // user.Name
		"username": user.Username,
	})
}

// FetchAllPostsFromUser fetch all posts from user
func FetchAllPostsFromUser(c echo.Context) error {
	username := c.Param("username")

	posts, err := models.FindAllPostsFromUser(username)
	if err != nil {
		return err
	}

	/*
	 *  res := make([]serializers.PostResponse, len(posts))
	 *
	 *  for i, post := range posts {
	 *    res[i].ID = fmt.Sprintf("%v", post.ID)
	 *    res[i].Text = post.Text
	 *    res[i].CreatedAt = post.CreatedAt
	 *  }
	 */

	return c.JSON(http.StatusOK, posts)
}
