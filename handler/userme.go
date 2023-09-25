package handler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
)

func UserMe(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	u, err := currentUser(r, stores)
	if err != nil {
		return err
	}

	writeJSON(w, 200, u)
	return nil
}
