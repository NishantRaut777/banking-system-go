package utils

import (
	"fmt"
	"hash/fnv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// takes userid and returns unique bankaccount value
func GenerateAccountNumber(userID uuid.UUID) string {
	year := time.Now().Year()

	// Take last 6 digits from UUID hash
	hash := fnv.New32a()
	hash.Write([]byte(userID.String()))
	unique := hash.Sum32() % 1_000_000

	return fmt.Sprintf("BANK%d%06d", year, unique)
}

var jwtSecret []byte

// SetJWTSecret should be called once at startup
func SetJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}

func GetJWTSecret() []byte {
	return jwtSecret
}

// GenerateJWT creates a signed JWT token
func GenerateJWT(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(), // UUID as string
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
