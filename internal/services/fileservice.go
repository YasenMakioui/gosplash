package services

import (
	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"time"
)

type FileService struct {
	Repository *repository.FileRepository
}

func NewFileService(repository *repository.FileRepository) *FileService {
	return &FileService{Repository: repository}
}

func (f *FileService) GetUserFiles(userId string) ([]domain.File, error) {
	files, err := f.Repository.GetFiles(userId)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func (f *FileService) GetFile(fileId string, userId string) (domain.File, error) {
	file, err := f.Repository.GetFile(fileId, userId)

	if err != nil {
		return domain.File{}, err
	}

	return file, nil
}

func (f *FileService) DeleteFile(fileId string, userId string) error {

	file, err := f.GetFile(fileId, userId)

	if err != nil {
		return err
	}

	if err := os.RemoveAll(path.Dir(file.StoragePath)); err != nil {
		log.Println(err)
		return err
	}
	log.Println("Deleted file ", fileId, " from storage")

	if err := f.Repository.Delete(fileId, userId); err != nil {
		return err
	}
	log.Println("Deleted file ", fileId, " from database")

	return nil
}

func (f *FileService) UploadFile(userId string, uploadedFile multipart.File, handler *multipart.FileHeader) (domain.File, error) {
	fileId := uuid.New().String()
	absolutePath := path.Join("/tmp", "gosplash", fileId, handler.Filename)

	log.Printf("Uploading file %s to %s", handler.Filename, absolutePath)

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

	defer uploadedFile.Close()

	// Create dir

	err := os.MkdirAll(path.Dir(absolutePath), os.ModePerm)

	dst, err := os.Create(absolutePath)

	if err != nil {
		log.Println(err)
		return file, err
	}

	defer dst.Close()

	if _, err := io.Copy(dst, uploadedFile); err != nil {
		return file, err
	}

	log.Println("File uploaded successfully")

	if err := f.Repository.Save(file); err != nil {
		log.Println("Failed to save file", err)
		log.Println("Deleting file...", absolutePath)

		err := os.Remove(absolutePath)
		if err != nil {
			log.Println("Could not remove file", absolutePath)
		}
		return file, err
	}

	// Uploaded file successfully, create file object
	return file, nil
}
