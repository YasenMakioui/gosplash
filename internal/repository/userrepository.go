package repository

import (
	"context"
	"log/slog"

	"github.com/YasenMakioui/gosplash/internal/db"
	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository returns a UserRepository object with the database connection
func NewUserRepository() (*UserRepository, error) {
	dbConn, err := db.NewDatabaseConnection()

	if err != nil {
		slog.Error("Could not connect to database")
		return nil, err
	}

	return &UserRepository{dbConn}, nil
}

// Find will return the user having the given id. If no users are found, an error of pgx.ErrNoRows is returned
func (r *UserRepository) Find(ctx context.Context, userId string) (domain.User, error) {
	var user domain.User

	query := `SELECT username FROM users WHERE id = $1`

	err := r.db.QueryRow(ctx, query, userId).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != pgx.ErrNoRows {
		slog.Debug("N")
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Debug("No such user", "userId", userId)
			return domain.User{}, pgx.ErrNoRows
		}

		slog.Error("Failed to execute query", "error", err)
		return domain.User{}, err
	}

	return user, nil
}

// Insert will insert the given user into the database
func (r *UserRepository) Insert(ctx context.Context, user *domain.User) error {
	defer r.db.Close()

	query := `INSERT INTO users (id, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(
		ctx,
		query,
		user.Id,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
	)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

// GetPasswordHash will return the user's password hash. If nothing is returned, a pgx.ErrNoRows is returned
func (r *UserRepository) GetPasswordHash(ctx context.Context, userId string) (string, error) {
	var passwordHash string

	query := `SELECT password_hash FROM users WHERE id = $1`

	err := r.db.QueryRow(ctx, query, userId).Scan(&passwordHash)

	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Debug("No passwordhash found for user", "userId", userId)
			return "", pgx.ErrNoRows
		}
		slog.Error("Failed to execute query", "error", err)
		return "", err
	}

	return passwordHash, nil
}

// GetPasswordHash will return the user's id. If nothing is returned, a pgx.ErrNoRows is returned
func (r *UserRepository) GetUUID(ctx context.Context, username string) (string, error) {
	var uuid string

	query := `SELECT id FROM users WHERE username = $1`

	err := r.db.QueryRow(ctx, query, username).Scan(&uuid)

	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Debug("No uuid found for user", "user", username)
			return "", pgx.ErrNoRows
		}
		slog.Error("Failed to execute query", "error", err)
		return "", err
	}

	return uuid, nil
}
