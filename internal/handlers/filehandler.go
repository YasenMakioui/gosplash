package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/YasenMakioui/gosplash/internal/services"
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
		slog.Error("Couldn't get username from context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, err := f.UserService.GetUserUUID(username)

	if err != nil {
		slog.Error("Couldn't get UUID from user", "username", username, "err", err)
		http.Error(w, "Could not get user id", http.StatusInternalServerError)
		return
	}

	files, err := f.FileService.GetUserFiles(userId)

	if err != nil {
		slog.Error("Couldn't get files from user", "username", username, "err", err)
		http.Error(w, "Could not get user files", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(files); err != nil {
		slog.Error("Couldn't encode files to JSON", "username", username, "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (f *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	const maxUploadSize = 50 << 30 // 50GB todo change to env

	username, ok := r.Context().Value("username").(string)
	if !ok {
		slog.Error("Couldn't get username from context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, err := f.UserService.GetUserUUID(username)

	if err != nil {
		slog.Error("Couldn't retrieve uuid from user", "err", err)
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	file, handler, err := r.FormFile("file")

	if err != nil {
		slog.Error("Couldn't get file from form", "username", username, "err", err)
		http.Error(w, "Failed to retrieve the file", http.StatusBadRequest)
		return
	}

	if handler.Size > maxUploadSize {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	uploadedFile, err := f.FileService.UploadFile(userId, file, handler)

	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(uploadedFile); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (f *FileHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	fileId := r.PathValue("fileId")

	if !ok {
		slog.Error("Couldn't get username from context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, err := f.UserService.GetUserUUID(username)

	if err != nil {
		slog.Error("Couldn't get UUID from user", "username", username, "err", err)
		http.Error(w, "Could not get user id", http.StatusInternalServerError)
		return
	}

	slog.Debug("Getting file %s from user %s", "fileId", fileId, "userId", userId)

	file, err := f.FileService.GetFile(fileId, userId)

	if err != nil {
		slog.Debug("Couldn't get the file", "fileId", fileId, "userId", userId, "err", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(file); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (f *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	fileId := r.PathValue("fileId")

	if !ok {
		slog.Error("Couldn't get username from context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, err := f.UserService.GetUserUUID(username)

	if err != nil {
		slog.Error("Couldn't get UUID from user", "username", username, "err", err)
		http.Error(w, "Could not get user id", http.StatusInternalServerError)
		return
	}

	slog.Debug("Deleting file %s from user %s", fileId, userId)

	if err := f.FileService.DeleteFile(fileId, userId); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Could not delete file", http.StatusInternalServerError)
		return
	}
}
