package storage

type Storage interface {
	// Uploads the file and returns an error on failure
	Upload(path string) error

	// Deletes the file and returns an error on failure
	Delete(path string) error
}
