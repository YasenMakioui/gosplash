package config

import (
	"log"
	"os"
)

// CheckConfig Looks for required environment variables. If one of these variables are not set, it won't start.
func CheckConfig() {
	requiredEnv := []string{
		"SECRET_KEY",
		"DATABASE_URL",
	}

	for _, env := range requiredEnv {
		if _, ok := os.LookupEnv(env); !ok {
			log.Fatalf("Missing required environment variable: %s", env)
		}
	}
}

func GetSecretKey() []byte {
	secret := os.Getenv("SECRET_KEY")
	return []byte(secret)
}

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
