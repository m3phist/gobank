package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaim struct {
	jwt.RegisteredClaims
	UserID    int64 `json:"user_id"`
	ExpiresAt int64 `json:"expires_at"`
}

func CreateToken(user_id int64, signingKey string) (string, error) {
	claims := jwtClaim{
		UserID:    user_id,
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		return "", err
	}
	return string(tokenString), nil
}
