package data

import (
	"reflect"
)

type Org struct {
	Name string `json:"name" bson:"name,omitempty"`
	// The org's prefix -- example `/api/org/{tag}/steps`
	Tag           string `json:"tag" bson:"tag,omitempty"`
	IsUnprotected bool   `json:"is_unprotected" bson:"is_unprotected,omitempty"`
	// an array of adventures tags
	Adventures []string       `json:"adventures" bson:"adventures,omitempty"`
	Codes      map[string]int `json:"-" bson:"codes,omitempty"`
}

func (o Org) IsEmpty() bool {
	return reflect.DeepEqual(o, Org{})
}

func (o Org) GetTag() string {
	return o.Tag
}
