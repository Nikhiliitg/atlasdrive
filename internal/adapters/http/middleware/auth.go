package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/Nikhiliitg/atlasdrive/internal/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if h == "" {
			http.Error(w, "missing token", 401)
			return
		}

		token := strings.TrimPrefix(h, "Bearer ")
		userID, err := auth.ParseToken(token)
		if err != nil {
			http.Error(w, "invalid token", 401)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
