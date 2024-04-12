package userhandler

import (
	"net/http"
	"time"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func Logout(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	cookie := http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	}

	http.SetCookie(w, &cookie)
	handler.WriteJSON(w, 200, handler.M{"success": "logged out properly"})
	return nil
}
