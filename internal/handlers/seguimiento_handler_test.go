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

func TestSeguimientos_SinToken_Responde401(t *testing.T) {
	// Preparar
	r := chi.NewRouter()

	r.Route("/api/v1/seguimientos", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Get("/", ObtenerSeguimientosHandler)
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/seguimientos/", nil)
	rec := httptest.NewRecorder()

	// Ejecutar
	r.ServeHTTP(rec, req)

	// Verificar
	require.Equal(t, http.StatusUnauthorized, rec.Code)
	require.Contains(t, rec.Body.String(), "token requerido")
}

func TestCrearSeguimientoHandler_CreaSeguimientoValido(t *testing.T) {
	// Preparar
	r := chi.NewRouter()

	r.Post("/api/v1/seguimientos", CrearSeguimientoHandler)

	body := []byte(`{
		"deportista_id": 1,
		"entrenador_id": 1,
		"peso": 70,
		"altura": 1.75
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/seguimientos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	// Ejecutar
	r.ServeHTTP(rec, req)

	// Verificar
	require.Equal(t, http.StatusCreated, rec.Code)
	require.Contains(t, rec.Body.String(), `"deportista_id":1`)
	require.Contains(t, rec.Body.String(), `"entrenador_id":1`)
	require.Contains(t, rec.Body.String(), `"peso":70`)
	require.Contains(t, rec.Body.String(), `"altura":1.75`)
}
