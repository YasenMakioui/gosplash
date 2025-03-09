package storage

import "fmt"

type LocalStorage struct {
	BaseDir string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{BaseDir: "/tmp/gosplash"}
}

func (l *LocalStorage) Upload(path string) error {
	fmt.Println(path)
	return nil
}

func (l *LocalStorage) Delete(path string) error {
	fmt.Println(path)
	return nil
}
