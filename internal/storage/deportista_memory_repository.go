package storage

import "LigaFit-AWII2026/internal/models"

type DeportistaMemoryRepository struct{}

func NewDeportistaMemoryRepository() *DeportistaMemoryRepository {
	return &DeportistaMemoryRepository{}
}

func (r *DeportistaMemoryRepository) ListarDeportistas() []models.Deportista {
	return Deportistas
}

func (r *DeportistaMemoryRepository) BuscarDeportistaPorID(id int) (models.Deportista, bool) {
	for _, deportista := range Deportistas {
		if deportista.ID == id {
			return deportista, true
		}
	}

	return models.Deportista{}, false
}

func (r *DeportistaMemoryRepository) CrearDeportista(d models.Deportista) models.Deportista {
	d.ID = DeportistaIDCounter
	DeportistaIDCounter++

	Deportistas = append(Deportistas, d)
	return d
}

func (r *DeportistaMemoryRepository) ActualizarDeportista(id int, datos models.Deportista) (models.Deportista, bool) {
	for i, deportista := range Deportistas {
		if deportista.ID == id {
			datos.ID = id
			datos.CreatedAt = deportista.CreatedAt
			Deportistas[i] = datos
			return datos, true
		}
	}

	return models.Deportista{}, false
}

func (r *DeportistaMemoryRepository) BorrarDeportista(id int) bool {
	for i, deportista := range Deportistas {
		if deportista.ID == id {
			Deportistas = append(Deportistas[:i], Deportistas[i+1:]...)
			return true
		}
	}

	return false
}
