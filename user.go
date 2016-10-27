package main

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int64
	Name     string
	Password string
}

func HashedPassword(password string) ([]byte, error) {
	hex := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(hex, 2)
	if err != nil {
		return nil, err
	}
	return hash, err

}
