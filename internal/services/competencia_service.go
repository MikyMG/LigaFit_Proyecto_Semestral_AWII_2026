package services

import (
	"errors"
	"strings"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

var competenciaRepo storage.CompetenciaRepository = storage.NewCompetenciaMemoryRepository()

func SetCompetenciaRepository(repo storage.CompetenciaRepository) {
	competenciaRepo = repo
}

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

	if competencia.Estado == "" {
		competencia.Estado = "Programada"
	}

	competencia.CreatedAt = time.Now()
	competencia.UpdatedAt = time.Now()

	competencia = competenciaRepo.CrearCompetencia(competencia)

	return competencia, nil
}

func ObtenerCompetencias() []models.Competencia {
	return competenciaRepo.ListarCompetencias()
}

func ObtenerCompetenciaPorID(id int) (models.Competencia, bool) {
	return competenciaRepo.BuscarCompetenciaPorID(id)
}

func ActualizarCompetencia(id int, datos models.Competencia) (models.Competencia, error) {
	competenciaAnterior, encontrada := competenciaRepo.BuscarCompetenciaPorID(id)
	if !encontrada {
		return models.Competencia{}, errors.New("competencia no encontrada")
	}

	if strings.TrimSpace(datos.Nombre) == "" {
		return competenciaAnterior, errors.New("el nombre de la competencia es obligatorio")
	}

	if datos.DeporteID <= 0 {
		return competenciaAnterior, errors.New("el deporte_id es obligatorio")
	}

	if datos.CategoriaID <= 0 {
		return competenciaAnterior, errors.New("el categoria_id es obligatorio")
	}

	if strings.TrimSpace(datos.Lugar) == "" {
		return competenciaAnterior, errors.New("el lugar es obligatorio")
	}

	datos.ID = id
	datos.CreatedAt = competenciaAnterior.CreatedAt
	datos.UpdatedAt = time.Now()

	if datos.Estado == "" {
		datos.Estado = competenciaAnterior.Estado
	}

	actualizada, ok := competenciaRepo.ActualizarCompetencia(id, datos)
	if !ok {
		return models.Competencia{}, errors.New("competencia no encontrada")
	}

	return actualizada, nil
}

func EliminarCompetencia(id int) bool {
	return competenciaRepo.BorrarCompetencia(id)
}
