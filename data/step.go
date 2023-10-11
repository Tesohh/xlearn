package data

import "reflect"

type StepCat string

const (
	StepCatLesson   StepCat = "lesson"
	StepCatExercise StepCat = "exercise"
	StepCatProject  StepCat = "project"
)

type Step struct {
	Name        string  `bson:"name,omitempty" json:"name,omitempty"`
	Tag         string  `bson:"tag,omitempty" json:"tag,omitempty"`
	Description string  `bson:"description,omitempty" json:"description,omitempty"`
	Content     string  `bson:"content,omitempty" json:"content,omitempty"`
	Category    StepCat `bson:"category,omitempty" json:"category,omitempty"`
	XPAward     int     `bson:"xp_award,omitempty" json:"xp_award,omitempty"`
	CoinsAward  int     `bson:"coins_award,omitempty" json:"coins_award,omitempty"`
	EnergyCost  int     `bson:"energy_cost,omitempty" json:"energy_cost,omitempty"`

	// slice of slices of tags to other Steps
	Children [][]string `bson:"children,omitempty" json:"children,omitempty"`
}

func (s Step) IsEmpty() bool {
	return reflect.DeepEqual(s, Step{})
}
