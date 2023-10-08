package userhandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func Me(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	u, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, u)
	return nil
}
