package data

import "reflect"

type Adventure struct {
	Name        string   `bson:"name,omitempty" json:"name,omitempty"`
	Tag         string   `bson:"tag,omitempty" json:"tag,omitempty"`
	Description string   `bson:"description,omitempty" json:"description,omitempty"`
	Steps       []string `bson:"steps,omitempty" json:"steps,omitempty"`
}

func (a Adventure) IsEmpty() bool {
	return reflect.DeepEqual(a, Adventure{})
}
