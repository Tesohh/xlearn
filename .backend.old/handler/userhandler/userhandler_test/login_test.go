package userhandler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tesohh/xlearn/handler/userhandler"
	"github.com/Tesohh/xlearn/mock"
)

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	stores, err := mock.Stores()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("login with wrong username errors", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/api/user/login", mock.JSON(mock.M{
			"username": "asdasdasjkdasdhjkasdhjkasdhjkasdhjkasdjkh",
			"password": "na no way boy",
		}))
		err = userhandler.Login(w, r, stores)
		if err == nil {
			t.Fail()
		}
	})

	t.Run("incorrect password errors", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/api/user/login", mock.JSON(mock.M{
			"username": "michele",
			"password": "clearlywrongpassword",
		}))
		err = userhandler.Login(w, r, stores)
		if err == nil {
			t.Fail()
		}
	})

	t.Run("correct login sets cookie", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/api/user/login", mock.JSON(mock.M{
			"username": "michele",
			"password": "michelepazzofolle",
		}))
		err = userhandler.Login(w, r, stores)
		if err != nil {
			t.Fail()
		}
		header := w.Header().Get("Set-Cookie")
		if header == "" || !strings.HasPrefix(header, "token=") {
			t.Fail()
		}
	})
}
