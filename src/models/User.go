package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) (err error) {
	if err := DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}
