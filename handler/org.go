package handler

import (
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
)

func Org(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	org, err := getOrg(r, stores)
	if err != nil {
		return err
	}

	writeJSON(w, 200, org)
	return nil
}

func OrgMeta(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	org, err := getOrg(r, stores)
	if err != nil {
		return err
	}

	org.Adventures = []data.Adventure{}

	writeJSON(w, 200, org)
	return nil
}
