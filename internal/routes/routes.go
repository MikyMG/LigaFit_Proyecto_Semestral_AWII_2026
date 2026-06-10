package routes

import (
	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/handlers"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/deportistas", func(r chi.Router) {
		r.Post("/", handlers.CrearDeportistaHandler)
		r.Get("/", handlers.ObtenerDeportistasHandler)
		r.Get("/{id}", handlers.ObtenerDeportistaPorIDHandler)
		r.Put("/{id}", handlers.ActualizarDeportistaHandler)
		r.Delete("/{id}", handlers.EliminarDeportistaHandler)
	})
}
