package userhandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func JoinedOrgs(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	orgs, err := db.Populate(user.JoinedOrgs, stores.Orgs)
	if err != nil {
		return err
	}

	allOrgs, err := stores.Orgs.Many(db.Query{})
	if err != nil {
		return err
	}

	for _, v := range allOrgs {
		if v.IsUnprotected {
			orgs = append(orgs, v)
		}
	}

	handler.WriteJSON(w, 200, orgs)
	return nil
}
