package storage

import "LigaFit-AWII2026/internal/models"

type CompetenciaMemoryRepository struct{}

func NewCompetenciaMemoryRepository() *CompetenciaMemoryRepository {
	return &CompetenciaMemoryRepository{}
}

func (r *CompetenciaMemoryRepository) ListarCompetencias() []models.Competencia {
	return Competencias
}

func (r *CompetenciaMemoryRepository) BuscarCompetenciaPorID(id int) (models.Competencia, bool) {
	for _, competencia := range Competencias {
		if competencia.ID == id {
			return competencia, true
		}
	}

	return models.Competencia{}, false
}

func (r *CompetenciaMemoryRepository) CrearCompetencia(c models.Competencia) models.Competencia {
	c.ID = CompetenciaIDCounter
	CompetenciaIDCounter++

	Competencias = append(Competencias, c)
	return c
}

func (r *CompetenciaMemoryRepository) ActualizarCompetencia(id int, datos models.Competencia) (models.Competencia, bool) {
	for i, competencia := range Competencias {
		if competencia.ID == id {
			datos.ID = id
			datos.CreatedAt = competencia.CreatedAt
			Competencias[i] = datos
			return datos, true
		}
	}

	return models.Competencia{}, false
}

func (r *CompetenciaMemoryRepository) BorrarCompetencia(id int) bool {
	for i, competencia := range Competencias {
		if competencia.ID == id {
			Competencias = append(Competencias[:i], Competencias[i+1:]...)
			return true
		}
	}

	return false
}
