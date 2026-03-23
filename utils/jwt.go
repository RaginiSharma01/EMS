package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(empID string, email string) (string, error) {
	claims := jwt.MapClaims{
		"empId": empID,
		"email": email,
		"exp":   jwt.TimeFunc().Add(20 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
