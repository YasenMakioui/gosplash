package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/YasenMakioui/gosplash/internal/services"
	"log"
	"net/http"
)

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Bind to user dto
	loginDTO := new(LoginDTO)

	if err := json.NewDecoder(r.Body).Decode(&loginDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Could not bind request data to userDTO: %v", err)
		return
	}

	log.Println("Processing login request")

	authService := services.NewAuthService(loginDTO.Username, loginDTO.Password)

	if err := authService.Login(); err != nil {
		log.Printf("Failed authentication for user: %s", loginDTO.Username)
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout")
}
