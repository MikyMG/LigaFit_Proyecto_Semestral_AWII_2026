package models

import "time"

type RecomendacionNutricional struct {
	ID              int       `json:"id"`
	SeguimientoID   int       `json:"seguimiento_id"`
	NutricionistaID int       `json:"nutricionista_id"`
	Recomendacion   string    `json:"recomendacion"`
	Observaciones   string    `json:"observaciones"`
	Fecha           string    `json:"fecha"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
