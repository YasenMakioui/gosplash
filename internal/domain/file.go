package domain

import (
	"github.com/google/uuid"
	"time"
)

type File struct {
	Id            uuid.UUID
	UploaderId    uuid.UUID
	FileName      string
	FileSize      int64
	StoragePath   string
	ExpiresAt     time.Time
	MaxDownloads  int64
	Downloads     int64
	EncryptionKey string
	CreatedAt     time.Time
}
