package routes

import (
	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/handlers"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/competencias", func(r chi.Router) {
		r.Post("/", handlers.CrearCompetenciaHandler)
		r.Get("/", handlers.ObtenerCompetenciasHandler)
		r.Get("/{id}", handlers.ObtenerCompetenciaPorIDHandler)
		r.Put("/{id}", handlers.ActualizarCompetenciaHandler)
		r.Delete("/{id}", handlers.EliminarCompetenciaHandler)

		r.Get("/{competencia_id}/participantes", handlers.ObtenerParticipantesPorCompetenciaHandler)
	})

	r.Route("/api/v1/participaciones", func(r chi.Router) {
		r.Post("/", handlers.CrearParticipacionHandler)
		r.Get("/", handlers.ObtenerParticipacionesHandler)
		r.Get("/{id}", handlers.ObtenerParticipacionPorIDHandler)
		r.Put("/{id}", handlers.ActualizarParticipacionHandler)
		r.Delete("/{id}", handlers.EliminarParticipacionHandler)
	})
}
