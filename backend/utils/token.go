package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	config *Config
}

type jwtClaim struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}

func NewJWTToken(config *Config) *JWTToken {
	return &JWTToken{config: config}
}

func (j *JWTToken) CreateToken(user_id int64) (string, error) {
	claims := jwtClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
		UserID: user_id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.config.Signing_key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWTToken) VerifyToken(tokenString string) (int64, error) {
	var claims jwtClaim
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid authentication token")
		}
		return []byte(j.config.Signing_key), nil
	})

	if err != nil {
		return 0, fmt.Errorf("invalid authentication token: %v", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid authentication token")
	}

	return claims.UserID, nil
}
