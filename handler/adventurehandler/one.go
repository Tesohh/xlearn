package adventurehandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func One(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	tag, ok := mux.Vars(r)["advtag"]
	if !ok {
		return handler.ErrPathVar
	}

	adv, err := stores.Adventures.One(db.Query{"tag": tag})
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, adv)
	return nil
}
