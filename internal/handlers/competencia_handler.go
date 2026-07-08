package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"
)

func CrearCompetenciaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var competencia models.Competencia

	if err := json.NewDecoder(r.Body).Decode(&competencia); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	nuevaCompetencia, err := services.CrearCompetencia(competencia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevaCompetencia)
}

func ObtenerCompetenciasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.ObtenerCompetencias())
}

func ObtenerCompetenciaPorIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	competencia, encontrado := services.ObtenerCompetenciaPorID(id)
	if !encontrado {
		http.Error(w, "Competencia no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(competencia)
}

func ActualizarCompetenciaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var competencia models.Competencia

	if err := json.NewDecoder(r.Body).Decode(&competencia); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	actualizada, err := services.ActualizarCompetencia(id, competencia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(actualizada)
}

func EliminarCompetenciaHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	if !services.EliminarCompetencia(id) {
		http.Error(w, "Competencia no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
