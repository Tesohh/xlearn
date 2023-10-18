package orghandler

import (
	"net/http"
	"strconv"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func varsToInt(r *http.Request, key string, def int) int {
	str, ok := mux.Vars(r)[key]
	if !ok {
		return def
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		return def
	}

	return value
}

func Code(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	uses := varsToInt(r, "uses", 1)

	org, err := handler.GetOrg(r, stores)
	if err != nil {
		return err
	}

	key := data.HexString()
	if org.Codes == nil {
		org.Codes = make(map[string]int)
	}
	org.Codes[key] = uses

	err = stores.Orgs.Update(db.Query{"tag": org.Tag}, *org)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, handler.M{"key": key, "uses": strconv.Itoa(uses)})
	return nil
}
