package handlers_test

import (
	"bytes"
	"net/http"

	"LigaFit-AWII2026/internal/handlers"
	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"

	"github.com/go-chi/chi/v5"
)

type fakeSeguimientoRepository struct {
	seguimientos []models.SeguimientoFisico
	nextID       int
}

func (f *fakeSeguimientoRepository) ListarSeguimientos() []models.SeguimientoFisico {
	return f.seguimientos
}

func (f *fakeSeguimientoRepository) BuscarSeguimientoPorID(id int) (models.SeguimientoFisico, bool) {
	for _, s := range f.seguimientos {
		if s.ID == id {
			return s, true
		}
	}
	return models.SeguimientoFisico{}, false
}

func (f *fakeSeguimientoRepository) CrearSeguimiento(s models.SeguimientoFisico) models.SeguimientoFisico {
	if f.nextID == 0 {
		f.nextID = 1
	}
	s.ID = f.nextID
	f.nextID++
	f.seguimientos = append(f.seguimientos, s)
	return s
}

func (f *fakeSeguimientoRepository) ActualizarSeguimiento(id int, datos models.SeguimientoFisico) (models.SeguimientoFisico, bool) {
	for i, s := range f.seguimientos {
		if s.ID == id {
			datos.ID = id
			datos.CreatedAt = s.CreatedAt
			f.seguimientos[i] = datos
			return datos, true
		}
	}
	return models.SeguimientoFisico{}, false
}

func (f *fakeSeguimientoRepository) BorrarSeguimiento(id int) bool {
	for i, s := range f.seguimientos {
		if s.ID == id {
			f.seguimientos = append(f.seguimientos[:i], f.seguimientos[i+1:]...)
			return true
		}
	}
	return false
}

func nuevoRouterSeguimiento(repo *fakeSeguimientoRepository) http.Handler {
	services.SetSeguimientoRepository(repo)

	r := chi.NewRouter()

	r.Route("/api/v1/seguimientos", func(r chi.Router) {
		r.Post("/", handlers.CrearSeguimientoHandler)
		r.Get("/", handlers.ObtenerSeguimientosHandler)
		r.Get("/{id}", handlers.ObtenerSeguimientoPorIDHandler)
		r.Put("/{id}", handlers.ActualizarSeguimientoHandler)
		r.Delete("/{id}", handlers.EliminarSeguimientoHandler)
	})

	return r
}

func jsonReq(body string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(body))
}
