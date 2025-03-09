package storage

import (
	"io"
	"log/slog"
	"mime/multipart"
	"os"
	"path"
)

type LocalStorage struct {
	BaseDir string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{BaseDir: "/tmp/gosplash"}
}

func (l *LocalStorage) Upload(filePath string, uploadedFile multipart.File) error {

	defer uploadedFile.Close()

	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		slog.Error("Could not crete directory", "path", filePath, "error", err)
		return err
	}

	slog.Debug("File directory created successfully")

	dst, err := os.Create(filePath)

	if err != nil {
		slog.Error("Could not create file", "path", filePath, "error", err)
		return err
	}

	slog.Debug("Empty file created successfully")

	defer dst.Close()

	slog.Debug("Uploading file")

	if _, err := io.Copy(dst, uploadedFile); err != nil {
		slog.Error("Error uploading file", "path", filePath, "error", err)
		return err
	}

	return nil
}

func (l *LocalStorage) Delete(filePath string) error {
	if err := os.RemoveAll(path.Dir(filePath)); err != nil {
		slog.Error("Could not remove file from storage", "path", filePath)
		return err
	}
	return nil
}
