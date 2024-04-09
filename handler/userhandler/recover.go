package userhandler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"golang.org/x/crypto/bcrypt"
)

type recoverBody struct {
	Username    string `json:"username"`
	Pin         string `json:"pin"`
	NewPassword string `json:"new_password"`
}

func Recover(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body recoverBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	user, err := stores.Users.One(db.Query{"username": body.Username})
	if err != nil {
		return err
	}

	if user.RecoverAttempts >= 3 {
		user.RecoverAttempts += 1
		err = stores.Users.Update(db.Query{"username": body.Username}, *user)
		if err != nil {
			return err
		}
		return handler.ErrTooManyAttempts.Context("please ask your admin to unlock your account from recovering")
	}

	if body.Pin != user.Pin {
		user.RecoverAttempts += 1
		err = stores.Users.Update(db.Query{"username": body.Username}, *user)
		if err != nil {
			return err
		}

		return handler.ErrWrongPin
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hash)
	err = stores.Users.Update(db.Query{"username": body.Username}, *user)
	if err != nil {
		return err
	}

	return handler.WriteJSON(w, 200, handler.M{"success": "successfully changed password"})
}
