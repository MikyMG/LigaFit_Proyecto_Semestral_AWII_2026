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
	// COMPETENCIAS (Solo Admin)
	// ==========================
	r.Route("/api/v1/competencias", func(r chi.Router) {

		r.Use(middleware.AuthMiddleware)
		r.Use(middleware.RequireRole("admin"))

		r.Post("/", handlers.CrearCompetenciaHandler)
		r.Get("/", handlers.ObtenerCompetenciasHandler)
		r.Get("/{id}", handlers.ObtenerCompetenciaPorIDHandler)
		r.Put("/{id}", handlers.ActualizarCompetenciaHandler)
		r.Delete("/{id}", handlers.EliminarCompetenciaHandler)

		r.Get("/{competencia_id}/participantes", handlers.ObtenerParticipantesPorCompetenciaHandler)
		r.Get("/{competencia_id}/resultados", handlers.ObtenerResultadosPorCompetenciaHandler)
	})

	// ==========================
	// PARTICIPACIONES (Entrenador)
	// ==========================
	r.Route("/api/v1/participaciones", func(r chi.Router) {

		r.Use(middleware.AuthMiddleware)
		r.Use(middleware.RequireRole("entrenador"))

		r.Post("/", handlers.CrearParticipacionHandler)
		r.Get("/", handlers.ObtenerParticipacionesHandler)
		r.Get("/{id}", handlers.ObtenerParticipacionPorIDHandler)
		r.Put("/{id}", handlers.ActualizarParticipacionHandler)
		r.Delete("/{id}", handlers.EliminarParticipacionHandler)
	})

	// ==========================
	// RESULTADOS (Admin)
	// ==========================
	r.Route("/api/v1/resultados-competencia", func(r chi.Router) {

		r.Use(middleware.AuthMiddleware)
		r.Use(middleware.RequireRole("admin"))

		r.Post("/", handlers.CrearResultadoCompetenciaHandler)
		r.Get("/", handlers.ObtenerResultadosCompetenciaHandler)
		r.Get("/{id}", handlers.ObtenerResultadoCompetenciaPorIDHandler)
		r.Put("/{id}", handlers.ActualizarResultadoCompetenciaHandler)
		r.Delete("/{id}", handlers.EliminarResultadoCompetenciaHandler)
	})
}
