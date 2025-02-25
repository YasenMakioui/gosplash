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
	domain.User
	// Add db conection here.
	db *pgxpool.Pool
}

// Constructor that returns based on the user values
func NewUserRepository(u domain.User) (*UserRepository, error) {
	// Inject the database connection
	dbConn, err := db.NewDatabaseConnection()

	if err != nil {
		log.Println("Could not connect to database")
		return nil, err
	}

	return &UserRepository{u, dbConn}, nil

}

// CheckUser returns error if the user exists
func (u *UserRepository) CheckUser() error {
	var username string

	query := `SELECT username FROM users WHERE username = $1`

	log.Printf("Executing query: %s\n", query)
	err := u.db.QueryRow(context.Background(), query, u.Username).Scan(&username)

	if err == nil {
		return fmt.Errorf("user exists")
	}

	return nil
}

func (u *UserRepository) CreateUser() error {
	defer u.db.Close()

	query := `INSERT INTO users (id, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)`

	log.Printf("Executing query: %s\n", query)
	_, err := u.db.Exec(
		context.Background(),
		query,
		u.Id,
		u.Username,
		u.Email,
		u.PasswordHash,
		u.CreatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u *UserRepository) GetPasswordHash() (string, error) {
	var passwordHash string

	query := `SELECT password_hash FROM users WHERE username = $1`

	log.Printf("Executing query: %s\n", query)
	err := u.db.QueryRow(context.Background(), query, u.Username).Scan(&passwordHash)

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
