package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Fullname string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(u *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
