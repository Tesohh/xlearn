package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Tesohh/xlearn/db"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	ErrInvalidPassword = errors.New("password invalid")
)

func UserLogin(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body loginBody
	json.NewDecoder(r.Body).Decode(&body)

	user, err := stores.Users.One(db.Query{"username": body.Username})
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
	if err != nil {
		return ErrInvalidPassword
	}

	expiration := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      expiration.Unix(),
	})

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Path:    "/",
		Expires: expiration,
	}

	http.SetCookie(w, &cookie)
	writeJSON(w, 200, M{"success": "set cookie properly"})
	return nil
}
