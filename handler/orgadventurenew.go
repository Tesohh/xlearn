package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
)

type orgAdventureNewBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func OrgAdventureNew(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body orgAdventureNewBody
	json.NewDecoder(r.Body).Decode(&body)

	adventure := data.Adventure{
		Name:        body.Name,
		Tag:         data.Tagify(body.Name, true),
		Description: body.Description,
	}

	// validate request
	if (body == orgAdventureNewBody{}) {
		return ErrEmptyBody
	} else if body.Name == "" {
		return ErrMalformedBody
	}

	err := stores.Adventures.Put(adventure)
	if err != nil {
		return err
	}

	// update the org's list
	org, err := getOrg(r, stores)
	if err != nil {
		return err
	}
	org.Adventures = append(org.Adventures, adventure.Tag)

	err = stores.Orgs.Update(db.Query{"tag": org.Tag}, *org)
	if err != nil {
		return err
	}

	writeJSON(w, 200, adventure)
	return nil
}
