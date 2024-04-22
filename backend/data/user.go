package data

import (
	"reflect"
)

type User struct {
	Tag          string `json:"tag" bson:"tag,omitempty"`
	PasswordHash string `json:"-" bson:"passwordhash,omitempty"`
	// Pin          string `json:"pin" bson:"pin,omitempty"`
	// RecoverAttempts int    `json:"recover_attempts" bson:"recover_attempts,omitempty"`
	// XP    int  `json:"xp" bson:"xp,omitempty"`
	// Level int  `json:"level" bson:"level,omitempty"`
	// Coins int  `json:"coins" bson:"coins,omitempty"`
	Role Role `json:"role" bson:"role,omitempty"`
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}

func (u User) GetTag() string {
	return u.Tag
}
