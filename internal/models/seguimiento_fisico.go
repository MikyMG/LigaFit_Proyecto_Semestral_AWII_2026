package models

import "time"

type SeguimientoFisico struct {
	ID                            int       `json:"id"`
	DeportistaID                  int       `json:"deportista_id"`
	EntrenadorID                  int       `json:"entrenador_id"`
	Peso                          float64   `json:"peso"`
	Altura                        float64   `json:"altura"`
	IMC                           float64   `json:"imc"`
	EstadoFisico                  string    `json:"estado_fisico"`
	RequiereEvaluacionNutricional bool      `json:"requiere_evaluacion_nutricional"`
	FechaRegistro                 string    `json:"fecha_registro"`
	CreatedAt                     time.Time `json:"created_at"`
	UpdatedAt                     time.Time `json:"updated_at"`
}
