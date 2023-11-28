package userhandler_test

import (
	"net/http/httptest"
	"testing"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler/userhandler"
	"github.com/Tesohh/xlearn/mock"
)

func TestSignup(t *testing.T) {
	w := httptest.NewRecorder()
	stores, err := mock.Stores()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("short password errors", func(t *testing.T) {
		r := mock.Request("POST", "/api/user/signup", &mock.M{
			"username": "newuser1",
			"password": "short",
		}, "", nil)

		err = userhandler.Signup(w, r, stores)
		if err == nil {
			t.Fail()
		}
	})

	t.Run("user is added to the db", func(t *testing.T) {
		r := mock.Request("POST", "/api/user/signup", &mock.M{
			"username": "newuser1",
			"password": "goodpassword",
		}, "", nil)

		err = userhandler.Signup(w, r, stores)
		if err != nil {
			t.Fail()
		}

		_, err := stores.Users.One(db.Query{"username": "newuser1"})
		if err != nil {
			t.Fail()
		}
	})

}
