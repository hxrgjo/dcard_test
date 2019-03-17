package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret string

func SetSecret(v string) {
	secret = v
}

func Sign(userID string) (authToken string, err error) {
	if secret == "" {
		return "", fmt.Errorf("secret not set")
	}
	// Set claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().AddDate(0, 0, 1).Unix(),
	})
	authToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return
	}
	return
}
