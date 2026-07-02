package storage

import (
	"testing"

	"LigaFit-AWII2026/internal/models"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCompetenciaGORM_CRUDCompleto(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	sqlDB, err := db.DB()
	require.NoError(t, err)
	sqlDB.SetMaxOpenConns(1)

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

	creada := repo.CrearCompetencia(competencia)

	require.NotZero(t, creada.ID)

	encontrada, ok := repo.BuscarCompetenciaPorID(creada.ID)

	require.True(t, ok)
	require.Equal(t, creada.ID, encontrada.ID)
	require.Equal(t, "Copa Provincial", encontrada.Nombre)

	actualizada := models.Competencia{
		Nombre:      "Copa Actualizada",
		DeporteID:   1,
		CategoriaID: 1,
		FechaInicio: "2026-08-25",
		FechaFin:    "2026-08-27",
		Lugar:       "Coliseo ULEAM",
		Descripcion: "Competencia actualizada",
		Estado:      "Finalizada",
	}

	guardada, ok := repo.ActualizarCompetencia(creada.ID, actualizada)

	require.True(t, ok)
	require.Equal(t, "Copa Actualizada", guardada.Nombre)
	require.Equal(t, "Finalizada", guardada.Estado)

	eliminada := repo.BorrarCompetencia(creada.ID)

	require.True(t, eliminada)

	_, ok = repo.BuscarCompetenciaPorID(creada.ID)

	require.False(t, ok)
}
