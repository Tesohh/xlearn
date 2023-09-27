package data

type Role int

const (
	RoleUser Role = iota
	RoleTeacher
	RoleAdmin
)

type User struct {
	Display      string `json:"display,omitempty" bson:"display,omitempty"`   // Tesohh Dockerton
	Username     string `json:"username,omitempty" bson:"username,omitempty"` // @tesohh
	PasswordHash string `json:"passwordhash,omitempty" bson:"passwordhash,omitempty"`
	XP           int    `json:"xp,omitempty" bson:"xp,omitempty"`
	Level        int    `json:"level,omitempty" bson:"level,omitempty"`
	Coins        int    `json:"coins,omitempty" bson:"coins,omitempty"`
	Role         Role   `json:"role,omitempty" bson:"role,omitempty"`
}

func (u User) IsEmpty() bool {
	return u == User{}
}
