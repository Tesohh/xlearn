package stephandler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func Edit(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	tag, ok := mux.Vars(r)["steptag"]
	if !ok {
		return handler.ErrPathVar
	}

	var step data.Step
	json.NewDecoder(r.Body).Decode(&step)
	if step.IsEmpty() {
		return handler.ErrEmptyBody
	}
	step.Tag = "" // omit the tag so that it can never be changed!

	err := stores.Steps.Update(db.Query{"tag": tag}, step)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, handler.M{"success": "updated step"})
	return nil
}
