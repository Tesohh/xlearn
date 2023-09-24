package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT() (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix() // set expiration in 1 hour

	return token.SignedString(secret)
}
