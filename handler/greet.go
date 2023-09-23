package handler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
)

func Greet(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	writeJSON(w, 200, "COmeon manee")
	return nil
}
