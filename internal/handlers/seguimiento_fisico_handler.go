package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"
)

func CrearSeguimientoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var seguimiento models.SeguimientoFisico

	if err := json.NewDecoder(r.Body).Decode(&seguimiento); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	nuevoSeguimiento, err := services.CrearSeguimiento(seguimiento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoSeguimiento)
}

func ObtenerSeguimientosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	seguimientos := services.ObtenerSeguimientos()
	json.NewEncoder(w).Encode(seguimientos)
}

func ObtenerSeguimientoPorIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	seguimiento, encontrado := services.ObtenerSeguimientoPorID(id)
	if !encontrado {
		http.Error(w, "Seguimiento no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(seguimiento)
}

func ObtenerSeguimientosPorDeportistaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deportistaID, err := strconv.Atoi(chi.URLParam(r, "deportista_id"))
	if err != nil {
		http.Error(w, "ID de deportista invalido", http.StatusBadRequest)
		return
	}

	seguimientos := services.ObtenerSeguimientosPorDeportista(deportistaID)
	json.NewEncoder(w).Encode(seguimientos)
}

func ActualizarSeguimientoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var seguimiento models.SeguimientoFisico

	if err := json.NewDecoder(r.Body).Decode(&seguimiento); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	seguimientoActualizado, err := services.ActualizarSeguimiento(id, seguimiento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(seguimientoActualizado)
}

func EliminarSeguimientoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	eliminado := services.EliminarSeguimiento(id)
	if !eliminado {
		http.Error(w, "Seguimiento no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
