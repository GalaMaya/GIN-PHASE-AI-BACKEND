package user

import "fmt"

func ValidateCreateUser(u User) error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}

	if len(u.Roles) == 0 {
		return fmt.Errorf("at least one role is required")
	}

	return nil
}
