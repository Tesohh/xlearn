package stephandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func One(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	tag, ok := mux.Vars(r)["steptag"]
	if !ok {
		return handler.ErrPathVar
	}

	step, err := stores.Steps.One(db.Query{"tag": tag})
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, step)
	return nil
}
