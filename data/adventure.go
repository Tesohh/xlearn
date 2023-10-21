package data

import "reflect"

type Adventure struct {
	Name        string   `bson:"name,omitempty" json:"name"`
	Tag         string   `bson:"tag,omitempty" json:"tag"`
	Description string   `bson:"description,omitempty" json:"description"`
	Steps       []string `bson:"steps,omitempty" json:"steps"`
}

func (a Adventure) IsEmpty() bool {
	return reflect.DeepEqual(a, Adventure{})
}
