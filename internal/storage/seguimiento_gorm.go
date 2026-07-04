package storage

import (
	"LigaFit-AWII2026/internal/models"

	"gorm.io/gorm"
)

type SeguimientoGORM struct {
	db *gorm.DB
}

func NewSeguimientoGORM(db *gorm.DB) *SeguimientoGORM {
	return &SeguimientoGORM{db: db}
}

func (r *SeguimientoGORM) ListarSeguimientos() []models.SeguimientoFisico {
	var seguimientos []models.SeguimientoFisico
	r.db.Find(&seguimientos)
	return seguimientos
}

func (r *SeguimientoGORM) BuscarSeguimientoPorID(id int) (models.SeguimientoFisico, bool) {
	var seguimiento models.SeguimientoFisico

	if err := r.db.First(&seguimiento, id).Error; err != nil {
		return models.SeguimientoFisico{}, false
	}

	return seguimiento, true
}

func (r *SeguimientoGORM) CrearSeguimiento(s models.SeguimientoFisico) models.SeguimientoFisico {
	r.db.Create(&s)
	return s
}

func (r *SeguimientoGORM) ActualizarSeguimiento(id int, datos models.SeguimientoFisico) (models.SeguimientoFisico, bool) {
	var seguimiento models.SeguimientoFisico

	if err := r.db.First(&seguimiento, id).Error; err != nil {
		return models.SeguimientoFisico{}, false
	}

	datos.ID = id
	datos.CreatedAt = seguimiento.CreatedAt

	r.db.Save(&datos)

	return datos, true
}

func (r *SeguimientoGORM) BorrarSeguimiento(id int) bool {
	var seguimiento models.SeguimientoFisico

	if err := r.db.First(&seguimiento, id).Error; err != nil {
		return false
	}

	r.db.Delete(&seguimiento)
	return true
}
