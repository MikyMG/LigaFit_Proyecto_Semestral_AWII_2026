package storage

import (
	"LigaFit-AWII2026/internal/models"

	"gorm.io/gorm"
)

type DeportistaGORM struct {
	db *gorm.DB
}

func NewDeportistaGORM(db *gorm.DB) *DeportistaGORM {
	return &DeportistaGORM{db: db}
}

func (r *DeportistaGORM) ListarDeportistas() []models.Deportista {
	var deportistas []models.Deportista
	r.db.Find(&deportistas)
	return deportistas
}

func (r *DeportistaGORM) BuscarDeportistaPorID(id int) (models.Deportista, bool) {
	var deportista models.Deportista

	if err := r.db.First(&deportista, id).Error; err != nil {
		return models.Deportista{}, false
	}

	return deportista, true
}

func (r *DeportistaGORM) CrearDeportista(d models.Deportista) models.Deportista {
	r.db.Create(&d)
	return d
}

func (r *DeportistaGORM) ActualizarDeportista(id int, datos models.Deportista) (models.Deportista, bool) {
	var deportista models.Deportista

	if err := r.db.First(&deportista, id).Error; err != nil {
		return models.Deportista{}, false
	}

	datos.ID = id

	r.db.Save(&datos)

	return datos, true
}

func (r *DeportistaGORM) BorrarDeportista(id int) bool {
	if err := r.db.Delete(&models.Deportista{}, id).Error; err != nil {
		return false
	}

	return true
}
