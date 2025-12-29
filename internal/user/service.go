package user

func CreateUser(u *User) error {
	if err := ValidateCreateUser(*u); err != nil {
		return err
	}

	NormalizeUser(u)

	return nil
}
