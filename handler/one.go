package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/db"
)

func OneUser(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var q db.Query
	json.NewDecoder(r.Body).Decode(&q)
	if len(q) == 0 {
		return ErrEmptyBody
	}

	u, err := stores.Users.One(q)
	if err != nil {
		return err
	}

	writeJSON(w, 200, u)
	return nil
}
