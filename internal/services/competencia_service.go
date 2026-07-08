package services

import (
	"errors"
	"strings"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

func CrearCompetencia(competencia models.Competencia) (models.Competencia, error) {
	if strings.TrimSpace(competencia.Nombre) == "" {
		return competencia, errors.New("el nombre de la competencia es obligatorio")
	}

	if competencia.DeporteID <= 0 {
		return competencia, errors.New("el deporte_id es obligatorio")
	}

	if competencia.CategoriaID <= 0 {
		return competencia, errors.New("el categoria_id es obligatorio")
	}

	if strings.TrimSpace(competencia.Lugar) == "" {
		return competencia, errors.New("el lugar es obligatorio")
	}

	competencia.ID = storage.CompetenciaIDCounter
	storage.CompetenciaIDCounter++

	if competencia.Estado == "" {
		competencia.Estado = "Programada"
	}

	competencia.CreatedAt = time.Now()
	competencia.UpdatedAt = time.Now()

	storage.Competencias = append(storage.Competencias, competencia)

	return competencia, nil
}

func ObtenerCompetencias() []models.Competencia {
	return storage.Competencias
}

func ObtenerCompetenciaPorID(id int) (models.Competencia, bool) {
	for _, competencia := range storage.Competencias {
		if competencia.ID == id {
			return competencia, true
		}
	}

	return models.Competencia{}, false
}

func ActualizarCompetencia(id int, datos models.Competencia) (models.Competencia, error) {
	for i, competencia := range storage.Competencias {
		if competencia.ID == id {
			if strings.TrimSpace(datos.Nombre) == "" {
				return competencia, errors.New("el nombre de la competencia es obligatorio")
			}

			if datos.DeporteID <= 0 {
				return competencia, errors.New("el deporte_id es obligatorio")
			}

			if datos.CategoriaID <= 0 {
				return competencia, errors.New("el categoria_id es obligatorio")
			}

			if strings.TrimSpace(datos.Lugar) == "" {
				return competencia, errors.New("el lugar es obligatorio")
			}

			datos.ID = id
			datos.CreatedAt = competencia.CreatedAt
			datos.UpdatedAt = time.Now()

			if datos.Estado == "" {
				datos.Estado = competencia.Estado
			}

			storage.Competencias[i] = datos
			return datos, nil
		}
	}

	return models.Competencia{}, errors.New("competencia no encontrada")
}

func EliminarCompetencia(id int) bool {
	for i, competencia := range storage.Competencias {
		if competencia.ID == id {
			storage.Competencias = append(storage.Competencias[:i], storage.Competencias[i+1:]...)
			return true
		}
	}

	return false
}
