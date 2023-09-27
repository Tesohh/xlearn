package handler

import (
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/gorilla/mux"
)

func OrgAdventuresAll(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	vars := mux.Vars(r)
	tag, ok := vars["tag"]
	if !ok {
		return ErrPathVar
	}

	org, err := stores.Orgs.One(db.Query{"tag": tag})
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
