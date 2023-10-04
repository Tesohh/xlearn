package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/data"
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

func OrgAdventureEdit(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	tag, ok := mux.Vars(r)["advtag"]
	if !ok {
		return ErrPathVar
	}

	var adv data.Adventure
	json.NewDecoder(r.Body).Decode(&adv)
	if adv.IsEmpty() {
		return ErrEmptyBody
	}
	adv.Tag = "" // omit the tag so that it can never be changed!

	err := stores.Adventures.Update(db.Query{"tag": tag}, adv)
	if err != nil {
		return err
	}

	writeJSON(w, 200, M{"success": "updated adventure"})
	return nil
}
