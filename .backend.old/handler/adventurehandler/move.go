package adventurehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

type operation string

const (
	operationInsert = "insert"
	operationSwap   = "swap"
)

type moveBody struct {
	Operation operation `json:"operation"`
	Step      string    `json:"step"`
	Target    string    `json:"target"`
}

func MoveStep(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body moveBody
	json.NewDecoder(r.Body).Decode(&body)

	tag, ok := mux.Vars(r)["advtag"]
	if !ok {
		return handler.ErrPathVar
	}

	if (body == moveBody{}) {
		return handler.ErrEmptyBody
	} else if body.Step == "" || body.Target == "" {
		return handler.ErrMalformedBody
	}

	adv, err := stores.Adventures.One(db.Query{"tag": tag})
	if err != nil {
		return err
	}

	stepIndex := slices.Index(adv.Steps, body.Step)
	targetIndex := slices.Index(adv.Steps, body.Target)
	if stepIndex == -1 || targetIndex == -1 {
		return handler.ErrRequestedItemInexistent
	}

	if body.Operation == operationInsert {
		adv.Steps = slices.Delete(adv.Steps, stepIndex, stepIndex+1)
		adv.Steps = slices.Insert(adv.Steps, targetIndex, body.Step)
	} else if body.Operation == operationSwap {
		adv.Steps[stepIndex] = body.Target
		adv.Steps[targetIndex] = body.Step
	} else {
		return fmt.Errorf("%w. Must be either \"insert\" or \"swap\"", handler.ErrInvalidOperation)
	}

	err = stores.Adventures.Update(db.Query{"tag": tag}, *adv)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, adv)
	return nil
}
