package services

import (
	"errors"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

func CrearParticipacion(participacion models.Participacion) (models.Participacion, error) {
	if participacion.CompetenciaID <= 0 {
		return participacion, errors.New("el competencia_id es obligatorio")
	}

	if participacion.DeportistaID <= 0 {
		return participacion, errors.New("el deportista_id es obligatorio")
	}

	if participacion.GrupoID <= 0 {
		return participacion, errors.New("el grupo_id es obligatorio")
	}

	for _, p := range storage.Participaciones {
		if p.CompetenciaID == participacion.CompetenciaID && p.DeportistaID == participacion.DeportistaID {
			return participacion, errors.New("el deportista ya está registrado en esta competencia")
		}
	}

	participacion.ID = storage.ParticipacionIDCounter
	storage.ParticipacionIDCounter++

	if participacion.Estado == "" {
		participacion.Estado = "Inscrito"
	}

	if participacion.FechaInscripcion == "" {
		participacion.FechaInscripcion = time.Now().Format("2006-01-02")
	}

	participacion.CreatedAt = time.Now()
	participacion.UpdatedAt = time.Now()

	storage.Participaciones = append(storage.Participaciones, participacion)

	return participacion, nil
}

func ObtenerParticipaciones() []models.Participacion {
	return storage.Participaciones
}

func ObtenerParticipacionPorID(id int) (models.Participacion, bool) {
	for _, participacion := range storage.Participaciones {
		if participacion.ID == id {
			return participacion, true
		}
	}

	return models.Participacion{}, false
}

func ObtenerParticipantesPorCompetencia(competenciaID int) []models.Participacion {
	var resultado []models.Participacion

	for _, participacion := range storage.Participaciones {
		if participacion.CompetenciaID == competenciaID {
			resultado = append(resultado, participacion)
		}
	}

	return resultado
}

func ActualizarParticipacion(id int, datos models.Participacion) (models.Participacion, error) {
	for i, participacion := range storage.Participaciones {
		if participacion.ID == id {
			if datos.CompetenciaID <= 0 {
				return participacion, errors.New("el competencia_id es obligatorio")
			}

			if datos.DeportistaID <= 0 {
				return participacion, errors.New("el deportista_id es obligatorio")
			}

			if datos.GrupoID <= 0 {
				return participacion, errors.New("el grupo_id es obligatorio")
			}

			datos.ID = id
			datos.CreatedAt = participacion.CreatedAt
			datos.UpdatedAt = time.Now()

			if datos.Estado == "" {
				datos.Estado = participacion.Estado
			}

			if datos.FechaInscripcion == "" {
				datos.FechaInscripcion = participacion.FechaInscripcion
			}

			storage.Participaciones[i] = datos
			return datos, nil
		}
	}

	return models.Participacion{}, errors.New("participación no encontrada")
}

func EliminarParticipacion(id int) bool {
	for i, participacion := range storage.Participaciones {
		if participacion.ID == id {
			storage.Participaciones = append(storage.Participaciones[:i], storage.Participaciones[i+1:]...)
			return true
		}
	}

	return false
}
