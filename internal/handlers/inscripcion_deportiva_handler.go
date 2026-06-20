package handlers

import (
	"encoding/json"
	"net/http"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"
)

func CrearInscripcionDeportivaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inscripcion models.InscripcionDeportiva

	if err := json.NewDecoder(r.Body).Decode(&inscripcion); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	nuevaInscripcion, err := services.CrearInscripcionDeportiva(inscripcion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevaInscripcion)
}

func ObtenerInscripcionesDeportivasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.ObtenerInscripcionesDeportivas())
}
