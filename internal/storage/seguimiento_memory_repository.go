package storage

import "LigaFit-AWII2026/internal/models"

type SeguimientoMemoryRepository struct{}

func NewSeguimientoMemoryRepository() *SeguimientoMemoryRepository {
	return &SeguimientoMemoryRepository{}
}

func (r *SeguimientoMemoryRepository) ListarSeguimientos() []models.SeguimientoFisico {
	return Seguimientos
}

func (r *SeguimientoMemoryRepository) BuscarSeguimientoPorID(id int) (models.SeguimientoFisico, bool) {
	for _, seguimiento := range Seguimientos {
		if seguimiento.ID == id {
			return seguimiento, true
		}
	}
	return models.SeguimientoFisico{}, false
}

func (r *SeguimientoMemoryRepository) CrearSeguimiento(s models.SeguimientoFisico) models.SeguimientoFisico {
	s.ID = SeguimientoIDCounter
	SeguimientoIDCounter++

	Seguimientos = append(Seguimientos, s)
	return s
}

func (r *SeguimientoMemoryRepository) ActualizarSeguimiento(id int, datos models.SeguimientoFisico) (models.SeguimientoFisico, bool) {
	for i, seguimiento := range Seguimientos {
		if seguimiento.ID == id {
			datos.ID = id
			datos.CreatedAt = seguimiento.CreatedAt
			Seguimientos[i] = datos
			return datos, true
		}
	}
	return models.SeguimientoFisico{}, false
}

func (r *SeguimientoMemoryRepository) BorrarSeguimiento(id int) bool {
	for i, seguimiento := range Seguimientos {
		if seguimiento.ID == id {
			Seguimientos = append(Seguimientos[:i], Seguimientos[i+1:]...)
			return true
		}
	}
	return false
}
