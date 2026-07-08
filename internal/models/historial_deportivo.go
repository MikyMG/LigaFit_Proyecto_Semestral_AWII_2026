package models

import "time"

type HistorialDeportivo struct {
	ID                 int       `json:"id"`
	DeportistaID       int       `json:"deportista_id"`
	AniosExperiencia   int       `json:"anios_experiencia"`
	Logros             string    `json:"logros"`
	LesionesRelevantes string    `json:"lesiones_relevantes"`
	Observaciones      string    `json:"observaciones"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
