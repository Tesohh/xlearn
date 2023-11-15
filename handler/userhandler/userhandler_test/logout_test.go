package userhandler_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tesohh/xlearn/handler/userhandler"
	"github.com/Tesohh/xlearn/mock"
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	stores, err := mock.Stores()
	if err != nil {
		t.Fatal(err)
	}
	r := mock.Request("POST", "/api/user/logout", nil, "", nil)

	err = userhandler.Logout(w, r, stores)
	if err != nil {
		t.Fail()
	}

	header := w.Header().Get("Set-Cookie")
	if header == "" || !strings.HasPrefix(header, "token=;") {
		t.Fail()
	}

}
