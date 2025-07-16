package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func GenerateJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}