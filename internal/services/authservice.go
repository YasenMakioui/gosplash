package services

import (
	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/pkg/utils"
	"log"
)

type AuthService struct {
	domain.User
	Password   string // plain password
	repository repository.UserRepository
}

func NewAuthService(username string, password string) (*AuthService, error) {

	authService := new(AuthService)

	authService.Username = username
	authService.Password = password

	userRepository, err := repository.NewUserRepository(authService.User)

	if err != nil {
		log.Println("Failed to instantiate UserRepository")
		return authService, err
	}

	authService.repository = *userRepository

	return authService, nil
}

func (a *AuthService) Login() error {

	// Create the user repository object

	log.Println("Connecting to the user repository")
	userRepository, err := repository.NewUserRepository(a.User)

	if err != nil {
		log.Println("Failed to instantiate UserRepository")
		return err
	}

	// Get user passwordhash

	passwordHash, err := userRepository.GetPasswordHash()

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
