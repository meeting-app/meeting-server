package handlers

import (
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ezradiniz/meeting-server/models"
	"github.com/ezradiniz/meeting-server/serializers"
	"github.com/labstack/echo"
)

var secret = os.Getenv("JWT_SECRET")

func createToken(user *serializers.UserResponse) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["name"] = user.Name
	claims["username"] = user.Username
	claims["email"] = user.Email

	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

func LoginUser(c echo.Context) error {
	user := new(serializers.UserResponse)

	if err := c.Bind(user); err != nil {
		return err
	}

	userModel, err := models.FindUser(user.Email, user.Password)
	if err != nil {
		return err
	}

	tokenString, _ := createToken(user)

	return c.JSON(http.StatusOK, map[string]string{
		"name":     userModel.Name,
		"username": userModel.Username,
		"token":    tokenString,
	})
}

func CreateUser(c echo.Context) error {
	user := new(serializers.UserResponse)

	if err := c.Bind(user); err != nil {
		return err
	}

	userModel := &models.UserModel{}
	userModel.Name = user.Name
	userModel.Username = user.Username
	userModel.Email = user.Email
	userModel.SetPassword(user.Password)

	if err := models.AddUser(userModel); err != nil {
		return err
	}

	tokenString, err := createToken(user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"name":     user.Name,
		"username": user.Username,
		"token":    tokenString,
	})
}
