package handler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/gorilla/mux"
)

func OrgAdventuresAll(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	org, err := getOrg(r, stores)
	if err != nil {
		return err
	}

	adventures, err := db.Populate(org.Adventures, stores.Adventures)
	if err != nil {
		return err
	}

	writeJSON(w, 200, adventures)
	return nil
}

func OrgAdventureOne(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	tag, ok := mux.Vars(r)["advtag"]
	if !ok {
		return ErrPathVar
	}

	adv, err := stores.Adventures.One(db.Query{"tag": tag})
	if err != nil {
		return err
	}

	writeJSON(w, 200, adv)
	return nil
}
