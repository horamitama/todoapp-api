package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(userId uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   string(userId),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	return "", nil
}
