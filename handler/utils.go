package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/db"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type APIFunc func(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error

func DecorateHTTPFunc(f APIFunc, stores db.StoreHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r, stores)
		if err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
		}
	}
}
