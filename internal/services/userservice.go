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
	repository *repository.UserRepository // Dependency injection
}

func NewUserService(repository *repository.UserRepository, username string, email string, password string) (*UserService, *domain.User, error) {

	userService := new(UserService)
	user := new(domain.User)

	// Validate email
	log.Printf("Checking email: %s", email)
	if err := utils.ValidateEmail(email); err != nil {
		return nil, nil, err
	}

	// Check if password respects the requirements and hash it
	log.Println("Hashing password...")
	passwordHash, err := utils.HashPassword(password)

	if err != nil {
		log.Println("Could not hash password")
		return nil, nil, err
	}

	// Validate user

	log.Printf("Checking username: %s", username)
	if len(username) <= 0 {
		log.Printf("Username is empty")
		return nil, nil, fmt.Errorf("username cannot be empty")
	}

	// Asign values

	user.Id = uuid.New()
	user.Username = username
	user.Email = email
	user.PasswordHash = passwordHash
	user.Role = "user"
	user.CreatedAt = time.Now()

	userService.repository = repository

	return userService, user, nil
}

func (u *UserService) SignUp(user *domain.User) error {
	// Check if the user exists
	log.Println("Checking if user exists")
	if err := u.repository.CheckUser(user); err != nil {
		return fmt.Errorf("username is already taken")
	}

	// Create user
	log.Println("Inserting new user")
	if err := u.repository.CreateUser(user); err != nil {
		return fmt.Errorf("Could not create user")
	}

	return nil
}
