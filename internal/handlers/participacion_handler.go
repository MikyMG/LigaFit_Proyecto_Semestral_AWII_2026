package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"
)

func CrearParticipacionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var participacion models.Participacion

	if err := json.NewDecoder(r.Body).Decode(&participacion); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	nuevaParticipacion, err := services.CrearParticipacion(participacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevaParticipacion)
}

func ObtenerParticipacionesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.ObtenerParticipaciones())
}

func ObtenerParticipacionPorIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	participacion, encontrado := services.ObtenerParticipacionPorID(id)
	if !encontrado {
		http.Error(w, "Participacion no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(participacion)
}

func ObtenerParticipantesPorCompetenciaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	competenciaID, err := strconv.Atoi(chi.URLParam(r, "competencia_id"))
	if err != nil {
		http.Error(w, "ID de competencia invalido", http.StatusBadRequest)
		return
	}

	participantes := services.ObtenerParticipantesPorCompetencia(competenciaID)
	json.NewEncoder(w).Encode(participantes)
}

func ActualizarParticipacionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var participacion models.Participacion

	if err := json.NewDecoder(r.Body).Decode(&participacion); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	actualizada, err := services.ActualizarParticipacion(id, participacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(actualizada)
}

func EliminarParticipacionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	if !services.EliminarParticipacion(id) {
		http.Error(w, "Participacion no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
