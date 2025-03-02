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
