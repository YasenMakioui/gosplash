package services

import (
	"fmt"
	"time"

	//"github.com/YasenMakioui/gosplash/src/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

// type jwtToken struct {
// 	token         string
// 	expiration    int64
// 	signingMethod types.SigningMethod
// }

// could come from config
var secretKey = []byte("secret-key")

func NewToken(username string) (string, error) {

	datetime := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      datetime,
		},
	)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", fmt.Errorf("could not create JWT: %v", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// func (t *jwtToken) Verify() error {
// 	token, err := jwt.Parse(t.token, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if !token.Valid {
// 		return fmt.Errorf("invalid token")
// 	}

// 	return nil
// }
