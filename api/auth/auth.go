package auth

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret string

func SetSecret(v string) {
	secret = v
}

func Sign(userID int64) (authToken string, err error) {
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

// Please see the documentation: http://jwt.io/
func Verify(authToken string) (userID int64, err error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		return 0, fmt.Errorf("claims valid failed")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("wrong claims type")
	}

	isVerified := claims.VerifyExpiresAt(time.Now().Unix(), true)
	if !isVerified {
		return 0, errors.New("JWT Token has expired")
	}

	switch claims["user_id"].(type) {
	case float64:
		userID = (int64)(claims["user_id"].(float64))
	default:
		return 0, errors.New("convert user id to int64 error")
	}
	return userID, nil
}
