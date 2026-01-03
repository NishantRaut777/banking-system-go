package utils

import "golang.org/x/crypto/bcrypt"

func HashString(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(bytes), err
}
