package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Tesohh/xlearn/backend/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var (
	ErrUnauthorized        = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	ErrInvalidToken        = echo.NewHTTPError(http.StatusUnauthorized, "invalid jwt token")
	ErrTokenExpired        = echo.NewHTTPError(http.StatusUnauthorized, "jwt token expired")
	ErrTagClaimNotExistant = echo.NewHTTPError(http.StatusUnauthorized, "jwt token doesn't contain tag claim")
)

func AuthMW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(CustomContext)
		tokenString, err := cc.Cookie("token")
		if err != nil {
			return ErrUnauthorized
		}

		token, err := jwt.Parse(tokenString.Value, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return err
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		// basic validation
		if !ok || !token.Valid {
			return ErrInvalidToken
		}
		// validate expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return ErrTokenExpired
		}
		tag, ok := claims["tag"]
		if !ok {
			return ErrTagClaimNotExistant
		}

		user, err := cc.Stores.Users.One(db.Query{"tag": tag})
		if err != nil {
			return err
		}

		cc.User = user

		return next(cc)
	}
}
