package models

import "golang.org/x/crypto/bcrypt"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CheckPasswordLogin(u *User, l *Login) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password))
	return err
}


