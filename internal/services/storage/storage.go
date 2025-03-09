package storage

import "mime/multipart"

type Storage interface {
	// Uploads the file and returns an error on failure
	Upload(path string, uploadedFile multipart.File) error

	// Deletes the file and returns an error on failure
	Delete(path string) error
}
