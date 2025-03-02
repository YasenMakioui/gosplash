package main

import (
	"encoding/json"
	"fmt"
	"github.com/YasenMakioui/gosplash/internal/config"
	"github.com/YasenMakioui/gosplash/internal/handlers"
	"github.com/YasenMakioui/gosplash/internal/middleware"
	"net/http"
)

func main() {

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

	//mux.HandleFunc("GET /files", handlers.GetFilesHandler)

	stack := middleware.CreateStack(
		middleware.ValidateJWT,
	)

	http.ListenAndServe(":8080", stack(mux))
}
