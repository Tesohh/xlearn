package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/Tesohh/xlearn/backend/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginBody struct {
	Tag      string `json:"tag"`
	Password string `json:"password"`
}

func UserLogin(c echo.Context) error {
	cc := c.(CustomContext)
	body, err := Decode[loginBody](cc.Request().Body)
	if err != nil {
		return err
	}

	user, err := cc.Stores.Users.One(db.Query{"tag": body.Tag})
	if err != nil {
		return err
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
	if err != nil {
		return err
	}

	expiration := time.Now().Add(24 * time.Hour)

	// give jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"tag": []byte(body.Tag),
		"exp": expiration.Unix(),
	})
	s, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	cc.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   s,
		Path:    "/",
		Expires: expiration,
	})

	return cc.JSON(200, nil)
}
