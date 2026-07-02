package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"LigaFit-AWII2026/internal/models"

	"github.com/stretchr/testify/require"
)

func TestObtenerDeportistas_Responde200(t *testing.T) {
	repo := &fakeDeportistaRepository{}
	router := nuevoRouterDeportista(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/deportistas/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}

func TestCrearDeportista_Valido_Responde201(t *testing.T) {
	repo := &fakeDeportistaRepository{}
	router := nuevoRouterDeportista(repo)

	body := `{
		"nombre": "Juan Perez",
		"edad": 15,
		"genero": "Masculino",
		"telefono": "0987654321",
		"deporte_id": 1,
		"grupo_id": 1
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/deportistas/", jsonReq(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
	require.Contains(t, rec.Body.String(), `"nombre":"Juan Perez"`)
	require.Contains(t, rec.Body.String(), `"categoria":"Juvenil"`)
}

func TestCrearDeportista_Invalido_Responde400(t *testing.T) {
	repo := &fakeDeportistaRepository{}
	router := nuevoRouterDeportista(repo)

	body := `{
		"nombre": "Jo",
		"edad": 15,
		"genero": "Masculino",
		"telefono": "0987654321",
		"deporte_id": 1,
		"grupo_id": 1
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/deportistas/", jsonReq(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "el nombre debe tener al menos 3 caracteres")
}

func TestObtenerDeportista_NoExiste_Responde404(t *testing.T) {
	repo := &fakeDeportistaRepository{}
	router := nuevoRouterDeportista(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/deportistas/99", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
	require.Contains(t, rec.Body.String(), "Deportista no encontrado")
}

func TestEliminarDeportista_Existe_Responde204(t *testing.T) {
	repo := &fakeDeportistaRepository{
		deportistas: []models.Deportista{
			{
				ID:        1,
				Nombre:    "Juan Perez",
				Edad:      15,
				Genero:    "Masculino",
				Telefono:  "0987654321",
				DeporteID: 1,
				Categoria: "Juvenil",
				GrupoID:   1,
			},
		},
		nextID: 2,
	}

	router := nuevoRouterDeportista(repo)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/deportistas/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNoContent, rec.Code)
}
