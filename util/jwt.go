package util

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(userId uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	key := os.Getenv("SECRET")
	claims := &jwt.StandardClaims{
		Subject:   strconv.Itoa(int(userId)),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
