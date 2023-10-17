package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type M map[string]string

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func CurrentUser(r *http.Request, stores db.StoreHolder) (*data.User, error) {
	username := r.Header.Get("jwt-username")
	if username == "" {
		return nil, ErrJwtUsernameInexistent
	}

	return stores.Users.One(db.Query{"username": username})
}

func GetOrgTag(r *http.Request) (string, bool) {
	vars := mux.Vars(r)
	tag, ok := vars["orgtag"]
	return tag, ok
}

func GetOrg(r *http.Request, stores db.StoreHolder) (*data.Org, error) {
	tag, ok := GetOrgTag(r)
	if !ok {
		return nil, ErrPathVar
	}

	return stores.Orgs.One(db.Query{"tag": tag})
}

type APIFunc func(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error

func MW(f APIFunc, stores db.StoreHolder, modifiers ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// authentication
		if !slices.Contains(modifiers, "unprotected") {
			tokenString, err := r.Cookie("token")
			if err != nil {
				WriteJSON(w, http.StatusUnauthorized, M{"error": "token cookie not found"})
				return
			}

			token, err := jwt.Parse(tokenString.Value, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				WriteJSON(w, 400, M{"error": err.Error()})
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			// basic validation
			if !ok || !token.Valid {
				WriteJSON(w, http.StatusUnauthorized, M{"error": err.Error()})
				return
			}
			// validate expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				WriteJSON(w, http.StatusUnauthorized, M{"error": "token is expired"})
				return
			}
			username, ok := claims["username"]
			if !ok {
				WriteJSON(w, http.StatusUnauthorized, M{"error": "username undefined in jwt"})
				return
			}
			r.Header.Add("jwt-username", username.(string))
			fmt.Printf("claims: %+v\n", claims)
		}

		if slices.Contains(modifiers, "admin") {
			user, err := CurrentUser(r, stores)
			if err != nil {
				WriteJSON(w, 400, M{"error": err.Error()})
				return
			}
			if user.Role < data.RoleAdmin {
				WriteJSON(w, http.StatusUnauthorized, M{"error": ErrUnauthorized.Error()})
				return
			}
		}
		if slices.Contains(modifiers, "teacher") {
			user, err := CurrentUser(r, stores)
			if err != nil {
				WriteJSON(w, 400, M{"error": err.Error()})
				return
			}
			if user.Role < data.RoleTeacher {
				WriteJSON(w, http.StatusUnauthorized, M{"error": ErrUnauthorized.Error()})
				return
			}
		}

		// if authentication / authorization failed at this point the function would have exited
		// so from this point on everything is protected
		err := f(w, r, stores)
		if err != nil {
			if apierr, ok := err.(APIError); ok {
				WriteJSON(w, apierr.Status, M{"error": err.Error()})
				return
			}
			WriteJSON(w, 400, M{"error": err.Error()})
			return
		}
	}
}
