package data

import (
	"reflect"
)

type Org struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// The org's prefix -- example `/api/org/{tag}/steps`
	Tag    string `json:"tag,omitempty" bson:"tag,omitempty"`
	Secret string `json:"-" bson:"secret,omitempty"`
	// an array of adventures tags
	Adventures []string       `json:"adventures,omitempty" bson:"adventures,omitempty"`
	Codes      map[string]int `json:"-" bson:"codes,omitempty"`
}

func (o Org) IsEmpty() bool {
	return reflect.DeepEqual(o, Org{})
}
