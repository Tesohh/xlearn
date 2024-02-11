package stephandler

import (
	"net/http"
	"slices"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func Complete(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	tag, ok := mux.Vars(r)["steptag"]
	if !ok {
		return handler.ErrPathVar
	}

	if slices.Contains(user.CompletedSteps, tag) {
		return handler.ErrAlreadyCompletedStep
	}

	step, err := stores.Steps.One(db.Query{"tag": tag})
	if err != nil {
		return err
	}

	user.CompletedSteps = append(user.CompletedSteps, tag)
	user.XP += step.XPAward
	user.Coins += step.CoinsAward

	err = stores.Users.Update(db.Query{"username": user.Username}, *user)
	if err != nil {
		return err
	}

	return handler.WriteJSON(w, 200, handler.M{"success": "successfully completed step"})
}
