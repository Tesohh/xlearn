package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
)

var (
	ErrEmptyBody     = errors.New("the request body is empty")
	ErrUsernameTaken = errors.New("username is taken")
)

func PutUser(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var u data.User
	json.NewDecoder(r.Body).Decode(&u)
	if u.IsEmpty() {
		return ErrEmptyBody
	}

	// check if username already exists
	_, err := stores.Users.One(db.Query{"username": u.Username})
	if err != db.ErrDocumentNotFound {
		return ErrUsernameTaken
	}

	if err := stores.Users.Put(u); err != nil {
		return err
	}

	writeJSON(w, 200, u)
	return nil
}
