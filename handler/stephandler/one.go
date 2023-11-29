package stephandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func One(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	tag, ok := mux.Vars(r)["steptag"]
	if !ok {
		return handler.ErrPathVar
	}

	user, err := handler.CurrentUser(r, stores)
	if err != nil {
		return err
	}

	step, err := stores.Steps.One(db.Query{"tag": tag})
	if err != nil {
		return err
	}

	localContent := map[string]string{
		"names":       localize(step.Names, user.Settings.Language),
		"description": localize(step.Descriptions, user.Settings.Language),
		"content":     localize(step.Contents, user.Settings.Language),
	}

	langs := mergeLangs(getLangs(step.Names), getLangs(step.Descriptions), getLangs(step.Contents))

	// omit everything to optimize
	step.Names = map[string]string{}
	step.Descriptions = map[string]string{}
	step.Contents = map[string]string{}

	res := map[string]any{
		"step":          step,
		"local_content": localContent,
		"langs":         langs,
	}

	handler.WriteJSON(w, 200, res)
	return nil
}
