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
	// Since we need to defer close everytime we execute a query, we should extract the exec part on another function.
	defer u.db.Close()

	var username string

	query := `SELECT username FROM users WHERE username = $1`

	fmt.Printf("Executing query: %s\n", query)
	err := u.db.QueryRow(context.Background(), query, u.Username).Scan(&username)

	if err == nil {
		return fmt.Errorf("user exists")
	}

	return nil
}

func (u *UserRepository) CreateUser() error {
	return nil
}

// queryExecutor Executes the query safely closing the pool after finishing.
func queryExector() error {
	return nil
}
