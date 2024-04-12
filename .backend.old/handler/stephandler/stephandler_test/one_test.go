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

type oneResponse struct {
	Step         data.Step `json:"step"`
	LocalContent struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Content     string `json:"content"`
	} `json:"local_content"`
	Langs []string `json:"langs"`
}

func TestOne(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/step/@forkliftstep1-123456", nil)
	stores, err := mock.Stores()
	if err != nil {
		t.Fatal(err)
	}

	r.Header.Add("jwt-username", "polaroidking123") // polaroid king is italian
	t.Run("check if language is italian on forkliftstep1-123456", func(t *testing.T) {
		r = mux.SetURLVars(r, map[string]string{"steptag": "forkliftstep1-123456"})
		err = stephandler.One(w, r, stores)
		if err != nil {
			t.Fatal(err)
		}

		res := w.Result()
		defer res.Body.Close()
		var data oneResponse
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			t.Fatal(err)
		}

		if data.LocalContent.Name != "Muletto 1 silandro" {
			t.Errorf("got name=%s, wanted %s", data.LocalContent.Name, "Muletto 1 silandro")
		} else if data.LocalContent.Content != "ciao" {
			t.Errorf("got content=%s, wanted %s", data.LocalContent.Content, "ciao")
		} else if data.LocalContent.Description != "ciao" {
			t.Errorf("got description=%s, wanted %s", data.LocalContent.Description, "ciao")
		}
	})
}
