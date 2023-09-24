package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"golang.org/x/crypto/bcrypt"
)

type signupBody struct {
	Username string `json:"username"`
	Display  string `json:"display,omitempty"`
	Password string `json:"password"`
}

var (
	ErrMalformedBody = errors.New("malformed body")
	ErrPWTooShort    = errors.New("password is too short. must be at least 12 characters")
)

func UserSignup(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body signupBody
	json.NewDecoder(r.Body).Decode(&body)

	// validate request
	if (body == signupBody{}) {
		return ErrEmptyBody
	} else if body.Username == "" {
		return ErrMalformedBody
	} else if body.Password == "" {
		return ErrMalformedBody
	} else if len(body.Password) < 12 {
		return ErrPWTooShort
	}

	// normalize username and display (just in case)
	body.Username = strings.ReplaceAll(body.Username, " ", "-")
	if body.Display == "" {
		body.Display = body.Username
		body.Display = strings.ReplaceAll(body.Display, "-", " ")
		body.Display = strings.ReplaceAll(body.Display, "_", " ")
		body.Display = strings.Title(body.Display)
	}

	// check if username already exists
	_, err := stores.Users.One(db.Query{"username": body.Username})
	if err != db.ErrDocumentNotFound {
		return ErrUsernameTaken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return err
	}

	user := data.User{
		Display:      body.Display,
		Username:     body.Username,
		PasswordHash: string(hash),
		Coins:        5,
	}
	err = stores.Users.Put(user)
	if err != nil {
		return err
	}

	writeJSON(w, 200, user)
	return nil
}
