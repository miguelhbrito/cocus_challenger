package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type Authorization struct{}

//Check if hashed request password is the same from database
func (a Authorization) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//Generate a new hashed password to not show clean password
func (a Authorization) GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
