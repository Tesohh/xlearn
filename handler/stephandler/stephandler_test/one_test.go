package stephandler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/handler/stephandler"
	"github.com/Tesohh/xlearn/mock"
	"github.com/gorilla/mux"
)

func TestOne(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/step/@forkliftstep1-123456", nil)
	stores, err := mock.Stores()
	if err != nil {
		t.Fatal(err)
	}

	r.Header.Add("jwt-username", "polaroidking123") // polaroid king is italian

	t.Run("check if `it` content is emptied out on content that has no `it` translation", func(t *testing.T) {
		r = mux.SetURLVars(r, map[string]string{"steptag": "forkliftstep1-123456"})
		err = stephandler.One(w, r, stores)
		if err != nil {
			t.Fatal(err)
		}

		res := w.Result()
		defer res.Body.Close()

		var step data.Step
		err = json.NewDecoder(res.Body).Decode(&step)
		if err != nil {
			t.Fatal(err)
		}

		if step.Content["it"] != "" {
			t.Error("got filled it content, wanted empty")
		}
	})

	w = httptest.NewRecorder()
	t.Run("check if `it` content is filled out on content that has `it` translation", func(t *testing.T) {
		r = mux.SetURLVars(r, map[string]string{"steptag": "forkliftstep1-abcdef"})
		err = stephandler.One(w, r, stores)
		if err != nil {
			t.Fatal(err)
		}

		res := w.Result()
		defer res.Body.Close()

		var step data.Step
		err = json.NewDecoder(res.Body).Decode(&step)
		if err != nil {
			t.Fatal(err)
		}

		if step.Content["it"] == "" {
			t.Error("got empty it content, wanted it filled")
		}
	})

}
