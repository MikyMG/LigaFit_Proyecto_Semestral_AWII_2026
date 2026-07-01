package services

import (
	"errors"
	"strings"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

var deportistaRepo storage.DeportistaRepository = storage.NewDeportistaMemoryRepository()

func SetDeportistaRepository(repo storage.DeportistaRepository) {
	deportistaRepo = repo
}

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

	if len(strings.TrimSpace(deportista.Nombre)) < 3 {
		return deportista, errors.New("el nombre debe tener al menos 3 caracteres")
	}

	if deportista.Edad <= 0 {
		return deportista, errors.New("la edad debe ser mayor a 0")
	}

	if strings.TrimSpace(deportista.Telefono) == "" {
		return deportista, errors.New("el telefono es obligatorio")
	}

	if len(deportista.Telefono) < 10 {
		return deportista, errors.New("el telefono debe tener al menos 10 digitos")
	}

	if strings.TrimSpace(deportista.Genero) == "" {
		return deportista, errors.New("el genero es obligatorio")
	}

	if deportista.DeporteID <= 0 {
		return deportista, errors.New("el deporte_id es obligatorio")
	}

	deportista.Categoria = asignarCategoriaPorEdad(deportista.Edad)
	deportista.CreatedAt = time.Now()
	deportista.UpdatedAt = time.Now()

	deportista = deportistaRepo.CrearDeportista(deportista)

	return deportista, nil
}

func ObtenerDeportistas() []models.Deportista {
	return deportistaRepo.ListarDeportistas()
}

func ObtenerDeportistaPorID(id int) (models.Deportista, bool) {
	return deportistaRepo.BuscarDeportistaPorID(id)
}

func ActualizarDeportista(id int, datos models.Deportista) (models.Deportista, error) {

	deportistaAnterior, encontrado := deportistaRepo.BuscarDeportistaPorID(id)
	if !encontrado {
		return models.Deportista{}, errors.New("deportista no encontrado")
	}

	if strings.TrimSpace(datos.Nombre) == "" {
		return deportistaAnterior, errors.New("el nombre es obligatorio")
	}

	if len(strings.TrimSpace(datos.Nombre)) < 3 {
		return deportistaAnterior, errors.New("el nombre debe tener al menos 3 caracteres")
	}

	if datos.Edad <= 0 {
		return deportistaAnterior, errors.New("la edad debe ser mayor a 0")
	}

	if strings.TrimSpace(datos.Telefono) == "" {
		return deportistaAnterior, errors.New("el telefono es obligatorio")
	}

	if len(datos.Telefono) < 10 {
		return deportistaAnterior, errors.New("el telefono debe tener al menos 10 digitos")
	}

	if strings.TrimSpace(datos.Genero) == "" {
		return deportistaAnterior, errors.New("el genero es obligatorio")
	}

	if datos.DeporteID <= 0 {
		return deportistaAnterior, errors.New("el deporte_id es obligatorio")
	}

	datos.ID = id
	datos.Categoria = asignarCategoriaPorEdad(datos.Edad)
	datos.CreatedAt = deportistaAnterior.CreatedAt
	datos.UpdatedAt = time.Now()

	actualizado, ok := deportistaRepo.ActualizarDeportista(id, datos)
	if !ok {
		return models.Deportista{}, errors.New("deportista no encontrado")
	}

	return actualizado, nil
}

func EliminarDeportista(id int) bool {
	return deportistaRepo.BorrarDeportista(id)
}
