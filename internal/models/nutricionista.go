package models

import "time"

type Nutricionista struct {
	ID           int       `json:"id"`
	Nombre       string    `json:"nombre"`
	Telefono     string    `json:"telefono"`
	Especialidad string    `json:"especialidad"`
	Estado       bool      `json:"estado"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
