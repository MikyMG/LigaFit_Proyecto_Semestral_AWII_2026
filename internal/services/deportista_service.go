package services

import (
	"errors"
	"strings"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

func asignarCategoriaPorEdad(edad int) string {
	if edad >= 6 && edad <= 12 {
		return "Infantil"
	}

	if edad >= 13 && edad <= 17 {
		return "Juvenil"
	}

	return "Adulto"
}

func CrearDeportista(deportista models.Deportista) (models.Deportista, error) {
	if strings.TrimSpace(deportista.Nombre) == "" {
		return deportista, errors.New("el nombre es obligatorio")
	}

	if deportista.Edad <= 0 {
		return deportista, errors.New("la edad debe ser mayor a 0")
	}

	if strings.TrimSpace(deportista.Telefono) == "" {
		return deportista, errors.New("el telefono es obligatorio")
	}

	if deportista.DeporteID <= 0 {
		return deportista, errors.New("el deporte_id es obligatorio")
	}

	deportista.ID = storage.DeportistaIDCounter
	storage.DeportistaIDCounter++

	deportista.Categoria = asignarCategoriaPorEdad(deportista.Edad)
	deportista.CreatedAt = time.Now()
	deportista.UpdatedAt = time.Now()

	storage.Deportistas = append(storage.Deportistas, deportista)

	return deportista, nil
}

func ObtenerDeportistas() []models.Deportista {
	return storage.Deportistas
}

func ObtenerDeportistaPorID(id int) (models.Deportista, bool) {
	for _, deportista := range storage.Deportistas {
		if deportista.ID == id {
			return deportista, true
		}
	}

	return models.Deportista{}, false
}

func ActualizarDeportista(id int, datos models.Deportista) (models.Deportista, error) {
	for i, deportista := range storage.Deportistas {
		if deportista.ID == id {
			if strings.TrimSpace(datos.Nombre) == "" {
				return deportista, errors.New("el nombre es obligatorio")
			}

			if datos.Edad <= 0 {
				return deportista, errors.New("la edad debe ser mayor a 0")
			}

			if strings.TrimSpace(datos.Telefono) == "" {
				return deportista, errors.New("el telefono es obligatorio")
			}

			if datos.DeporteID <= 0 {
				return deportista, errors.New("el deporte_id es obligatorio")
			}

			datos.ID = id
			datos.Categoria = asignarCategoriaPorEdad(datos.Edad)
			datos.CreatedAt = deportista.CreatedAt
			datos.UpdatedAt = time.Now()

			storage.Deportistas[i] = datos
			return datos, nil
		}
	}

	return models.Deportista{}, errors.New("deportista no encontrado")
}

func EliminarDeportista(id int) bool {
	for i, deportista := range storage.Deportistas {
		if deportista.ID == id {
			storage.Deportistas = append(storage.Deportistas[:i], storage.Deportistas[i+1:]...)
			return true
		}
	}

	return false
}
