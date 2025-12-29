package user

import "strings"

func NormalizeUser(u *User) {
	for i, role := range u.Roles {
		u.Roles[i] = strings.TrimSpace(role)
	}
}
