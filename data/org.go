package data

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

type Org struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// The org's prefix -- example `/api/org/{tag}/steps`
	Tag        string      `json:"tag,omitempty" bson:"tag,omitempty"`
	Secret     string      `json:"-" bson:"secret,omitempty"`
	Adventures []Adventure `json:"adventures,omitempty" bson:"adventures,omitempty"`
}

func (o Org) IsEmpty() bool {
	return reflect.DeepEqual(o, Org{})
}

func Tagify(display string, random bool) string {
	tag := display
	tag = strings.ReplaceAll(tag, " ", "-")
	tag = strings.ReplaceAll(tag, "_", "-")
	tag = strings.ToLower(tag)

	if random {
		tag += fmt.Sprintf("-%06s", strconv.FormatUint(rand.Uint64(), 16))[:6]
	}
	return tag
}
