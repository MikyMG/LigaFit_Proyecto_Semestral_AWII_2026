package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"LigaFit-AWII2026/internal/models"

	"github.com/stretchr/testify/require"
)

func TestObtenerCompetencias_Responde200(t *testing.T) {
	repo := &fakeCompetenciaRepository{}
	router := nuevoRouterCompetencia(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/competencias/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}

func TestCrearCompetencia_Valida_Responde201(t *testing.T) {
	repo := &fakeCompetenciaRepository{}
	router := nuevoRouterCompetencia(repo)

	body := `{
		"nombre": "Copa Provincial 2026",
		"deporte_id": 1,
		"categoria_id": 1,
		"fecha_inicio": "2026-08-20",
		"fecha_fin": "2026-08-22",
		"lugar": "Estadio ULEAM",
		"descripcion": "Competencia interclubes",
		"estado": "Programada"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/competencias/", jsonReq(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
	require.Contains(t, rec.Body.String(), `"nombre":"Copa Provincial 2026"`)
	require.Contains(t, rec.Body.String(), `"deporte_id":1`)
	require.Contains(t, rec.Body.String(), `"categoria_id":1`)
	require.Contains(t, rec.Body.String(), `"lugar":"Estadio ULEAM"`)
}

func TestCrearCompetencia_Invalida_Responde400(t *testing.T) {
	repo := &fakeCompetenciaRepository{}
	router := nuevoRouterCompetencia(repo)

	body := `{
		"nombre": "",
		"deporte_id": 1,
		"categoria_id": 1,
		"fecha_inicio": "2026-08-20",
		"fecha_fin": "2026-08-22",
		"lugar": "Estadio ULEAM",
		"descripcion": "Competencia interclubes"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/competencias/", jsonReq(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "el nombre de la competencia es obligatorio")
}

func TestObtenerCompetencia_NoExiste_Responde404(t *testing.T) {
	repo := &fakeCompetenciaRepository{}
	router := nuevoRouterCompetencia(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/competencias/99", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
	require.Contains(t, rec.Body.String(), "Competencia no encontrada")
}

func TestEliminarCompetencia_Existe_Responde204(t *testing.T) {
	repo := &fakeCompetenciaRepository{
		competencias: []models.Competencia{
			{
				ID:          1,
				Nombre:      "Copa Provincial",
				DeporteID:   1,
				CategoriaID: 1,
				FechaInicio: "2026-08-20",
				FechaFin:    "2026-08-22",
				Lugar:       "Estadio ULEAM",
				Estado:      "Programada",
			},
		},
		nextID: 2,
	}

	router := nuevoRouterCompetencia(repo)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/competencias/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNoContent, rec.Code)
}
