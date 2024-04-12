package userhandler_test

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/handler/userhandler"
	"github.com/Tesohh/xlearn/mock"
)

func TestJoinedOrgs(t *testing.T) {
	stores, err := mock.Stores()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("regular users can only see joined and unprotected orgs", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := mock.BuildRequest("POST", "/api/user/org/joined", nil, "zestyman", nil)
		err = userhandler.JoinedOrgs(w, r, stores)
		if err != nil {
			t.Fatal(err)
		}

		res := mock.Unpopulate(mock.Unmarshal[[]data.Org](w))
		if !reflect.DeepEqual(res, []string{"silandro-investors", "merano-holdings"}) {
			t.Fail()
		}
	})

	t.Run("admins are able to see all orgs", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := mock.BuildRequest("POST", "/api/user/org/joined", nil, "michele", nil)
		err = userhandler.JoinedOrgs(w, r, stores)
		if err != nil {
			t.Fatal(err)
		}

		res := mock.Unmarshal[[]data.Org](w)
		if len(res) != len(mock.Orgs) {
			t.Fail()
		}
	})
}
