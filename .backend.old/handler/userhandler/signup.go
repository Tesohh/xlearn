package userhandler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"golang.org/x/crypto/bcrypt"
)

type signupBody struct {
	Username string `json:"username"`
	Display  string `json:"display,omitempty"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body signupBody
	json.NewDecoder(r.Body).Decode(&body)

	// validate request
	if (body == signupBody{}) {
		return handler.ErrEmptyBody
	} else if body.Username == "" {
		return handler.ErrMalformedBody
	} else if body.Password == "" {
		return handler.ErrMalformedBody
	} else if len(body.Password) < 12 {
		return handler.ErrPWTooShort
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
		return handler.ErrUsernameTaken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return err
	}

	user := data.User{
		Display:      body.Display,
		Username:     body.Username,
		Role:         data.RoleUser,
		Pin:          data.PinString(),
		PasswordHash: string(hash),
		Level:        1,
		Coins:        5,
		Settings: data.UserSettings{
			Language: "en_us",
			Theme:    "workshop",
		},
	}

	// check if user collection is empty, then in that case the user should be admin
	if allUsers, _ := stores.Users.Many(db.Query{}); len(allUsers) == 0 {
		user.Role = data.RoleAdmin
	}

	err = stores.Users.Put(user)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, user)
	return nil
}
