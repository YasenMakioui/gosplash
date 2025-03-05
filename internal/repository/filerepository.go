package repository

import (
	"context"
	"github.com/YasenMakioui/gosplash/internal/db"
	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type FileRepository struct {
	db *pgxpool.Pool
}

func NewFileRepository() (*FileRepository, error) {
	// Inject the database connection
	dbConn, err := db.NewDatabaseConnection()

	if err != nil {
		log.Println("Could not connect to database")
		return nil, err
	}

	return &FileRepository{dbConn}, nil
}

func (r *FileRepository) GetFiles(userId string) ([]domain.File, error) {

	var files []domain.File

	query := `SELECT * FROM files WHERE uploader_id = $1`

	log.Printf("Executing query: %s\n", query)

	rows, err := r.db.Query(context.Background(), query, userId)

	if err != nil {
		return files, err
	}

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
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func (r *FileRepository) GetFile(fileId string, userId string) (domain.File, error) {
	var file domain.File

	query := `SELECT * FROM files WHERE id = $1 AND uploader_id = $2`
	log.Printf("Executing query: SELECT * FROM files WHERE id = %s AND uploader_id = %s", fileId, userId)

	err := r.db.QueryRow(context.Background(), query, fileId, userId).Scan(
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
		log.Println(err)
		return file, err
	}

	return file, nil
}

func (r *FileRepository) Delete(fileId string, userId string) error {
	query := `DELETE FROM files WHERE id = $1 AND uploader_id = $2`

	_, err := r.db.Query(context.Background(), query, fileId, userId)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *FileRepository) Save(file domain.File) error {
	
	query := `INSERT INTO files (id, uploader_id, file_name, file_size, storage_path, expires_at, max_downloads, downloads, encryption_key, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	log.Printf("Executing query: %s\n", query)
	_, err := r.db.Exec(
		context.Background(),
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
		log.Println(err.Error())
		return err
	}

	return nil
}
