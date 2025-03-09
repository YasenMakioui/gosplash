package services

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/YasenMakioui/gosplash/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	Secret []byte
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// NewJwtService returns a serivce using the secret key from the configuration
func NewJwtService() *JwtService {
	jwtSecret := config.GetSecretKey()

	return &JwtService{
		Secret: jwtSecret,
	}
}

// GenerateToken returns a token for the given user
func (js *JwtService) GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(js.Secret)
}

// ValidateToken Validates the given token and returns the a Claims struct if the validation was successfull
func (js *JwtService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Debug("Invalid token signing method", "method", token.Header["alg"])
			return claims, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return js.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		slog.Debug("Invalid Token", "token", tokenString)
		return claims, fmt.Errorf("invalid token")
	}

	return claims, nil
}
