package stephandler

import (
	"net/http"
	"slices"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func LastCompletedSpecific(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	// get adventure tag
	advtag, ok := mux.Vars(r)["advtag"]
	if !ok {
		return handler.ErrPathVar.Context("advtag")
	}

	// get the adventure
	adv, err := stores.Adventures.One(db.Query{"tag": advtag})
	if err != nil {
		return err
	}

	// get the current user
	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	// filter out steps that aren't from the adventure
	filtered := []string{}
	for _, s := range user.CompletedSteps {
		if slices.Contains(adv.Steps, s) {
			filtered = append(filtered, s)
		}
	}

	// error if none
	if len(filtered) == 0 {
		return handler.ErrNoStepsCompletedFromAdventure
	}

	// return that step
	last := filtered[len(filtered)-1]
	return handler.WriteJSON(w, 200, handler.M{"last": last})
}
