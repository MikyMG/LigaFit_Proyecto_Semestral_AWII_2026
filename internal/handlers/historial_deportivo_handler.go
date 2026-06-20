package handlers

import (
	"encoding/json"
	"net/http"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"
)

func CrearHistorialDeportivoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var historial models.HistorialDeportivo

	if err := json.NewDecoder(r.Body).Decode(&historial); err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	nuevoHistorial, err := services.CrearHistorialDeportivo(historial)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoHistorial)
}

func ObtenerHistorialesDeportivosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.ObtenerHistorialesDeportivos())
}
