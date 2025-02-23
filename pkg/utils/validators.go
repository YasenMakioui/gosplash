package utils

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

// ValidateEmail Validates email format
func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

// ValidatePassword Validates password hash
func ValidatePassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
