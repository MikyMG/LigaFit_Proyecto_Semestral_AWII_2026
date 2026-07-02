package handlers_test

import (
	"bytes"
	"net/http"

	"LigaFit-AWII2026/internal/handlers"
	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"

	"github.com/go-chi/chi/v5"
)

type fakeCompetenciaRepository struct {
	competencias []models.Competencia
	nextID       int
}

func (f *fakeCompetenciaRepository) ListarCompetencias() []models.Competencia {
	return f.competencias
}

func (f *fakeCompetenciaRepository) BuscarCompetenciaPorID(id int) (models.Competencia, bool) {
	for _, c := range f.competencias {
		if c.ID == id {
			return c, true
		}
	}
	return models.Competencia{}, false
}

func (f *fakeCompetenciaRepository) CrearCompetencia(c models.Competencia) models.Competencia {
	if f.nextID == 0 {
		f.nextID = 1
	}
	c.ID = f.nextID
	f.nextID++
	f.competencias = append(f.competencias, c)
	return c
}

func (f *fakeCompetenciaRepository) ActualizarCompetencia(id int, datos models.Competencia) (models.Competencia, bool) {
	for i, c := range f.competencias {
		if c.ID == id {
			datos.ID = id
			datos.CreatedAt = c.CreatedAt
			f.competencias[i] = datos
			return datos, true
		}
	}
	return models.Competencia{}, false
}

func (f *fakeCompetenciaRepository) BorrarCompetencia(id int) bool {
	for i, c := range f.competencias {
		if c.ID == id {
			f.competencias = append(f.competencias[:i], f.competencias[i+1:]...)
			return true
		}
	}
	return false
}

func nuevoRouterCompetencia(repo *fakeCompetenciaRepository) http.Handler {
	services.SetCompetenciaRepository(repo)

	r := chi.NewRouter()

	r.Route("/api/v1/competencias", func(r chi.Router) {
		r.Post("/", handlers.CrearCompetenciaHandler)
		r.Get("/", handlers.ObtenerCompetenciasHandler)
		r.Get("/{id}", handlers.ObtenerCompetenciaPorIDHandler)
		r.Put("/{id}", handlers.ActualizarCompetenciaHandler)
		r.Delete("/{id}", handlers.EliminarCompetenciaHandler)
	})

	return r
}

func jsonReq(body string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(body))
}
