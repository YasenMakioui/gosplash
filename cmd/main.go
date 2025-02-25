package main

import (
	"fmt"
	"github.com/YasenMakioui/gosplash/internal/handlers"
	"github.com/YasenMakioui/gosplash/internal/middleware"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "gosplash!!")
	})

	mux.HandleFunc("POST /auth/login", handlers.LoginHandler)
	mux.HandleFunc("/auth/logout", handlers.LogoutHandler)
	mux.HandleFunc("POST /auth/signup", handlers.SignupHandler)

	stack := middleware.CreateStack()

	http.ListenAndServe(":8080", stack(mux))
}
