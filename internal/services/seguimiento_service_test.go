package services

import (
	"testing"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type seguimientoRepoMock struct {
	mock.Mock
}

func (m *seguimientoRepoMock) ListarSeguimientos() []models.SeguimientoFisico {
	args := m.Called()
	return args.Get(0).([]models.SeguimientoFisico)
}

func (m *seguimientoRepoMock) BuscarSeguimientoPorID(id int) (models.SeguimientoFisico, bool) {
	args := m.Called(id)
	return args.Get(0).(models.SeguimientoFisico), args.Bool(1)
}

func (m *seguimientoRepoMock) CrearSeguimiento(s models.SeguimientoFisico) models.SeguimientoFisico {
	args := m.Called(s)
	return args.Get(0).(models.SeguimientoFisico)
}

func (m *seguimientoRepoMock) ActualizarSeguimiento(id int, datos models.SeguimientoFisico) (models.SeguimientoFisico, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.SeguimientoFisico), args.Bool(1)
}

func (m *seguimientoRepoMock) BorrarSeguimiento(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func TestCrearSeguimiento_EntrenadorIDInvalido_NoLlamaRepositorio(t *testing.T) {
	repo := new(seguimientoRepoMock)

	SetSeguimientoRepository(repo)
	defer SetSeguimientoRepository(storage.NewSeguimientoMemoryRepository())

	seguimiento := models.SeguimientoFisico{
		DeportistaID: 1,
		EntrenadorID: 0,
		Peso:         70,
		Altura:       1.75,
	}

	_, err := CrearSeguimiento(seguimiento)

	require.Error(t, err)
	require.EqualError(t, err, "el entrenador_id es obligatorio")

	repo.AssertNotCalled(t, "CrearSeguimiento", mock.Anything)
}
