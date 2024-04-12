package orghandler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func Edit(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	tag, ok := handler.CurrentOrgTag(r)
	if !ok {
		return handler.ErrPathVar
	}

	var org data.Org
	json.NewDecoder(r.Body).Decode(&org)
	if org.IsEmpty() {
		return handler.ErrEmptyBody
	}
	org.Tag = "" // omit the tag so that it can never be changed!

	err := stores.Orgs.Update(db.Query{"tag": tag}, org)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, handler.M{"success": "updated org"})
	return nil
}
