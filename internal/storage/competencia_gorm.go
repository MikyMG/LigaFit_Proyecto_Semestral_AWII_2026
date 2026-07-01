package storage

import (
	"LigaFit-AWII2026/internal/models"

	"gorm.io/gorm"
)

type CompetenciaGORM struct {
	db *gorm.DB
}

func NewCompetenciaGORM(db *gorm.DB) *CompetenciaGORM {
	return &CompetenciaGORM{db: db}
}

func (r *CompetenciaGORM) ListarCompetencias() []models.Competencia {
	var competencias []models.Competencia
	r.db.Find(&competencias)
	return competencias
}

func (r *CompetenciaGORM) BuscarCompetenciaPorID(id int) (models.Competencia, bool) {
	var competencia models.Competencia

	if err := r.db.First(&competencia, id).Error; err != nil {
		return models.Competencia{}, false
	}

	return competencia, true
}

func (r *CompetenciaGORM) CrearCompetencia(c models.Competencia) models.Competencia {
	r.db.Create(&c)
	return c
}

func (r *CompetenciaGORM) ActualizarCompetencia(id int, datos models.Competencia) (models.Competencia, bool) {
	var competencia models.Competencia

	if err := r.db.First(&competencia, id).Error; err != nil {
		return models.Competencia{}, false
	}

	datos.ID = id

	r.db.Save(&datos)

	return datos, true
}

func (r *CompetenciaGORM) BorrarCompetencia(id int) bool {
	if err := r.db.Delete(&models.Competencia{}, id).Error; err != nil {
		return false
	}

	return true
}
