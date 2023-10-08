package adventurehandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func All(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	org, err := handler.GetOrg(r, stores)
	if err != nil {
		return err
	}

	adventures, err := db.Populate(org.Adventures, stores.Adventures)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, adventures)
	return nil
}
