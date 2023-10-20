package orghandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func One(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	org, err := handler.CurrentOrg(r, stores)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, org)
	return nil
}
