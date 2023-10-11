package stephandler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

type manyBody struct {
	Items []string `json:"items"`
}

func Many(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body manyBody
	json.NewDecoder(r.Body).Decode(&body)

	steps := make([]data.Step, 0)
	for _, v := range body.Items {
		step, err := stores.Steps.One(db.Query{"tag": v})
		if err != nil {
			return err
		}

		steps = append(steps, *step)
	}

	handler.WriteJSON(w, 200, steps)
	return nil
}
