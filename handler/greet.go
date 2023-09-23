package handler

import "net/http"

func Greet(w http.ResponseWriter, r *http.Request) error {
	writeJSON(w, 200, "COmeon manee")
	return nil
}
