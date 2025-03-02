package handlers

import (
	"encoding/json"
	"github.com/YasenMakioui/gosplash/internal/config"
	"github.com/YasenMakioui/gosplash/internal/repository"
	"github.com/YasenMakioui/gosplash/internal/services"
	"log"
	"net/http"
)

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessTokenDTO struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	userRepository, err := repository.NewUserRepository()

	if err != nil {
		log.Println(err)
	}
	// Bind to user dto
	loginDTO := new(LoginDTO)

	if err := json.NewDecoder(r.Body).Decode(&loginDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Could not bind request data to userDTO: %v", err)
		return
	}

	log.Println("Processing login request")

	// Get the auth service and log in the user
	authService, err := services.NewAuthService(userRepository, loginDTO.Username, loginDTO.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Could not bind request data to userDTO: %v", err)
	}

	if err := authService.Login(); err != nil {
		log.Printf("Failed authentication for user: %s", loginDTO.Username)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var secretKey = config.GetSecretKey()

	jwtService := services.NewJwtService(secretKey)

	token, err := jwtService.GenerateToken(loginDTO.Username)

	if err != nil {
		log.Printf("Failed generating token: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&AccessTokenDTO{Token: token}); err != nil {
		log.Printf("Failed encoding token: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
