package mock

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/Tesohh/xlearn/data"
	"github.com/gorilla/mux"
)

type M map[string]any
type MS map[string]string

// for convenience's sake, since it's only used for tests, will panic instead of returning errors.
func JSON(data map[string]any) io.Reader {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error while encoding data: %s", err.Error())
	}
	return bytes.NewReader(b)
}

type Request struct {
	Method   string
	Target   string
	Body     M
	Username string
	Vars     MS
}

func (r Request) Build() *http.Request {
	var reader io.Reader
	if r.Body == nil {
		reader = nil
	} else {
		reader = JSON(r.Body)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/user/org/joined", reader)
	req = mux.SetURLVars(req, r.Vars)

	if r.Username != "" {
		req.Header.Add("jwt-username", r.Username)
	}

	return req
}

func BuildRequest(method string, target string, body *M, username string, vars map[string]string) *http.Request {
	var reader io.Reader
	if body == nil {
		reader = nil
	} else {
		reader = JSON(*body)
	}

	r := httptest.NewRequest(http.MethodPost, "/api/user/org/joined", reader)

	r = mux.SetURLVars(r, vars)

	if username != "" {
		r.Header.Add("jwt-username", username)
	}

	return r
}

func Unmarshal[T any](w *httptest.ResponseRecorder) T {
	res := w.Result()
	defer res.Body.Close()

	var doc T
	err := json.NewDecoder(res.Body).Decode(&doc)
	if err != nil {
		log.Fatalf("Error while unmarshaling: %s", err.Error())
	}

	return doc
}

func Unpopulate[T data.GetTagger](docs []T) []string {
	tags := []string{}
	for _, v := range docs {
		tags = append(tags, v.GetTag())
	}
	return tags
}
