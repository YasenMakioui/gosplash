package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/YasenMakioui/gosplash/internal/services"
)

type contextKey string

const (
	UserClaimsKey contextKey = "userClaims"
	UsernameKey   contextKey = "username"
)

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if it's a public route
		publicPaths := []string{
			"/healthz",
			"/auth/login",
			"/auth/signup",
		}

		for _, path := range publicPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				next.ServeHTTP(w, r)
				return
			}
		}

		jwtService := services.NewJwtService()

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			slog.Debug("No Authorization header found")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		slog.Debug("Validating JWT", "token", tokenString)

		claims, err := jwtService.ValidateToken(tokenString)

		if err != nil {
			slog.Debug("Could not validate token", "token", tokenString)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		slog.Debug("Token is valid", "token", tokenString)

		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
		ctx = context.WithValue(ctx, UserClaimsKey, claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
