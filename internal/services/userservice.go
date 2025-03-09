package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/pkg/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	Repository *repository.UserRepository // Dependency injection
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{Repository: repository}
}

// NewUser creates a domain.User based on the UserDTO and validates  username, password and email
func (u *UserService) NewUser(username string, email string, password string) (*domain.User, error) {
	user := new(domain.User)

	slog.Debug("Checking email", "email", email)
	if err := utils.ValidateEmail(email); err != nil {
		return nil, err
	}

	slog.Debug("Hashing password...")
	passwordHash, err := utils.HashPassword(password)

	if err != nil {
		slog.Debug("Could not hash password")
		return nil, err
	}

	slog.Debug("Checking username", "username", username)
	if len(username) <= 0 {
		slog.Debug("Username is empty")
		return nil, fmt.Errorf("username cannot be empty")
	}

	user.Id = uuid.New()
	user.Username = username
	user.Email = email
	user.PasswordHash = passwordHash
	user.Role = "user"
	user.CreatedAt = time.Now()

	return user, nil
}

// SignUp Gives the user to the repository for its insertion in the database
func (u *UserService) SignUp(ctx context.Context, user *domain.User) error {

	if _, err := u.Repository.Find(ctx, user.Id.String()); err != nil {
		if err != pgx.ErrNoRows {
			slog.Error("Could not check user", "error", err)
		}
		slog.Debug("User exists", "user", user)
		return fmt.Errorf("username is already taken")
	}

	slog.Debug("Performing user creation on user", "user", user)

	if err := u.Repository.Insert(ctx, user); err != nil {
		slog.Error("Could not insert user", "user", user, "error", err)
		return fmt.Errorf("could not create user")
	}

	return nil
}

// GetUserUUID will return the user id
func (u *UserService) GetUserUUID(ctx context.Context, username string) (string, error) {
	userUUID, err := u.Repository.GetUUID(ctx, username)

	if err != nil {
		return "", err
	}

	return userUUID, nil
}
