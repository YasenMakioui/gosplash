package services

import (
	"context"
	"fmt"

	"log/slog"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/internal/services/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type FileService struct {
	Repository *repository.FileRepository
	Storage    storage.Storage
}

func NewFileService(repository *repository.FileRepository, storage storage.Storage) *FileService {
	return &FileService{Repository: repository, Storage: storage}
}

// GetFiles will return all files owned by userId.
func (f *FileService) GetFiles(ctx context.Context, userId string) ([]domain.File, error) {
	files, err := f.Repository.FindAll(ctx, userId)

	if err != nil {
		slog.Error("Could not get files")
		return nil, err
	}

	return files, nil
}

// GetFile gets the specific file owned by userId. If no file is found, an error is returned.
func (f *FileService) GetFile(ctx context.Context, fileId string, userId string) (domain.File, error) {
	file, err := f.Repository.Find(ctx, fileId, userId)

	if err != nil {
		if err != pgx.ErrNoRows {
			slog.Error("Could not get file", "fileId", fileId, "userId", userId)
			return file, err
		}

		return file, fmt.Errorf("file not found")
	}

	slog.Debug("Found file", "fileId", fileId, "userId", userId)
	return file, nil
}

// DeleteFile will delete the given fileId owned by userId
func (f *FileService) DeleteFile(ctx context.Context, fileId string, userId string) error {

	file, err := f.Repository.Find(ctx, fileId, userId)

	if err != nil {
		if err != pgx.ErrNoRows {
			slog.Error("Could not get file", "fileId", fileId, "userId", userId)
			return err
		}

		return fmt.Errorf("file not found")
	}

	if err := f.Repository.Delete(ctx, fileId, userId); err != nil {
		if err == pgx.ErrNoRows {
			slog.Debug("File not found", "fileId", fileId, "userId", userId)
			return err
		}
		slog.Error("Could not delete file", "fileId", fileId, "userId", userId)
		return err
	}

	slog.Debug("Deleted file from database", "fileId", fileId)

	if err := f.Storage.Delete(file.StoragePath); err != nil {
		slog.Error("Could not remove file from storage", "path", file.StoragePath)
		return err
	}

	slog.Debug("Deleted file", "path", file.StoragePath)

	return nil
}

// UploadFile uploads the file to the storage
func (f *FileService) UploadFile(ctx context.Context, userId string, uploadedFile multipart.File, handler *multipart.FileHeader) (domain.File, error) {
	fileId := uuid.New().String()
	absolutePath := path.Join("/tmp", "gosplash", fileId, handler.Filename)

	slog.Debug("Uploading file", "file", handler.Filename, "path", absolutePath)

	file := domain.File{
		Id:            fileId,
		UploaderId:    userId,
		FileName:      handler.Filename,
		FileSize:      handler.Size,
		StoragePath:   absolutePath,
		ExpiresAt:     time.Now().Add(time.Hour * 24),
		MaxDownloads:  3,
		Downloads:     0,
		EncryptionKey: "",
		CreatedAt:     time.Now(),
	}

	if err := f.Storage.Upload(absolutePath, uploadedFile); err != nil {
		slog.Error("Failed uploading file", "file", absolutePath)
		return file, err
	}

	slog.Debug("File uploaded successfully")

	if err := f.Repository.Insert(ctx, file); err != nil {
		slog.Error("Failed to save file", "error", err)
		slog.Debug("Deleting file...", "path", absolutePath)

		err := os.Remove(absolutePath)
		if err != nil {
			slog.Error("Could not remove file", "path", absolutePath)
		}
		return file, err
	}

	// Uploaded file successfully, create file object
	return file, nil
}
