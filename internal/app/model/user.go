package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID               int
	Email            string
	Password         string
	EncryptoPassword string
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptoPassword = enc
	}
	return nil
}
func encryptString(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
