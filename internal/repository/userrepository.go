package repository

import (
	"context"
	"fmt"
	"github.com/YasenMakioui/gosplash/internal/db"
	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type UserRepository struct {
	// Add db conection here.
	db *pgxpool.Pool
}

// Constructor that returns based on the user values
func NewUserRepository() (*UserRepository, error) {
	// Inject the database connection
	dbConn, err := db.NewDatabaseConnection()

	if err != nil {
		log.Println("Could not connect to database")
		return nil, err
	}

	return &UserRepository{dbConn}, nil

}

// CheckUser returns error if the user exists
func (r *UserRepository) CheckUser(user *domain.User) error {
	var username string

	query := `SELECT username FROM users WHERE username = $1`

	log.Printf("Executing query: %s\n", query)
	err := r.db.QueryRow(context.Background(), query, user.Username).Scan(&username)

	if err == nil {
		return fmt.Errorf("user exists")
	}

	return nil
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	defer r.db.Close()

	query := `INSERT INTO users (id, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)`

	log.Printf("Executing query: %s\n", query)
	_, err := r.db.Exec(
		context.Background(),
		query,
		user.Id,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (r *UserRepository) GetPasswordHash(username string) (string, error) {
	var passwordHash string

	query := `SELECT password_hash FROM users WHERE username = $1`

	log.Printf("Executing query: %s\n", query)
	err := r.db.QueryRow(context.Background(), query, username).Scan(&passwordHash)

	if err != nil {
		return "", fmt.Errorf("Could not get password hash")
	}
	log.Printf("Got password hash %s", passwordHash)

	return passwordHash, nil
}

// queryExecutor Executes the query safely closing the pool after finishing.
func queryExector() error {
	return nil
}
