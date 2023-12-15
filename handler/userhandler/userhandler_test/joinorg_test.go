package userhandler_test

import (
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler/userhandler"
	"github.com/Tesohh/xlearn/mock"
	"github.com/gorilla/mux"
)

func TestJoinOrg(t *testing.T) {
	stores, err := mock.Stores()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("joins the org `silandro-investors` with the code 123456", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := mock.Request("POST", "/api/user/org/join/123456", nil, "donerkebab", map[string]string{"code": "123456"})
		mux.SetURLVars(r, map[string]string{"code": "123456"})

		err = userhandler.JoinOrg(w, r, stores)
		if err != nil {
			t.Fatal(err)
		}

		user, err := stores.Users.One(db.Query{"username": "donerkebab"})
		if err != nil {
			t.Fatal(err)
		}

		if !slices.Contains(user.JoinedOrgs, "silandro-investors") {
			t.Fatal("user.JoinedOrgs doesnt contain silandro-investors")
		}
	})

	t.Run("code is reduced", func(t *testing.T) {
		org, err := stores.Orgs.One(db.Query{"tag": "silandro-investors"})
		if err != nil {
			t.Fatal(err)
		}

		code, ok := org.Codes["123456"]
		if !ok {
			t.Fatal()
		}

		if code != 19 {
			t.Fatal("code wasnt reduced")
		}
	})
}
