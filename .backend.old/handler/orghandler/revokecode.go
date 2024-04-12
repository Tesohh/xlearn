package orghandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func RevokeCode(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	code, ok := mux.Vars(r)["code"]
	if !ok {
		return handler.ErrPathVar
	}

	org, err := handler.CurrentOrg(r, stores)
	if err != nil {
		return err
	}

	if _, ok := org.Codes[code]; !ok {
		return handler.ErrRequestedItemInexistent
	}

	delete(org.Codes, code)

	err = stores.Orgs.Update(db.Query{"tag": org.Tag}, *org)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, handler.M{"success": "successfully revoked code"})
	return nil
}
