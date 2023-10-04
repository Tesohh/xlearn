package handler

import (
	"encoding/json"
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

	org.Adventures = []string{}

	writeJSON(w, 200, org)
	return nil
}

func OrgEdit(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	tag, ok := getOrgTag(r)
	if !ok {
		return ErrPathVar
	}

	var org data.Org
	json.NewDecoder(r.Body).Decode(&org)
	if org.IsEmpty() {
		return ErrEmptyBody
	}
	org.Tag = "" // omit the tag so that it can never be changed!

	err := stores.Orgs.Update(db.Query{"tag": tag}, org)
	if err != nil {
		return err
	}

	writeJSON(w, 200, M{"success": "updated org"})
	return nil
}
