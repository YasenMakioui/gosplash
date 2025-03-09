package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/YasenMakioui/gosplash/internal/services"
)

// UserDTO struct used only to decode the request body
type UserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (u *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {

	// Collect data and encode it to the userdto
	userDTO := new(UserDTO)

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		slog.Error("Could not bind request data to userDTO", "err", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	slog.Debug("Processing signup request", "userDTO", userDTO)

	// using the userdto attributes create a new user service object
	user, err := u.UserService.NewUser(userDTO.Username, userDTO.Email, userDTO.Password)

	if err != nil {
		slog.Error("Could not create user", "err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second) // Use request context
	defer cancel()

	// SignUp the user if the data is correct
	// If further operations as saving to the database fail or another thing fails we return an error
	if err := u.UserService.SignUp(ctx, user); err != nil {
		slog.Error("Aborting user creation due to error: %v\n", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	slog.Debug("User created successfully", "username", userDTO.Username)
	w.WriteHeader(http.StatusCreated)
}
