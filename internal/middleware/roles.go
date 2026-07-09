package middleware

import (
	"net/http"

	"LigaFit-AWII2026/internal/services"
)

func RequireRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims, ok := r.Context().Value(UsuarioContextKey).(*services.Claims)
			if !ok {
				http.Error(w, "No autorizado", http.StatusUnauthorized)
				return
			}

			if claims.Rol != role {
				http.Error(w, "Acceso denegado", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
