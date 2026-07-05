package models

import "time"

type Competencia struct {
	ID          int       `json:"id"`
	Nombre      string    `json:"nombre"`
	DeporteID   int       `json:"deporte_id"`
	CategoriaID int       `json:"categoria_id"`
	FechaInicio string    `json:"fecha_inicio"`
	FechaFin    string    `json:"fecha_fin"`
	Lugar       string    `json:"lugar"`
	Descripcion string    `json:"descripcion"`
	Estado      string    `json:"estado"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
