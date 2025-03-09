package config

import (
	"log/slog"
	"os"
	"strings"
)

// CheckConfig Looks for required environment variables. If one of these variables are not set, an error is shown and the code exits with status code 1.
func CheckConfig() {
	requiredEnv := []string{
		"SECRET_KEY",
		"DATABASE_URL",
	}

	for _, env := range requiredEnv {
		if _, ok := os.LookupEnv(env); !ok {
			slog.Error("Missing required environment variable", "name", env)
			os.Exit(1)
		}
	}
}

// GetSecretKey returns the secret key in a slice of bytes for later use
func GetSecretKey() []byte {
	secret := os.Getenv("SECRET_KEY")
	return []byte(secret)
}

// GetDatabaseURL returns the connection string for the database
func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

// GetLogLevel returns the log level inside slog.HandlerOptions for its implementation into the logger
func GetLogLevel() slog.Level {
	switch strings.ToUpper(os.Getenv("LOG_LEVEL")) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	}

	return slog.LevelInfo
}

func GetStorageBackend() string {
	return os.Getenv("STORAGE")
}
