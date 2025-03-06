package main

import (
	"github.com/YasenMakioui/gosplash/internal/config"
	"github.com/YasenMakioui/gosplash/internal/handlers"
	"github.com/YasenMakioui/gosplash/internal/logger"
	"github.com/YasenMakioui/gosplash/internal/middleware"
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/internal/services"
	"log/slog"
	"net/http"
	"os"
)

func setupFileHandler() *handlers.FileHandler {
	fileRepository, err := repository.NewFileRepository()
	if err != nil {
		slog.Error(err.Error())
	}

	fileService := services.NewFileService(fileRepository)

	userRepository, err := repository.NewUserRepository()
	if err != nil {
		slog.Error(err.Error())
	}

	userService := services.NewUserService(userRepository)

	return handlers.NewFileHandler(userService, fileService)
}

func setupUserHandler() *handlers.UserHandler {
	userRepository, err := repository.NewUserRepository()
	if err != nil {
		slog.Error(err.Error())
	}

	userService := services.NewUserService(userRepository)

	return handlers.NewUserHandler(userService)
}

func setupAuthHandler() *handlers.AuthHandler {
	userRepository, err := repository.NewUserRepository()
	if err != nil {
		slog.Error(err.Error())
	}

	authService := services.NewAuthService(userRepository)
	jwtService := services.NewJwtService()

	return handlers.NewAuthHandler(authService, jwtService)
}

func setupHealthHandler() *handlers.HealthHandler {
	return handlers.NewHealthHandler()
}

func main() {

	// Setup logger

	logger.SetupLogger()

	// Setup handlers. If one of them fails, the program won't start
	fileHandler := setupFileHandler()
	userHandler := setupUserHandler()
	authHandler := setupAuthHandler()
	healthHandler := setupHealthHandler()

	// Check required environment variables
	config.CheckConfig()

	// Routes

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/healthz", healthHandler.CheckServerStatus)

	mux.HandleFunc("POST /auth/login", authHandler.LoginHandler)
	mux.HandleFunc("POST /auth/signup", userHandler.Signup)

	mux.HandleFunc("GET /files", fileHandler.GetFiles)
	mux.HandleFunc("POST /files", fileHandler.UploadFile)
	mux.HandleFunc("GET /files/{fileId}/metadata", fileHandler.GetFile)
	mux.HandleFunc("DELETE /files/{fileId}", fileHandler.DeleteFile)

	stack := middleware.CreateStack(
		middleware.ValidateJWT,
	)

	if err := http.ListenAndServe(":8080", stack(mux)); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
