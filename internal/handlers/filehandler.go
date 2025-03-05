package handlers

import (
	"encoding/json"
	"github.com/YasenMakioui/gosplash/internal/services"
	"log"
	"net/http"
	"time"
)

type FileDTO struct {
	Filename     string    `json:"filename"`
	Filesize     int64     `json:"filesize"`
	ExpiresAt    time.Time `json:"expires_at"`
	MaxDownloads int64     `json:"max_downloads"`
}
type FileHandler struct {
	UserService *services.UserService
	FileService *services.FileService
}

func NewFileHandler(userService *services.UserService, fileService *services.FileService) *FileHandler {

	return &FileHandler{
		UserService: userService,
		FileService: fileService,
	}
}

func (f *FileHandler) GetFiles(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, err := f.UserService.GetUserUUID(username)

	if err != nil {
		http.Error(w, "Could not get user id", http.StatusInternalServerError)
	}

	files, err := f.FileService.GetUserFiles(userId)

	if err != nil {
		http.Error(w, "Could not get user files", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(files)
}

func (f *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	const maxUploadSize = 50 << 30 // 50GB todo change to env

	username, ok := r.Context().Value("username").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, err := f.UserService.GetUserUUID(username)

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	file, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Failed to retrieve the file", http.StatusBadRequest)
		return
	}

	if handler.Size > maxUploadSize {
		http.Error(w, "File too big", http.StatusBadRequest)
	}

	uploadedFile, err := f.FileService.UploadFile(userId, file, handler)

	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(uploadedFile)
}

func (f *FileHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	fileId := r.PathValue("fileId")

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, err := f.UserService.GetUserUUID(username)

	if err != nil {
		http.Error(w, "Could not get user id", http.StatusInternalServerError)
	}

	log.Printf("Getting file %s from user %s", fileId, userId)

	file, err := f.FileService.GetFile(fileId, userId)

	if err != nil {
		http.Error(w, "Failed to retrieve the file", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(file)
}
