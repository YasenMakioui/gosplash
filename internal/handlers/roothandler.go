package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// Functions that are not methods and can be used by any that imported the package handlers

// RootHandler will write a json with a message
func RootHandler(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]interface{})

	resp["message"] = "gosplash!"
	resp["status"] = http.StatusOK
	resp["user"] = r.Context().Value("username")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
