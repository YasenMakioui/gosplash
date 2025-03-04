package handlers

import (
	"encoding/json"
	"github.com/YasenMakioui/gosplash/internal/services"
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

func (f *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {

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
