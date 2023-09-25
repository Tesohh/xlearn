package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/dgrijalva/jwt-go"
)

type M map[string]string

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func currentUser(r *http.Request, stores db.StoreHolder) (*data.User, error) {
	username := r.Header.Get("jwt-username")
	if username == "" {
		return nil, ErrJwtUsernameInexistent
	}

	return stores.Users.One(db.Query{"username": username})
}

type APIFunc func(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error

func DecorateHTTPFunc(f APIFunc, stores db.StoreHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// authentication
		if strings.Split(r.URL.Path, "/")[2] != "unprotected" {
			tokenString, err := r.Cookie("token")
			if err != nil {
				writeJSON(w, http.StatusUnauthorized, M{"error": "token cookie not found"})
				return
			}

			token, err := jwt.Parse(tokenString.Value, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				writeJSON(w, 400, M{"error": err.Error()})
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			// basic validation
			if !ok || !token.Valid {
				writeJSON(w, http.StatusUnauthorized, M{"error": err.Error()})
				return
			}
			// validate expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				writeJSON(w, http.StatusUnauthorized, M{"error": "token is expired"})
				return
			}
			username, ok := claims["username"]
			if !ok {
				writeJSON(w, http.StatusUnauthorized, M{"error": "username undefined in jwt"})
				return
			}
			r.Header.Add("jwt-username", username.(string))
			fmt.Printf("claims: %+v\n", claims)
		}

		// if authentication failed at this point the function would have exited
		// so from this point on everything is protected

		err := f(w, r, stores)
		if err != nil {
			writeJSON(w, 400, M{"error": err.Error()})
		}
	}
}
