package handlers_test

import (
	"bytes"
	"net/http"

	"LigaFit-AWII2026/internal/handlers"
	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"

	"github.com/go-chi/chi/v5"
)

type fakeDeportistaRepository struct {
	deportistas []models.Deportista
	nextID      int
}

func (f *fakeDeportistaRepository) ListarDeportistas() []models.Deportista {
	return f.deportistas
}

func (f *fakeDeportistaRepository) BuscarDeportistaPorID(id int) (models.Deportista, bool) {
	for _, d := range f.deportistas {
		if d.ID == id {
			return d, true
		}
	}
	return models.Deportista{}, false
}

func (f *fakeDeportistaRepository) CrearDeportista(d models.Deportista) models.Deportista {
	if f.nextID == 0 {
		f.nextID = 1
	}
	d.ID = f.nextID
	f.nextID++
	f.deportistas = append(f.deportistas, d)
	return d
}

func (f *fakeDeportistaRepository) ActualizarDeportista(id int, datos models.Deportista) (models.Deportista, bool) {
	for i, d := range f.deportistas {
		if d.ID == id {
			datos.ID = id
			datos.CreatedAt = d.CreatedAt
			f.deportistas[i] = datos
			return datos, true
		}
	}
	return models.Deportista{}, false
}

func (f *fakeDeportistaRepository) BorrarDeportista(id int) bool {
	for i, d := range f.deportistas {
		if d.ID == id {
			f.deportistas = append(f.deportistas[:i], f.deportistas[i+1:]...)
			return true
		}
	}
	return false
}

func nuevoRouterDeportista(repo *fakeDeportistaRepository) http.Handler {
	services.SetDeportistaRepository(repo)

	r := chi.NewRouter()

	r.Route("/api/v1/deportistas", func(r chi.Router) {
		r.Post("/", handlers.CrearDeportistaHandler)
		r.Get("/", handlers.ObtenerDeportistasHandler)
		r.Get("/{id}", handlers.ObtenerDeportistaPorIDHandler)
		r.Put("/{id}", handlers.ActualizarDeportistaHandler)
		r.Delete("/{id}", handlers.EliminarDeportistaHandler)
	})

	return r
}

func jsonReq(body string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(body))
}
