package data

import "reflect"

type Role int

const (
	RoleUser Role = iota
	RoleTeacher
	RoleAdmin
)

type User struct {
	Display      string   `json:"display" bson:"display,omitempty"`   // Tesohh Dockerton
	Username     string   `json:"username" bson:"username,omitempty"` // @tesohh
	PasswordHash string   `json:"passwordhash" bson:"passwordhash,omitempty"`
	XP           int      `json:"xp" bson:"xp,omitempty"`
	Level        int      `json:"level" bson:"level,omitempty"`
	Coins        int      `json:"coins" bson:"coins,omitempty"`
	Role         Role     `json:"role" bson:"role,omitempty"`
	JoinedOrgs   []string `json:"joined_orgs" bson:"joined_orgs,omitempty"`
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}
