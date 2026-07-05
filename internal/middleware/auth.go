package middleware

import (
	"context"
	"net/http"
	"strings"

	"LigaFit-AWII2026/internal/services"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UsuarioContextKey contextKey = "usuario"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "token requerido", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &services.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("ligafit-secret-key"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "token invalido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UsuarioContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
