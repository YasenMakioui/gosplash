package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func NewDatabaseConnection() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:mysecretpassword@localhost:5432")

	if err != nil {
		log.Printf("Aborting database creation due to error: %v\n", err)
		return nil, err
	}

	return dbpool, nil
}
