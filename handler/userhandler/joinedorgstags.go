package userhandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func JoinedOrgsTags(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	orgs := user.JoinedOrgs

	if user.Role < data.RoleAdmin {
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
	} else {
		allOrgs, err := stores.Orgs.Many(db.Query{})
		if err != nil {
			return nil
		}

		for _, v := range allOrgs {
			orgs = append(orgs, v.Tag)
		}
	}

	handler.WriteJSON(w, 200, orgs)
	return nil
}
