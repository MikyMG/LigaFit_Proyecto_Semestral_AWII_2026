package services

import (
	"errors"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

func CrearResultadoCompetencia(resultado models.ResultadoCompetencia) (models.ResultadoCompetencia, error) {
	if resultado.CompetenciaID <= 0 {
		return resultado, errors.New("el competencia_id es obligatorio")
	}

	if resultado.ParticipacionID <= 0 {
		return resultado, errors.New("el participacion_id es obligatorio")
	}

	if resultado.DeportistaID <= 0 {
		return resultado, errors.New("el deportista_id es obligatorio")
	}

	if resultado.Posicion <= 0 {
		return resultado, errors.New("la posicion debe ser mayor a 0")
	}

	if resultado.Puntaje < 0 {
		return resultado, errors.New("el puntaje no puede ser negativo")
	}

	resultado.ID = storage.ResultadoCompetenciaIDCounter
	storage.ResultadoCompetenciaIDCounter++

	if resultado.FechaRegistro == "" {
		resultado.FechaRegistro = time.Now().Format("2006-01-02")
	}

	resultado.CreatedAt = time.Now()
	resultado.UpdatedAt = time.Now()

	storage.ResultadosCompetencia = append(storage.ResultadosCompetencia, resultado)

	return resultado, nil
}

func ObtenerResultadosCompetencia() []models.ResultadoCompetencia {
	return storage.ResultadosCompetencia
}

func ObtenerResultadoCompetenciaPorID(id int) (models.ResultadoCompetencia, bool) {
	for _, resultado := range storage.ResultadosCompetencia {
		if resultado.ID == id {
			return resultado, true
		}
	}

	return models.ResultadoCompetencia{}, false
}

func ObtenerResultadosPorCompetencia(competenciaID int) []models.ResultadoCompetencia {
	var resultados []models.ResultadoCompetencia

	for _, resultado := range storage.ResultadosCompetencia {
		if resultado.CompetenciaID == competenciaID {
			resultados = append(resultados, resultado)
		}
	}

	return resultados
}

func ActualizarResultadoCompetencia(id int, datos models.ResultadoCompetencia) (models.ResultadoCompetencia, error) {
	for i, resultado := range storage.ResultadosCompetencia {
		if resultado.ID == id {
			if datos.CompetenciaID <= 0 {
				return resultado, errors.New("el competencia_id es obligatorio")
			}

			if datos.ParticipacionID <= 0 {
				return resultado, errors.New("el participacion_id es obligatorio")
			}

			if datos.DeportistaID <= 0 {
				return resultado, errors.New("el deportista_id es obligatorio")
			}

			if datos.Posicion <= 0 {
				return resultado, errors.New("la posicion debe ser mayor a 0")
			}

			if datos.Puntaje < 0 {
				return resultado, errors.New("el puntaje no puede ser negativo")
			}

			datos.ID = id
			datos.CreatedAt = resultado.CreatedAt
			datos.UpdatedAt = time.Now()

			if datos.FechaRegistro == "" {
				datos.FechaRegistro = resultado.FechaRegistro
			}

			storage.ResultadosCompetencia[i] = datos
			return datos, nil
		}
	}

	return models.ResultadoCompetencia{}, errors.New("resultado de competencia no encontrado")
}

func EliminarResultadoCompetencia(id int) bool {
	for i, resultado := range storage.ResultadosCompetencia {
		if resultado.ID == id {
			storage.ResultadosCompetencia = append(storage.ResultadosCompetencia[:i], storage.ResultadosCompetencia[i+1:]...)
			return true
		}
	}

	return false
}
