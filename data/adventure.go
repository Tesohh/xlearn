package data

import "reflect"

type Adventure struct {
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	Tag         string `bson:"tag,omitempty" json:"tag,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Steps       []Step `bson:"steps,omitempty" json:"steps,omitempty"`
}

func (a Adventure) IsEmpty() bool {
	return reflect.DeepEqual(a, Adventure{})
}

type Step struct {
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	Tag         string `bson:"tag,omitempty" json:"tag,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Content     string `bson:"content,omitempty" json:"content,omitempty"`
	XPAward     int    `bson:"xp_award,omitempty" json:"xp_award,omitempty"`
	CoinsAward  int    `bson:"coins_award,omitempty" json:"coins_award,omitempty"`
	EnergyCost  int    `bson:"energy_cost,omitempty" json:"energy_cost,omitempty"`

	Children []Step `bson:"children,omitempty" json:"children,omitempty"`
}
