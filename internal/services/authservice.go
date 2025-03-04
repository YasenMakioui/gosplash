package services

import (
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/pkg/utils"
	"log"
)

type AuthService struct {
	repository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) (*AuthService, error) {

	authService := new(AuthService)

	authService.repository = userRepository

	return authService, nil
}

func (a *AuthService) Login(username string, password string) error {

	// Get user passwordhash

	passwordHash, err := a.repository.GetPasswordHash(username)

	if err != nil {
		log.Println("Failed to get password hash")
		return err
	}

	if err := utils.ValidatePassword(password, passwordHash); err != nil {
		log.Println("Failed to validate password")
		return err
	}

	log.Println("Successfully authenticated user")

	return nil
}
