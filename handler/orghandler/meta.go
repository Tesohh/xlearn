package orghandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func Meta(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	org, err := handler.GetOrg(r, stores)
	if err != nil {
		return err
	}

	org.Adventures = []string{}

	handler.WriteJSON(w, 200, org)
	return nil
}
