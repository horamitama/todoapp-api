package util

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	key := os.Getenv("SECRET")
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUserIDFromAuthHeader(authHeader string) (uint, error) {
	if authHeader == "" {
		return 0, errors.New("unauthorized")
	}

	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secretKey := os.Getenv("SECRET")
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	claim := claims["user_id"]
	if claim == nil {
		return 0, errors.New("user_id is invalid")
	}

	userID := int(claims["user_id"].(float64))
	return uint(userID), nil
}
