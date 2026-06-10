package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"
)

func CrearDeportistaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var deportista models.Deportista

	if err := json.NewDecoder(r.Body).Decode(&deportista); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	nuevoDeportista, err := services.CrearDeportista(deportista)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoDeportista)
}

func ObtenerDeportistasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.ObtenerDeportistas())
}

func ObtenerDeportistaPorIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	deportista, encontrado := services.ObtenerDeportistaPorID(id)
	if !encontrado {
		http.Error(w, "Deportista no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(deportista)
}

func ActualizarDeportistaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var deportista models.Deportista

	if err := json.NewDecoder(r.Body).Decode(&deportista); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	deportistaActualizado, err := services.ActualizarDeportista(id, deportista)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(deportistaActualizado)
}

func EliminarDeportistaHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	if !services.EliminarDeportista(id) {
		http.Error(w, "Deportista no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
