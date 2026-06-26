package routes

import (
	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/handlers"
	"LigaFit-AWII2026/internal/middleware"
)

func RegisterRoutes(r chi.Router) {

	// ==========================
	// AUTENTICACIÓN
	// ==========================
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", handlers.RegisterHandler)
		r.Post("/login", handlers.LoginHandler)
	})

	// ==========================
	// SEGUIMIENTO FÍSICO
	// ==========================
	r.Route("/api/v1/seguimientos", func(r chi.Router) {

		// Todas estas rutas requieren un JWT válido
		r.Use(middleware.AuthMiddleware)

		r.Post("/", handlers.CrearSeguimientoHandler)
		r.Get("/", handlers.ObtenerSeguimientosHandler)
		r.Get("/{id}", handlers.ObtenerSeguimientoPorIDHandler)
		r.Get("/deportista/{deportista_id}", handlers.ObtenerSeguimientosPorDeportistaHandler)
		r.Put("/{id}", handlers.ActualizarSeguimientoHandler)
		r.Delete("/{id}", handlers.EliminarSeguimientoHandler)
	})
}
