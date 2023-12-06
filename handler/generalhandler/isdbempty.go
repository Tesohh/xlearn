package generalhandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func IsDBEmpty(stores db.StoreHolder) bool {
	if docs, err := stores.Adventures.Many(db.Query{}); len(docs) > 0 || err != nil {
		return false
	} else if docs, err := stores.Orgs.Many(db.Query{}); len(docs) > 0 || err != nil {
		return false
	} else if docs, err := stores.Steps.Many(db.Query{}); len(docs) > 0 || err != nil {
		return false
	} else if docs, err := stores.Users.Many(db.Query{}); len(docs) > 0 || err != nil {
		return false
	}
	return true
}

func IsDBEmptyEndpoint(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	resp := map[string]any{"empty": true}

	ok := IsDBEmpty(stores)
	if !ok {
		resp["empty"] = false
	}

	handler.WriteJSON(w, 200, resp)
	return nil
}
