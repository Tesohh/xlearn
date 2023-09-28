package handler

import (
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

	// omit steps
	for i := range org.Adventures {
		org.Adventures[i].Steps = []data.Step{}
	}

	writeJSON(w, 200, org.Adventures)
	return nil
}

func OrgAdventureOne(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	org, err := getOrg(r, stores)
	if err != nil {
		return err
	}

	tag, ok := mux.Vars(r)["advtag"]
	if !ok {
		return ErrPathVar
	}

	var adv data.Adventure
	for _, v := range org.Adventures {
		if v.Tag == tag {
			adv = v
		}
	}
	if adv.IsEmpty() {
		return db.ErrDocumentNotFound
	}

	writeJSON(w, 200, adv)
	return nil
}
