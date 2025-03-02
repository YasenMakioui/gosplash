package services

import (
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/pkg/utils"
	"log"
)

type AuthService struct {
	Username   string
	Password   string // plain password
	repository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository, username string, password string) (*AuthService, error) {

	authService := new(AuthService)

	authService.Username = username
	authService.Password = password
	authService.repository = userRepository

	return authService, nil
}

func (a *AuthService) Login() error {

	// Get user passwordhash

	passwordHash, err := a.repository.GetPasswordHash(a.Username)

	if err != nil {
		log.Println("Failed to get password hash")
		return err
	}

	if err := utils.ValidatePassword(a.Password, passwordHash); err != nil {
		log.Println("Failed to validate password")
		return err
	}

	log.Println("Successfully authenticated user")

	return nil
}
