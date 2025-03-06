// Package db This package is used to get database connection and information
package db

import (
	"context"
	"fmt"
	"github.com/YasenMakioui/gosplash/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

// NewDatabaseConnection Will return a pointer to the database connection
func NewDatabaseConnection() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), config.GetDatabaseURL())

	if err != nil {
		slog.Error(fmt.Sprintf("Error connecting to database: %s", err))
		return nil, err
	}

	return pool, nil
}
