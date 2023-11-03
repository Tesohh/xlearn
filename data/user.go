package data

import "reflect"

type Role int

const (
	RoleUser Role = iota
	RoleTeacher
	RoleAdmin
)

type UserSettings struct {
	Language string `json:"language" bson:"language,omitempty"`
	Theme    string `json:"theme" bson:"theme,omitempty"`
	// AutoNightMode bool   `json:"auto_night_mode" bson:"auto_night_mode,omitempty"`
}

func (s UserSettings) IsEmpty() bool {
	return s == UserSettings{}
}

type User struct {
	Display      string       `json:"display" bson:"display,omitempty"`   // Tesohh Dockerton
	Username     string       `json:"username" bson:"username,omitempty"` // @tesohh
	PasswordHash string       `json:"passwordhash" bson:"passwordhash,omitempty"`
	XP           int          `json:"xp" bson:"xp,omitempty"`
	Level        int          `json:"level" bson:"level,omitempty"`
	Coins        int          `json:"coins" bson:"coins,omitempty"`
	Role         Role         `json:"role" bson:"role,omitempty"`
	JoinedOrgs   []string     `json:"joined_orgs" bson:"joined_orgs,omitempty"`
	Settings     UserSettings `json:"settings" bson:"settings,omitempty"`
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}

func (u User) GetTag() string {
	return u.Username
}
