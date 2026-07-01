package storage

import "LigaFit-AWII2026/internal/models"

type CompetenciaRepository interface {
	ListarCompetencias() []models.Competencia
	BuscarCompetenciaPorID(id int) (models.Competencia, bool)
	CrearCompetencia(c models.Competencia) models.Competencia
	ActualizarCompetencia(id int, datos models.Competencia) (models.Competencia, bool)
	BorrarCompetencia(id int) bool
}