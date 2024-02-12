package data

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func Tagify(display string, random bool) string {
	tag := display
	tag = strings.ReplaceAll(tag, " ", "-")
	tag = strings.ReplaceAll(tag, "_", "-")
	tag = strings.ToLower(tag)

	if random {
		tag += fmt.Sprintf("-%s", HexString())
	}
	return tag
}

func HexString() string {
	return fmt.Sprintf("%06s", strconv.FormatUint(rand.Uint64(), 16))[:6]
}

type GetTagger interface {
	GetTag() string
}

type Success struct {
	Success string `json:"success"`
}
type Error struct {
	Error string `json:"error"`
}
