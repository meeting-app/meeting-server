package handlers

import (
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ezradiniz/meeting-server/serializers"
	"github.com/labstack/echo"
)

var secret = os.Getenv("JWT_SECRET")

func createToken(user *serializers.UserResponse) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["email"] = user.Email
	claims["username"] = user.Username

	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

func GetUser(c echo.Context) error {
	user := new(serializers.UserResponse)

	if err := c.Bind(user); err != nil {
		return err
	}

	// TODO: Check user authenticate in db

	return nil
}

func CreateUser(c echo.Context) error {
	user := new(serializers.UserResponse)

	if err := c.Bind(user); err != nil {
		return err
	}

	// TODO: Insert user into db

	tokenString, err := createToken(user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
