package services

import (
	"context"
	"log/slog"

	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/pkg/utils"
	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	Repository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{Repository: userRepository}
}

// Login will perform a login operation and will return a nil if succeded
func (a *AuthService) Login(ctx context.Context, username string, password string) error {

	userId, err := a.Repository.GetUUID(ctx, username)

	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Debug("User does not exist", "error", err)
			return err
		}
		slog.Error("Failed to get user id")
		return err
	}

	passwordHash, err := a.Repository.GetPasswordHash(ctx, userId)

	if err != nil {
		slog.Error("Failed to get password hash")
		return err
	}

	if err := utils.ValidatePassword(password, passwordHash); err != nil {
		slog.Debug("Failed to validate password", "user", username, "passwordHash", passwordHash)
		return err
	}

	slog.Debug("Successfully authenticated user", "user", username)

	return nil
}
