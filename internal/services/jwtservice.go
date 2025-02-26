package services

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	secret []byte
	claims *jwt.MapClaims
}

func NewJwtService(secret []byte, claims *jwt.MapClaims) *JwtService {
	return &JwtService{secret: secret, claims: claims}
}

func (j JwtService) GenerateToken() (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, j.claims)

	tokenString, err := claims.SignedString(j.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
