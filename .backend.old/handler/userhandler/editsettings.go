package userhandler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func EditSettings(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var settings data.UserSettings
	json.NewDecoder(r.Body).Decode(&settings)
	if settings.IsEmpty() {
		return handler.ErrEmptyBody
	}

	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	err = stores.Users.Update(db.Query{"username": user.Username}, data.User{Settings: settings})
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, settings)
	return nil
}
