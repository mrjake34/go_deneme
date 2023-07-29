package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	UserID    string `json:"_id" bson:"_id, omitempty"`
	Fullname  string `json:"fullName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	IP        string `json:"ip"`
	CreatedAt string `json:"createdAt"`
	Role      string `json:"role"`
}

func HashPasswordUser(u *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
