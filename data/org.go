package data

type Org struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// The org's prefix -- example `/api/org/{tag}/steps`
	Tag    string `json:"tag,omitempty" bson:"tag,omitempty"`
	Secret string `json:"-" bson:"secret,omitempty"`
	// Steps...
}

func (o Org) IsEmpty() bool {
	return o == Org{}
}
