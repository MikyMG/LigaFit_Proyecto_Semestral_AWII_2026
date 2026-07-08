package models

import "time"

type Deportista struct {
	ID        int       `json:"id"`
	Nombre    string    `json:"nombre"`
	Edad      int       `json:"edad"`
	Genero    string    `json:"genero"`
	Telefono  string    `json:"telefono"`
	DeporteID int       `json:"deporte_id"`
	Categoria string    `json:"categoria"`
	GrupoID   int       `json:"grupo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
