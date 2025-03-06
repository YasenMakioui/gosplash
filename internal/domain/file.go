package domain

import (
	"time"
)

// File represents a file in the filesystem.
type File struct {
	Id            string
	UploaderId    string
	FileName      string
	FileSize      int64
	StoragePath   string
	ExpiresAt     time.Time
	MaxDownloads  int64
	Downloads     int64
	EncryptionKey string
	CreatedAt     time.Time
}
