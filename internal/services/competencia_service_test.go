package services

import (
	"testing"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type competenciaRepoMock struct {
	mock.Mock
}

func (m *competenciaRepoMock) ListarCompetencias() []models.Competencia {
	args := m.Called()
	return args.Get(0).([]models.Competencia)
}

func (m *competenciaRepoMock) BuscarCompetenciaPorID(id int) (models.Competencia, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Competencia), args.Bool(1)
}

func (m *competenciaRepoMock) CrearCompetencia(c models.Competencia) models.Competencia {
	args := m.Called(c)
	return args.Get(0).(models.Competencia)
}

func (m *competenciaRepoMock) ActualizarCompetencia(id int, datos models.Competencia) (models.Competencia, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Competencia), args.Bool(1)
}

func (m *competenciaRepoMock) BorrarCompetencia(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func TestCrearCompetencia_NombreVacio_NoLlamaRepositorio(t *testing.T) {
	repo := new(competenciaRepoMock)

	SetCompetenciaRepository(repo)
	defer SetCompetenciaRepository(storage.NewCompetenciaMemoryRepository())

	competencia := models.Competencia{
		Nombre:      "",
		DeporteID:   1,
		CategoriaID: 1,
		Lugar:       "Estadio ULEAM",
	}

	_, err := CrearCompetencia(competencia)

	require.Error(t, err)
	require.EqualError(t, err, "el nombre de la competencia es obligatorio")

	repo.AssertNotCalled(t, "CrearCompetencia", mock.Anything)
}
