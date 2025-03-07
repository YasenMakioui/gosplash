package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/YasenMakioui/gosplash/internal/db"
	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FileRepository struct {
	db *pgxpool.Pool
}

func NewFileRepository() (*FileRepository, error) {
	dbConn, err := db.NewDatabaseConnection()

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return &FileRepository{dbConn}, nil
}

// GetFiles Will return a list of file objects retrieved from the database. If either the query fails or the row iteration fails, an error is returned as well as a nil value. If there are no results
// An empty slice is returned
func (r *FileRepository) GetFiles(userId string) ([]domain.File, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var files []domain.File

	query := `SELECT * FROM files WHERE uploader_id = $1`

	rows, err := r.db.Query(ctx, query, userId)

	if err != nil {
		slog.Error("Query execution failed")
		return nil, err
	}

	if !rows.Next() {
		slog.Debug("No files found for user", "userId", userId)
		return files, nil
	}

	defer rows.Close()

	for rows.Next() {
		file := domain.File{}
		err := rows.Scan(
			&file.Id,
			&file.UploaderId,
			&file.FileName,
			&file.FileSize,
			&file.StoragePath,
			&file.ExpiresAt,
			&file.MaxDownloads,
			&file.Downloads,
			&file.EncryptionKey,
			&file.CreatedAt,
		)
		if err != nil {
			slog.Error("Failed to scan resultset from database")
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

// GetFile reutrns a file owned by the userId. If the query fails or no results are found a nil and error are returned. If no results are found the error will be of pgx.ErrNoRows
func (r *FileRepository) GetFile(fileId string, userId string) (domain.File, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var file domain.File

	query := `SELECT * FROM files WHERE id = $1 AND uploader_id = $2`

	err := r.db.QueryRow(ctx, query, fileId, userId).Scan(
		&file.Id,
		&file.UploaderId,
		&file.FileName,
		&file.FileSize,
		&file.StoragePath,
		&file.ExpiresAt,
		&file.MaxDownloads,
		&file.Downloads,
		&file.EncryptionKey,
		&file.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Debug("No file found for user", "fileId", fileId, "userId", userId)
			return file, nil
		}

		slog.Error("Failed to execute query", "error", err)
		return file, err
	}

	return file, nil
}

// Delete Will delete the row with fileId and userId. If the query fails. On success a nil value is returned. If no rows where affected a pgx.ErrNoRows is returned
func (r *FileRepository) Delete(fileId string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `DELETE FROM files WHERE id = $1 AND uploader_id = $2`

	commandTag, err := r.db.Exec(ctx, query, fileId, userId)

	if err != nil {
		slog.Debug("Failed to delete file", "error", err)
		return err
	}

	if commandTag.RowsAffected() == 0 {
		slog.Debug("No file deleted (fileId may not exist)", "fileId", fileId, "userId", userId)
		return pgx.ErrNoRows
	}

	slog.Debug(fmt.Sprintf("%v rows affected on delete operation", commandTag.RowsAffected()))
	return nil
}

// Save will save the given file object in the database and will return nil if succeded
func (r *FileRepository) Save(file domain.File) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `INSERT INTO files (id, uploader_id, file_name, file_size, storage_path, expires_at, max_downloads, downloads, encryption_key, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := r.db.Exec(
		ctx,
		query,
		file.Id,
		file.UploaderId,
		file.FileName,
		file.FileSize,
		file.StoragePath,
		file.ExpiresAt,
		file.MaxDownloads,
		file.Downloads,
		file.EncryptionKey,
		file.CreatedAt,
	)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
