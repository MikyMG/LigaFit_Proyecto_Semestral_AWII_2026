package services

import (
	"errors"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

func CrearInscripcionDeportiva(inscripcion models.InscripcionDeportiva) (models.InscripcionDeportiva, error) {
	if inscripcion.DeportistaID <= 0 {
		return inscripcion, errors.New("el deportista_id es obligatorio")
	}

	if inscripcion.DeporteID <= 0 {
		return inscripcion, errors.New("el deporte_id es obligatorio")
	}

	if inscripcion.GrupoID <= 0 {
		return inscripcion, errors.New("el grupo_id es obligatorio")
	}

	if inscripcion.Estado == "" {
		inscripcion.Estado = "Activa"
	}

	inscripcion.ID = storage.InscripcionDeportivaIDCounter
	storage.InscripcionDeportivaIDCounter++

	if inscripcion.FechaInscripcion == "" {
		inscripcion.FechaInscripcion = time.Now().Format("2006-01-02")
	}

	inscripcion.CreatedAt = time.Now()
	inscripcion.UpdatedAt = time.Now()

	storage.InscripcionesDeportivas = append(storage.InscripcionesDeportivas, inscripcion)

	return inscripcion, nil
}

func ObtenerInscripcionesDeportivas() []models.InscripcionDeportiva {
	return storage.InscripcionesDeportivas
}
