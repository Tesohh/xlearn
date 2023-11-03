package mock

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type M map[string]any

// for convenience's sake, since it's only used for tests, will panic instead of returning errors.
func JSON(data map[string]any) io.Reader {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error while encoding data: %s", err.Error())
	}
	return bytes.NewReader(b)
}
