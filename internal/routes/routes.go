package routes

import (
	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/handlers"
	"LigaFit-AWII2026/internal/middleware"
)

func RegisterRoutes(r chi.Router) {

	// ============================
	// Autenticación
	// ============================

	r.Post("/api/v1/register", handlers.RegisterHandler)
	r.Post("/api/v1/login", handlers.LoginHandler)

	// ============================
	// Rutas protegidas
	// ============================

	r.Group(func(r chi.Router) {

		r.Use(middleware.AuthMiddleware)

		r.Route("/api/v1/deportistas", func(r chi.Router) {
			r.Post("/", handlers.CrearDeportistaHandler)
			r.Get("/", handlers.ObtenerDeportistasHandler)
			r.Get("/{id}", handlers.ObtenerDeportistaPorIDHandler)
			r.Put("/{id}", handlers.ActualizarDeportistaHandler)
			r.Delete("/{id}", handlers.EliminarDeportistaHandler)
		})

		r.Route("/api/v1/inscripciones-deportivas", func(r chi.Router) {
			r.Post("/", handlers.CrearInscripcionDeportivaHandler)
			r.Get("/", handlers.ObtenerInscripcionesDeportivasHandler)
		})

		r.Route("/api/v1/historiales-deportivos", func(r chi.Router) {
			r.Post("/", handlers.CrearHistorialDeportivoHandler)
			r.Get("/", handlers.ObtenerHistorialesDeportivosHandler)
		})

	})
}
