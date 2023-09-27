package handler

import (
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
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
