package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"LigaFit-AWII2026/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestCompetencias_SinToken_Responde401(t *testing.T) {
	// Preparar
	r := chi.NewRouter()

	r.Route("/api/v1/competencias", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Get("/", ObtenerCompetenciasHandler)
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/competencias/", nil)
	rec := httptest.NewRecorder()

	// Ejecutar
	r.ServeHTTP(rec, req)

	// Verificar
	require.Equal(t, http.StatusUnauthorized, rec.Code)
	require.Contains(t, rec.Body.String(), "token requerido")
}

func TestCrearCompetenciaHandler_CreaCompetenciaValida(t *testing.T) {
	// Preparar
	r := chi.NewRouter()

	r.Post("/api/v1/competencias", CrearCompetenciaHandler)

	body := []byte(`{
		"nombre": "Copa Provincial 2026",
		"deporte_id": 1,
		"categoria_id": 1,
		"fecha_inicio": "2026-08-20",
		"fecha_fin": "2026-08-22",
		"lugar": "Estadio ULEAM",
		"descripcion": "Competencia interclubes",
		"estado": "Programada"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/competencias", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	// Ejecutar
	r.ServeHTTP(rec, req)

	// Verificar
	require.Equal(t, http.StatusCreated, rec.Code)
	require.Contains(t, rec.Body.String(), `"nombre":"Copa Provincial 2026"`)
	require.Contains(t, rec.Body.String(), `"deporte_id":1`)
	require.Contains(t, rec.Body.String(), `"categoria_id":1`)
	require.Contains(t, rec.Body.String(), `"lugar":"Estadio ULEAM"`)
}
