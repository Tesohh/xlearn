package generalhandler

import (
	"net/http"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/Tesohh/xlearn/mock"
)

func MockDB(w http.ResponseWriter, r *http.Request, stores db.StoreHolder) error {
	if !IsDBEmpty(stores) {
		return handler.ErrDatabaseNotEmpty
	}
	err := mock.AddDataToStores(stores)
	if err != nil {
		return err
	}

	handler.WriteJSON(w, 200, handler.M{"success": "mocked MAIN database"})
	return nil
}
