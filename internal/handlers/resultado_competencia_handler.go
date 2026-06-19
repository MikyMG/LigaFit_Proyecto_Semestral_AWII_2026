package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"
)

func CrearResultadoCompetenciaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resultado models.ResultadoCompetencia

	if err := json.NewDecoder(r.Body).Decode(&resultado); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	nuevoResultado, err := services.CrearResultadoCompetencia(resultado)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoResultado)
}

func ObtenerResultadosCompetenciaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.ObtenerResultadosCompetencia())
}

func ObtenerResultadoCompetenciaPorIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	resultado, encontrado := services.ObtenerResultadoCompetenciaPorID(id)
	if !encontrado {
		http.Error(w, "Resultado de competencia no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(resultado)
}

func ObtenerResultadosPorCompetenciaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	competenciaID, err := strconv.Atoi(chi.URLParam(r, "competencia_id"))
	if err != nil {
		http.Error(w, "ID de competencia invalido", http.StatusBadRequest)
		return
	}

	resultados := services.ObtenerResultadosPorCompetencia(competenciaID)
	json.NewEncoder(w).Encode(resultados)
}

func ActualizarResultadoCompetenciaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var resultado models.ResultadoCompetencia

	if err := json.NewDecoder(r.Body).Decode(&resultado); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	resultadoActualizado, err := services.ActualizarResultadoCompetencia(id, resultado)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(resultadoActualizado)
}

func EliminarResultadoCompetenciaHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	if !services.EliminarResultadoCompetencia(id) {
		http.Error(w, "Resultado de competencia no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
