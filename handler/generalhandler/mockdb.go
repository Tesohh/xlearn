package generalhandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/Tesohh/xlearn/mock"
)

func MockDB(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	if docs, _ := stores.Adventures.Many(db.Query{}); len(docs) > 0 {
		return handler.ErrDatabaseNotEmpty
	} else if docs, _ := stores.Orgs.Many(db.Query{}); len(docs) > 0 {
		return handler.ErrDatabaseNotEmpty
	} else if docs, _ := stores.Steps.Many(db.Query{}); len(docs) > 0 {
		return handler.ErrDatabaseNotEmpty
	} else if docs, _ := stores.Users.Many(db.Query{}); len(docs) > 0 {
		return handler.ErrDatabaseNotEmpty
	}

	err := mock.AddDataToStores(stores)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, handler.M{"success": "mocked MAIN database"})
	return nil
}
