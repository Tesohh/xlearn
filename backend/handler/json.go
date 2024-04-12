package handler

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.Reader) (T, error) {
	var target T
	err := json.NewDecoder(body).Decode(&target)
	return target, err
}
