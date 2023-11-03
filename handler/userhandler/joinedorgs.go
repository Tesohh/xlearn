package userhandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func JoinedOrgs(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	var orgs []data.Org

	if user.Role < data.RoleAdmin { // not admin
		orgs, err = db.Populate(user.JoinedOrgs, stores.Orgs)
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
	} else { // admin
		orgs, err = stores.Orgs.Many(db.Query{})
		if err != nil {
			return err
		}

	}

	handler.WriteJSON(w, 200, orgs)
	return nil
}
