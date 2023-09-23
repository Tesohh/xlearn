package data

type User struct {
	Display  string `db:"display"`  // Tesohh Dockerton
	Username string `db:"username"` // @tesohh
	XP       int    `db:"xp"`
	Level    int    `db:"lvl"`
	Coins    int    `db:"coins"`
}

func (u User) IsEmpty() bool {
	return u == User{}
}
