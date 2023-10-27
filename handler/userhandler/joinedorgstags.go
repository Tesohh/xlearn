package userhandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func JoinedOrgsTags(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	orgs := user.JoinedOrgs

	allOrgs, err := stores.Orgs.Many(db.Query{})
	if err != nil {
		return err
	}

	for _, v := range allOrgs {
		if v.IsUnprotected {
			orgs = append(orgs, v.Tag)
		}
	}

	if orgs == nil {
		orgs = make([]string, 0)
	}

	handler.WriteJSON(w, 200, orgs)
	return nil
}
