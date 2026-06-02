package routes

import (
	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/handlers"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/seguimientos", func(r chi.Router) {
		r.Post("/", handlers.CrearSeguimientoHandler)
		r.Get("/", handlers.ObtenerSeguimientosHandler)
		r.Get("/{id}", handlers.ObtenerSeguimientoPorIDHandler)
		r.Get("/deportista/{deportista_id}", handlers.ObtenerSeguimientosPorDeportistaHandler)
		r.Put("/{id}", handlers.ActualizarSeguimientoHandler)
		r.Delete("/{id}", handlers.EliminarSeguimientoHandler)
	})
}
