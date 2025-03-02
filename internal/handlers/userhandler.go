package handlers

import (
	"encoding/json"
	"github.com/YasenMakioui/gosplash/internal/services"
	"log"
	"net/http"
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
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Could not bind request data to userDTO: %v", err)
		return
	}

	log.Println("Processing signup request")

	// using the userdto attributes create a new user service object
	user, err := services.NewUser(userDTO.Username, userDTO.Email, userDTO.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// If the email is bad, the password is bad length or other validations fail, we get an error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// SignUp the user if the data is correct
	// If further operations as saving to the database fail or another thing fails we return an error
	if err := u.UserService.SignUp(user); err != nil {
		log.Printf("Aborting user creation due to error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("User created successfully")
}
