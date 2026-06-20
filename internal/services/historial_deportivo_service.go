package services

import (
	"errors"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

func CrearHistorialDeportivo(historial models.HistorialDeportivo) (models.HistorialDeportivo, error) {
	if historial.DeportistaID <= 0 {
		return historial, errors.New("el deportista_id es obligatorio")
	}

	historial.ID = storage.HistorialDeportivoIDCounter
	storage.HistorialDeportivoIDCounter++

	historial.CreatedAt = time.Now()
	historial.UpdatedAt = time.Now()

	storage.HistorialesDeportivos = append(storage.HistorialesDeportivos, historial)

	return historial, nil
}

func ObtenerHistorialesDeportivos() []models.HistorialDeportivo {
	return storage.HistorialesDeportivos
}
