package middleware

import (
	"context"
	"github.com/YasenMakioui/gosplash/internal/services"
	"log"
	"net/http"
	"strings"
)

type contextKey string

const UserClaimsKey contextKey = "userClaims"

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
		
		jwtService := services.NewJwtService()

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			log.Println("No Authorization header")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		log.Println("Validating JWT")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtService.ValidateToken(tokenString)

		if err != nil {
			log.Println("Could not validate token")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		log.Println("Token validated")

		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
		ctx = context.WithValue(ctx, "username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
