package services

import (
	"fmt"
	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/pkg/utils"
	"github.com/google/uuid"
	"log"
	"time"
)

type UserService struct {
	domain.User                           // Composition
	repository  repository.UserRepository // Dependency injection
}

func NewUserService(username string, email string, password string) (*UserService, error) {

	userService := new(UserService)

	// Validate email
	log.Printf("Checking email: %s", email)
	if err := utils.ValidateEmail(email); err != nil {
		return nil, err
	}

	// Check if password respects the requirements and hash it
	log.Println("Hashing password...")
	passwordHash, err := utils.HashPassword(password)

	if err != nil {
		log.Println("Could not hash password")
		return nil, err
	}

	// Validate user

	log.Printf("Checking username: %s", username)
	if len(username) <= 0 {
		log.Printf("Username is empty")
		return nil, fmt.Errorf("username cannot be empty")
	}

	// Asign values

	userService.Id = uuid.New()
	userService.Username = username
	userService.Email = email
	userService.PasswordHash = passwordHash
	userService.Role = "user"
	userService.CreatedAt = time.Now()

	userRepository, err := repository.NewUserRepository(userService.User)

	if err != nil {
		log.Println("Could not create user repository")
		return nil, err
	}

	userService.repository = *userRepository

	return userService, nil
}

func (u *UserService) SignUp() error {
	// Check if the user exists
	log.Println("Checking if user exists")
	if err := u.repository.CheckUser(); err != nil {
		return fmt.Errorf("username is already taken")
	}

	// Create user
	log.Println("Inserting new user")
	if err := u.repository.CreateUser(); err != nil {
		return fmt.Errorf("Could not create user")
	}

	return nil
}
