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

func createToken(user *models.UserModel) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email

	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

// LoginUser user login
func LoginUser(c echo.Context) error {
	user := new(serializers.UserResponse)

	if err := c.Bind(user); err != nil {
		return err
	}

	userModel, err := models.FindUserLogin(user.Email, user.Password)
	if err != nil {
		return err
	}

	tokenString, err := createToken(&userModel)

	return c.JSON(http.StatusOK, map[string]string{
		"name":     userModel.Name,
		"username": userModel.Username,
		"token":    tokenString,
	})
}
