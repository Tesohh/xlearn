package userhandler

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func JoinOrg(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	code, ok := mux.Vars(r)["code"]
	if !ok {
		return handler.ErrPathVar
	}

	orgs, err := stores.Orgs.Many(db.Query{})
	if err != nil {
		return err
	}

	var org data.Org
	for _, o := range orgs {
		for k := range o.Codes {
			if k == code {
				org = o
				fmt.Println(org)
				break
			}
		}
	}

	if org.IsEmpty() {
		return fmt.Errorf("error while figuring out org: %w", handler.ErrRequestedItemInexistent)
	}

	if slices.Contains(user.JoinedOrgs, org.Tag) {
		return handler.ErrAlreadyJoinedOrg
	}

	fmt.Println(org)
	if _, ok := org.Codes[code]; !ok {
		return handler.ErrRequestedItemInexistent
	}
	org.Codes[code] -= 1
	if org.Codes[code] <= 0 {
		delete(org.Codes, code)
	}
	err = stores.Orgs.Update(db.Query{"tag": org.Tag}, org)
	if err != nil {
		return err
	}

	user.JoinedOrgs = append(user.JoinedOrgs, org.Tag)
	err = stores.Users.Update(db.Query{"username": user.Username}, *user)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, handler.M{"joined": org.Tag})
	return nil
}
