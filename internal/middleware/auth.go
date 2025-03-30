// This middleware just takes the username from the request that is comming from the oauth proxy
// and adds it to the request context.

package middleware

import (
	"context"
	"net/http"
)

type contextKey string

const UsernameKey contextKey = "username"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), UsernameKey, "test") // only for debug
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
