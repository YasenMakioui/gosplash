package main

import (
	"fmt"
	"github.com/YasenMakioui/gosplash/internal/handlers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "gosplash!!")
	})

	mux.HandleFunc("/auth/login", handlers.LoginHandler)
	mux.HandleFunc("/auth/logout", handlers.LogoutHandler)
	mux.HandleFunc("POST /auth/signup", handlers.SignupHandler)

	http.ListenAndServe(":8080", mux)
}
