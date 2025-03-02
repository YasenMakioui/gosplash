package services

import (
	"github.com/YasenMakioui/gosplash/internal/domain"
	"github.com/YasenMakioui/gosplash/internal/repository"
)

type FileService struct {
	Repository *repository.FileRepository
}

func NewFileService(repository *repository.FileRepository) (*FileService, error) {

	fileService := new(FileService)
	fileService.Repository = repository

	return fileService, nil
}

func (f *FileService) GetUserFiles(userId string) ([]domain.File, error) {
	files, err := f.Repository.GetFiles(userId)

	if err != nil {
		return nil, err
	}

	return files, nil
}
