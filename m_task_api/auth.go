package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//nolint:gosec
var secret = []byte("secret")

func generateToken() (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func validateToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid signing method")
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return err
	}
	return nil
}