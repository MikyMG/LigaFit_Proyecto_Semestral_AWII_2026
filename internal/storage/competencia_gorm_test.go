package storage

import (
	"testing"

	"LigaFit-AWII2026/internal/models"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCompetenciaGORM_CrearYBuscarCompetencia(t *testing.T) {

	// Crear una base SQLite en memoria
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&models.Competencia{})
	require.NoError(t, err)

	repo := NewCompetenciaGORM(db)

	competencia := models.Competencia{
		Nombre:      "Copa Provincial",
		DeporteID:   1,
		CategoriaID: 1,
		FechaInicio: "2026-08-20",
		FechaFin:    "2026-08-22",
		Lugar:       "Estadio ULEAM",
		Descripcion: "Competencia interclubes",
		Estado:      "Programada",
	}

	// Crear
	creada := repo.CrearCompetencia(competencia)

	require.NotZero(t, creada.ID)

	// Buscar
	encontrada, ok := repo.BuscarCompetenciaPorID(creada.ID)

	require.True(t, ok)
	require.Equal(t, creada.ID, encontrada.ID)
	require.Equal(t, "Copa Provincial", encontrada.Nombre)
	require.Equal(t, 1, encontrada.DeporteID)
	require.Equal(t, "Programada", encontrada.Estado)
}
