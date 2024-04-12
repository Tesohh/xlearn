package userhandler

import (
	"errors"
	"net/http"
	"slices"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func LeaveOrg(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	tag, ok := mux.Vars(r)["orgtag"]
	if !ok {
		return handler.ErrPathVar
	}

	// TEMP
	if len(user.JoinedOrgs) == 1 {
		return handler.APIError{Err: errors.New("you temporarily cannot leave your last org"), Status: 500}
	}
	// TEMP

	index := slices.Index(user.JoinedOrgs, tag)
	if index < 0 {
		return handler.ErrRequestedItemInexistent
	}

	user.JoinedOrgs = slices.Delete(user.JoinedOrgs, index, index+1)

	err = stores.Users.Update(db.Query{"username": user.Username}, *user)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, handler.M{"left": tag})
	return nil
}
