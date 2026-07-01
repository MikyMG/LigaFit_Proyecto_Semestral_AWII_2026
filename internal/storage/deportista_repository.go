package storage

import "LigaFit-AWII2026/internal/models"

type DeportistaRepository interface {
	ListarDeportistas() []models.Deportista
	BuscarDeportistaPorID(id int) (models.Deportista, bool)
	CrearDeportista(d models.Deportista) models.Deportista
	ActualizarDeportista(id int, datos models.Deportista) (models.Deportista, bool)
	BorrarDeportista(id int) bool
}
