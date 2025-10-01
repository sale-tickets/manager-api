package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Key []byte
}

type ClaimJWT struct {
	Uuid string `json:"uuid"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func (u *JWT) CreateToken(data ClaimJWT) (string, error) {
	claims := data
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(u.Key)
}

func (u *JWT) ValidateToken(tokenStr string) (*ClaimJWT, error) {
	claims := &ClaimJWT{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(u.Key), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}
