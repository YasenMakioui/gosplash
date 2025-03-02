package main

import (
	"encoding/json"
	"fmt"
	"github.com/YasenMakioui/gosplash/internal/config"
	"github.com/YasenMakioui/gosplash/internal/handlers"
	"github.com/YasenMakioui/gosplash/internal/middleware"
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/internal/services"
	"log"
	"net/http"
)

func setupFileHandler() (*handlers.FileHandler, error) {
	fileRepository, err := repository.NewFileRepository()

	if err != nil {
		log.Fatal(err)
	}

	fileService, err := services.NewFileService(fileRepository)

	if err != nil {
		log.Fatal(err)
	}

	userRepository, err := repository.NewUserRepository()

	if err != nil {
		log.Fatal(err)
	}

	userService, err := services.NewUserService(userRepository)

	if err != nil {
		log.Fatal(err)
	}

	return handlers.NewFileHandler(userService, fileService), nil
}

func main() {

	fileHandler, err := setupFileHandler()

	if err != nil {
		log.Fatal(err)
	}

	// Check required environment variables
	config.CheckConfig()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Context().Value(middleware.UserClaimsKey))
		resp := make(map[string]interface{})
		resp["message"] = "gosplash!"
		resp["status"] = http.StatusOK
		resp["user"] = r.Context().Value("username")
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	mux.HandleFunc("POST /auth/login", handlers.LoginHandler)
	mux.HandleFunc("POST /auth/signup", handlers.SignupHandler)

	mux.HandleFunc("GET /files", fileHandler.GetFiles)

	stack := middleware.CreateStack(
		middleware.ValidateJWT,
	)

	http.ListenAndServe(":8080", stack(mux))
}
