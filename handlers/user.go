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
