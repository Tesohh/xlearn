package data

import "reflect"

type StepCat string

const (
	StepCatLesson   StepCat = "lesson"
	StepCatExercise StepCat = "exercise"
	StepCatProject  StepCat = "project"
)

type Step struct {
	Name        string  `bson:"name,omitempty" json:"name"`
	Tag         string  `bson:"tag,omitempty" json:"tag"`
	Description string  `bson:"description,omitempty" json:"description"`
	Content     string  `bson:"content,omitempty" json:"content"`
	Category    StepCat `bson:"category,omitempty" json:"category"`
	XPAward     int     `bson:"xp_award,omitempty" json:"xp_award"`
	CoinsAward  int     `bson:"coins_award,omitempty" json:"coins_award"`
	EnergyCost  int     `bson:"energy_cost,omitempty" json:"energy_cost"`

	// slice of slices of tags to other Steps
	Children [][]string `bson:"children,omitempty" json:"children"`
}

func (s Step) IsEmpty() bool {
	return reflect.DeepEqual(s, Step{})
}
