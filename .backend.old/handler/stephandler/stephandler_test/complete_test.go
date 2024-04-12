package stephandler_test

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler/stephandler"
	"github.com/Tesohh/xlearn/mock"
)

func TestComplete(t *testing.T) {

	stores, err := mock.Stores()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should not error in right conditions", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := mock.Request{
			Method:   http.MethodGet,
			Target:   "/api/step/@xyz/complete",
			Username: "michele",
			Vars:     mock.MS{"steptag": "forkliftstep1-123456"},
		}.Build()
		err = stephandler.Complete(w, r, stores)
		if err != nil {
			t.Fatal(err)
		}

		// will panic in case of Error being returned
		mock.Unmarshal[data.Success](w)
	})

	t.Run("step should be added to user's completed list", func(t *testing.T) {
		user, err := stores.Users.One(db.Query{"username": "michele"})
		if err != nil {
			t.Fatal(err)
		}

		if !slices.Contains(user.CompletedSteps, "forkliftstep1-123456") {
			t.Fail()
		}
	})

	t.Run("should error when the step is already completed", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := mock.Request{
			Method:   http.MethodGet,
			Target:   "/api/step/@xyz/complete",
			Username: "michele",
			Vars:     mock.MS{"steptag": "forkliftstep1-123456"},
		}.Build()
		err = stephandler.Complete(w, r, stores)

		// fails in case no error is returned
		if err == nil {
			t.Fatal("should error, but instead it's allowed to complete the step twice")
		}

	})
	// t.Run("check if language is italian on forkliftstep1-123456", func(t *testing.T) {
	// 	r = mux.SetURLVars(r, map[string]string{"steptag": "forkliftstep1-123456"})
	// 	err = stephandler.One(w, r, stores)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	//
	// 	res := w.Result()
	// 	defer res.Body.Close()
	// 	var data oneResponse
	// 	err = json.NewDecoder(res.Body).Decode(&data)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	//
	// 	if data.LocalContent.Name != "Muletto 1 silandro" {
	// 		t.Errorf("got name=%s, wanted %s", data.LocalContent.Name, "Muletto 1 silandro")
	// 	} else if data.LocalContent.Content != "ciao" {
	// 		t.Errorf("got content=%s, wanted %s", data.LocalContent.Content, "ciao")
	// 	} else if data.LocalContent.Description != "ciao" {
	// 		t.Errorf("got description=%s, wanted %s", data.LocalContent.Description, "ciao")
	// 	}
	// })
}
