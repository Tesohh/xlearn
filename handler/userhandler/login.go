package userhandler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body loginBody
	json.NewDecoder(r.Body).Decode(&body)

	user, err := stores.Users.One(db.Query{"username": body.Username})
	if err != nil {
		return handler.APIError{Err: err, Status: http.StatusNotFound}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
	if err != nil {
		return handler.ErrInvalidPassword
	}

	expiration := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      expiration.Unix(),
	})

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return handler.APIError{Err: err, Status: http.StatusInternalServerError}
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Path:    "/",
		Expires: expiration,
	}

	http.SetCookie(w, &cookie)
	handler.WriteJSON(w, 200, handler.M{"success": "set cookie properly"})
	return nil
}
