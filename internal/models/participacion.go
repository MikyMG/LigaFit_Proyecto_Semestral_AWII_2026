package models

import "time"

type Participacion struct {
	ID               int       `json:"id"`
	CompetenciaID    int       `json:"competencia_id"`
	DeportistaID     int       `json:"deportista_id"`
	GrupoID          int       `json:"grupo_id"`
	FechaInscripcion string    `json:"fecha_inscripcion"`
	Estado           string    `json:"estado"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
