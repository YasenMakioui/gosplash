package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}
