package utils

import "golang.org/x/crypto/bcrypt"

func HashString(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(bytes), err
}

// CompareHash compares hashed value with plain input
func CompareHash(hash string, plain string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(plain),
	)
	return err == nil
}
