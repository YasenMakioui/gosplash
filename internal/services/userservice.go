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
	Repository *repository.UserRepository // Dependency injection
}

func NewUserService(repository *repository.UserRepository) (*UserService, error) {

	userService := new(UserService)
	userService.Repository = repository

	return userService, nil
}

func NewUser(username string, email string, password string) (*domain.User, error) {
	user := new(domain.User)

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

	user.Id = uuid.New()
	user.Username = username
	user.Email = email
	user.PasswordHash = passwordHash
	user.Role = "user"
	user.CreatedAt = time.Now()

	return user, nil
}

func (u *UserService) SignUp(user *domain.User) error {
	// Check if the user exists
	log.Println("Checking if user exists")
	if err := u.Repository.CheckUser(user); err != nil {
		return fmt.Errorf("username is already taken")
	}

	// Create user
	log.Println("Inserting new user")
	if err := u.Repository.CreateUser(user); err != nil {
		return fmt.Errorf("Could not create user")
	}

	return nil
}

func (u *UserService) GetUserUUID(username string) (string, error) {
	userUUID, err := u.Repository.GetUUID(username)

	if err != nil {
		return "", err
	}

	return userUUID, nil
}
