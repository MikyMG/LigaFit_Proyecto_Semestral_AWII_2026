package services

import (
	"testing"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type deportistaRepoMock struct {
	mock.Mock
}

func (m *deportistaRepoMock) ListarDeportistas() []models.Deportista {
	args := m.Called()
	return args.Get(0).([]models.Deportista)
}

func (m *deportistaRepoMock) BuscarDeportistaPorID(id int) (models.Deportista, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Deportista), args.Bool(1)
}

func (m *deportistaRepoMock) CrearDeportista(d models.Deportista) models.Deportista {
	args := m.Called(d)
	return args.Get(0).(models.Deportista)
}

func (m *deportistaRepoMock) ActualizarDeportista(id int, datos models.Deportista) (models.Deportista, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Deportista), args.Bool(1)
}

func (m *deportistaRepoMock) BorrarDeportista(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func TestCrearDeportista_NombreCorto_NoLlamaRepositorio(t *testing.T) {
	repo := new(deportistaRepoMock)

	SetDeportistaRepository(repo)
	defer SetDeportistaRepository(storage.NewDeportistaMemoryRepository())

	deportista := models.Deportista{
		Nombre:    "Jo",
		Edad:      15,
		Genero:    "Masculino",
		Telefono:  "0987654321",
		DeporteID: 1,
		GrupoID:   1,
	}

	_, err := CrearDeportista(deportista)

	require.Error(t, err)
	require.EqualError(t, err, "el nombre debe tener al menos 3 caracteres")

	repo.AssertNotCalled(t, "CrearDeportista", mock.Anything)
}
