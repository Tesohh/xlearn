package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
)

type orgNewBody struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func OrgNew(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body orgNewBody
	json.NewDecoder(r.Body).Decode(&body)

	org := data.Org{
		Name:   body.Name,
		Tag:    data.Tagify(body.Name, false),
		Secret: body.Secret,
	}
	// validate request
	if (body == orgNewBody{}) {
		return ErrEmptyBody
	} else if body.Name == "" || body.Secret == "" {
		return ErrMalformedBody
	}

	// TODO: check if tag already exists

	err := stores.Orgs.Put(org)
	if err != nil {
		return err
	}

	writeJSON(w, 200, org)
	return nil
}
