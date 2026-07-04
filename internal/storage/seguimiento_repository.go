package storage

import "LigaFit-AWII2026/internal/models"

type SeguimientoRepository interface {
	ListarSeguimientos() []models.SeguimientoFisico
	BuscarSeguimientoPorID(id int) (models.SeguimientoFisico, bool)
	CrearSeguimiento(s models.SeguimientoFisico) models.SeguimientoFisico
	ActualizarSeguimiento(id int, datos models.SeguimientoFisico) (models.SeguimientoFisico, bool)
	BorrarSeguimiento(id int) bool
}
