package models

import "time"

type ResultadoCompetencia struct {
	ID              int       `json:"id"`
	CompetenciaID   int       `json:"competencia_id"`
	ParticipacionID int       `json:"participacion_id"`
	DeportistaID    int       `json:"deportista_id"`
	Posicion        int       `json:"posicion"`
	Puntaje         float64   `json:"puntaje"`
	Observacion     string    `json:"observacion"`
	FechaRegistro   string    `json:"fecha_registro"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
