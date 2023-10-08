package orghandler

import (
	"encoding/json"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
)

type orgNewBody struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func New(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	var body orgNewBody
	json.NewDecoder(r.Body).Decode(&body)

	tag := data.Tagify(body.Name, false)

	org := data.Org{
		Name:   body.Name,
		Tag:    tag,
		Secret: body.Secret,
	}
	// validate request
	if (body == orgNewBody{}) {
		return handler.ErrEmptyBody
	} else if body.Name == "" || body.Secret == "" {
		return handler.ErrMalformedBody
	}

	if _, err := stores.Orgs.One(db.Query{"tag": tag}); err == nil {
		return handler.ErrTagTaken
	}

	err := stores.Orgs.Put(org)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, org)
	return nil
}
