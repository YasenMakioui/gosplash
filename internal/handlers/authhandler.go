package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/YasenMakioui/gosplash/internal/services"
)

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessTokenDTO struct {
	Token string `json:"token"`
}

type AuthHandler struct {
	AuthService *services.AuthService
	JwtService  *services.JwtService
}

func NewAuthHandler(authService *services.AuthService, jwtService *services.JwtService) *AuthHandler {
	return &AuthHandler{AuthService: authService, JwtService: jwtService}
}

func (a *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	loginDTO := new(LoginDTO)

	if err := json.NewDecoder(r.Body).Decode(&loginDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Debug("Could not bind request data to userDTO", "err", err.Error())
		return
	}

	slog.Debug("Processing login request", "username", loginDTO.Username)

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second) // Use request context
	defer cancel()

	if err := a.AuthService.Login(ctx, loginDTO.Username, loginDTO.Password); err != nil {
		slog.Debug("Failed authentication for user: %s", "username", loginDTO.Username, "err", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := a.JwtService.GenerateToken(loginDTO.Username)

	if err != nil {
		slog.Debug("Failed generating token", "username", loginDTO.Username, "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&AccessTokenDTO{Token: token}); err != nil {
		slog.Debug("Failed encoding token", "username", loginDTO.Username, "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
