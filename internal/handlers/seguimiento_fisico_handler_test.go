package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"
	"LigaFit-AWII2026/internal/storage"
)

func resetHandlerTest() {
	storage.Seguimientos = nil
	storage.SeguimientoIDCounter = 1
	services.SetSeguimientoRepository(storage.NewSeguimientoMemoryRepository())
}

func withChiContext(req *http.Request, urlParams map[string]string) *http.Request {
	chiCtx := chi.NewRouteContext()
	for k, v := range urlParams {
		chiCtx.URLParams.Add(k, v)
	}
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
}

func TestCrearSeguimientoHandler_Exitoso(t *testing.T) {
	resetHandlerTest()
	body := `{"deportista_id":1,"entrenador_id":1,"peso":70,"altura":1.75}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/seguimientos", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CrearSeguimientoHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusCreated)
	}

	var result models.SeguimientoFisico
	json.NewDecoder(w.Body).Decode(&result)
	if result.ID == 0 {
		t.Error("ID no deberia ser 0")
	}
}

func TestCrearSeguimientoHandler_DatosInvalidos(t *testing.T) {
	resetHandlerTest()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/seguimientos", bytes.NewBufferString("{"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CrearSeguimientoHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusBadRequest)
	}
}

func TestCrearSeguimientoHandler_ValidacionFallida(t *testing.T) {
	resetHandlerTest()
	body := `{"deportista_id":0,"entrenador_id":1,"peso":70,"altura":1.75}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/seguimientos", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CrearSeguimientoHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusBadRequest)
	}
}

func TestObtenerSeguimientosHandler_Vacio(t *testing.T) {
	resetHandlerTest()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/seguimientos", nil)
	w := httptest.NewRecorder()

	ObtenerSeguimientosHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusOK)
	}

	var result []models.SeguimientoFisico
	json.NewDecoder(w.Body).Decode(&result)
	if len(result) != 0 {
		t.Errorf("devolvio %d elementos, esperaba 0", len(result))
	}
}

func TestObtenerSeguimientoPorIDHandler_Encontrado(t *testing.T) {
	resetHandlerTest()
	creado, _ := services.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/seguimientos/1", nil)
	req = withChiContext(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	ObtenerSeguimientoPorIDHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusOK)
	}

	var result models.SeguimientoFisico
	json.NewDecoder(w.Body).Decode(&result)
	if result.ID != creado.ID {
		t.Errorf("ID = %d; want %d", result.ID, creado.ID)
	}
}

func TestObtenerSeguimientoPorIDHandler_NoEncontrado(t *testing.T) {
	resetHandlerTest()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/seguimientos/999", nil)
	req = withChiContext(req, map[string]string{"id": "999"})
	w := httptest.NewRecorder()

	ObtenerSeguimientoPorIDHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusNotFound)
	}
}

func TestObtenerSeguimientoPorIDHandler_IDInvalido(t *testing.T) {
	resetHandlerTest()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/seguimientos/abc", nil)
	req = withChiContext(req, map[string]string{"id": "abc"})
	w := httptest.NewRecorder()

	ObtenerSeguimientoPorIDHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusBadRequest)
	}
}

func TestObtenerSeguimientosPorDeportistaHandler_Filtra(t *testing.T) {
	resetHandlerTest()
	services.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})
	services.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 2, EntrenadorID: 1, Peso: 80, Altura: 1.80})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/seguimientos/deportista/1", nil)
	req = withChiContext(req, map[string]string{"deportista_id": "1"})
	w := httptest.NewRecorder()

	ObtenerSeguimientosPorDeportistaHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusOK)
	}

	var result []models.SeguimientoFisico
	json.NewDecoder(w.Body).Decode(&result)
	if len(result) != 1 {
		t.Errorf("devolvio %d elementos, esperaba 1", len(result))
	}
}

func TestActualizarSeguimientoHandler_Exitoso(t *testing.T) {
	resetHandlerTest()
	services.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})

	body := `{"deportista_id":1,"entrenador_id":1,"peso":80,"altura":1.75}`
	req := httptest.NewRequest(http.MethodPut, "/api/v1/seguimientos/1", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	req = withChiContext(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	ActualizarSeguimientoHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusOK)
	}

	var result models.SeguimientoFisico
	json.NewDecoder(w.Body).Decode(&result)
	if result.Peso != 80 {
		t.Errorf("Peso = %v; want 80", result.Peso)
	}
}

func TestEliminarSeguimientoHandler_Exitoso(t *testing.T) {
	resetHandlerTest()
	services.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/seguimientos/1", nil)
	req = withChiContext(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	EliminarSeguimientoHandler(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusNoContent)
	}
}

func TestEliminarSeguimientoHandler_NoEncontrado(t *testing.T) {
	resetHandlerTest()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/seguimientos/999", nil)
	req = withChiContext(req, map[string]string{"id": "999"})
	w := httptest.NewRecorder()

	EliminarSeguimientoHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status code = %d; want %d", w.Code, http.StatusNotFound)
	}
}
