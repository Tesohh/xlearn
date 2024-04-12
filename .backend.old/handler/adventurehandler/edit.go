package adventurehandler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func Edit(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	tag, ok := mux.Vars(r)["advtag"]
	if !ok {
		return handler.ErrPathVar
	}

	var adv data.Adventure
	json.NewDecoder(r.Body).Decode(&adv)
	if adv.IsEmpty() {
		return handler.ErrEmptyBody
	}
	adv.Tag = "" // omit the tag so that it can never be changed!

	err := stores.Adventures.Update(db.Query{"tag": tag}, adv)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, handler.M{"success": "updated adventure"})
	return nil
}
