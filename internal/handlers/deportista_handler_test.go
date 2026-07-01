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

func TestDeportistas_SinToken_Responde401(t *testing.T) {
	r := chi.NewRouter()

	r.Route("/api/v1/deportistas", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/", ObtenerDeportistasHandler)
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/deportistas/", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	require.Equal(t, http.StatusUnauthorized, rec.Code)
	require.Contains(t, rec.Body.String(), "token requerido")
}

func TestCrearDeportistaHandler_CreaDeportistaValido(t *testing.T) {
	r := chi.NewRouter()

	r.Post("/api/v1/deportistas", CrearDeportistaHandler)

	body := []byte(`{
		"nombre": "Juan Perez",
		"edad": 15,
		"genero": "Masculino",
		"telefono": "0987654321",
		"deporte_id": 1,
		"grupo_id": 1
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/deportistas", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
	require.Contains(t, rec.Body.String(), `"nombre":"Juan Perez"`)
	require.Contains(t, rec.Body.String(), `"edad":15`)
	require.Contains(t, rec.Body.String(), `"categoria":"Juvenil"`)
}
