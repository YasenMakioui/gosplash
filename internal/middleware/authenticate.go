package middleware

import (
	"fmt"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticate")
		// Check if jwt is valid. If not, write 401.
		//w.WriteHeader(http.StatusUnauthorized)
		next.ServeHTTP(w, r)
	})
}
