package stephandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

func LastCompleted(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	// get the current user
	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	if len(user.CompletedSteps) == 0 {
		return handler.ErrNoStepsEverCompleted
	}

	// return the last item in the user's completed list
	last := user.CompletedSteps[len(user.CompletedSteps)-1]

	return handler.WriteJSON(w, 200, handler.M{"last": last})
}
