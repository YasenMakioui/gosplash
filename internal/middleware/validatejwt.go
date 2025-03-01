package middleware

import (
	"github.com/YasenMakioui/gosplash/internal/services"
	"log"
	"net/http"
	"strings"
)

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Auth middleware")

		// Check if it's a public route
		publicPaths := []string{
			"/healthz",
			"/auth/login",
			"/auth/signup",
		}

		for _, path := range publicPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				log.Println("Public path")
				next.ServeHTTP(w, r)
				return
			}
		}

		var secretKey = []byte("your-secret-key")
		jwtService := services.NewJwtService(secretKey)

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			log.Println("No Authorization header")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		log.Println("Validating JWT")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := jwtService.ValidateToken(tokenString)

		if err != nil {
			log.Println("Could not validate token")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		log.Println("Token validated")
		next.ServeHTTP(w, r)
	})
}
