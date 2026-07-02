package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"LigaFit-AWII2026/internal/models"

	"github.com/stretchr/testify/require"
)

func TestObtenerSeguimientos_Responde200(t *testing.T) {
	repo := &fakeSeguimientoRepository{}
	router := nuevoRouterSeguimiento(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/seguimientos/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}

func TestCrearSeguimiento_Valido_Responde201(t *testing.T) {
	repo := &fakeSeguimientoRepository{}
	router := nuevoRouterSeguimiento(repo)

	body := `{
		"deportista_id": 1,
		"entrenador_id": 1,
		"peso": 70,
		"altura": 1.75
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/seguimientos/", jsonReq(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
	require.Contains(t, rec.Body.String(), `"deportista_id":1`)
	require.Contains(t, rec.Body.String(), `"entrenador_id":1`)
}

func TestCrearSeguimiento_Invalido_Responde400(t *testing.T) {
	repo := &fakeSeguimientoRepository{}
	router := nuevoRouterSeguimiento(repo)

	body := `{
		"deportista_id": 1,
		"entrenador_id": 0,
		"peso": 70,
		"altura": 1.75
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/seguimientos/", jsonReq(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "el entrenador_id es obligatorio")
}

func TestObtenerSeguimiento_NoExiste_Responde404(t *testing.T) {
	repo := &fakeSeguimientoRepository{}
	router := nuevoRouterSeguimiento(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/seguimientos/99", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
	require.Contains(t, rec.Body.String(), "Seguimiento no encontrado")
}

func TestEliminarSeguimiento_Existe_Responde204(t *testing.T) {
	repo := &fakeSeguimientoRepository{
		seguimientos: []models.SeguimientoFisico{
			{
				ID:           1,
				DeportistaID: 1,
				EntrenadorID: 1,
				Peso:         70,
				Altura:       1.75,
			},
		},
		nextID: 2,
	}

	router := nuevoRouterSeguimiento(repo)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/seguimientos/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNoContent, rec.Code)
}
