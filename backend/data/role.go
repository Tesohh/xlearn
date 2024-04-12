package data

type Role string

var (
	RoleUser    Role = "user"
	RoleTeacher Role = "teacher"
	RoleAdmin   Role = "admin"
)

func (r Role) CanAccess(level Role) bool {
	if level == RoleAdmin && r == RoleAdmin {
		return true
	}

	if level == RoleTeacher && (r == RoleTeacher || r == RoleAdmin) {
		return true
	}

	if level == RoleUser {
		return true
	}

	return false
}
