package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/YasenMakioui/gosplash/internal/config"
	"github.com/YasenMakioui/gosplash/internal/handlers"
	"github.com/YasenMakioui/gosplash/internal/logger"
	"github.com/YasenMakioui/gosplash/internal/middleware"
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/internal/services"
	"github.com/YasenMakioui/gosplash/internal/services/storage"
)

func setupFileHandler() *handlers.FileHandler {
	fileRepository, err := repository.NewFileRepository()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	storageBackend := config.GetStorageBackend()
	var gosplashStorage storage.Storage

	switch storageBackend {
	case "LOCAL":
		gosplashStorage = storage.NewLocalStorage()
	// case "S3":
	// 	storage = storage.NewS3Storage() -- NOT IMPLEMENTED YET
	default:
		storageBackend = "LOCAL"
		gosplashStorage = storage.NewLocalStorage()
	}

	slog.Debug("Setting up storage backend", "storage", gosplashStorage)

	fileService := services.NewFileService(fileRepository, gosplashStorage)

	userRepository, err := repository.NewUserRepository()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	userService := services.NewUserService(userRepository)

	return handlers.NewFileHandler(userService, fileService)
}

// func setupUserHandler() *handlers.UserHandler {
// 	userRepository, err := repository.NewUserRepository()
// 	if err != nil {
// 		slog.Error(err.Error())
// 		os.Exit(1)
// 	}

// 	userService := services.NewUserService(userRepository)

// 	return handlers.NewUserHandler(userService)
// }

// func setupAuthHandler() *handlers.AuthHandler {
// 	userRepository, err := repository.NewUserRepository()
// 	if err != nil {
// 		slog.Error(err.Error())
// 		os.Exit(1)
// 	}

// 	authService := services.NewAuthService(userRepository)
// 	jwtService := services.NewJwtService()

// 	return handlers.NewAuthHandler(authService, jwtService)
// }

func setupHealthHandler() *handlers.HealthHandler {
	return handlers.NewHealthHandler()
}

func main() {

	// Setup logger

	logger.SetupLogger()

	slog.Info("Setting up handlers")

	// Setup handlers. If one of them fails, the program won't start
	fileHandler := setupFileHandler()
	// userHandler := setupUserHandler()
	// authHandler := setupAuthHandler()
	healthHandler := setupHealthHandler()

	slog.Info("Checking configuration")

	// Check required environment variables
	config.CheckConfig()

	// Routes

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/healthz", healthHandler.CheckServerStatus)

	// mux.HandleFunc("POST /auth/login", authHandler.LoginHandler) Using third party provider for that. OATUH2 PROXY
	// mux.HandleFunc("POST /auth/signup", userHandler.Signup)

	mux.HandleFunc("GET /files", fileHandler.GetFiles)
	mux.HandleFunc("POST /files", fileHandler.UploadFile)
	mux.HandleFunc("GET /files/{fileId}/metadata", fileHandler.GetFile)
	mux.HandleFunc("DELETE /files/{fileId}", fileHandler.DeleteFile)

	mux.HandleFunc("GET /files/{fileId}", fileHandler.DownloadFile)
	mux.HandleFunc("POST /files/{fileId}/share/{entity}", fileHandler.ShareFile)

	stack := middleware.CreateStack(
	//middleware.ValidateJWT, Not using our own Auth methods.
	//middleware.Auth,
	)

	slog.Info("Listening on port :8080")

	if err := http.ListenAndServe(":8080", stack(mux)); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
