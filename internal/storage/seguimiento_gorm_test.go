package storage

import (
	"testing"

	"LigaFit-AWII2026/internal/models"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestSeguimientoGORM_CRUDCompleto(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	sqlDB, err := db.DB()
	require.NoError(t, err)
	sqlDB.SetMaxOpenConns(1)

	err = db.AutoMigrate(&models.SeguimientoFisico{})
	require.NoError(t, err)

	repo := NewSeguimientoGORM(db)

	seguimiento := models.SeguimientoFisico{
		DeportistaID: 1,
		EntrenadorID: 1,
		Peso:         70,
		Altura:       1.75,
		IMC:          22.86,
		EstadoFisico: "Normal",
	}

	creado := repo.CrearSeguimiento(seguimiento)

	require.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarSeguimientoPorID(creado.ID)

	require.True(t, ok)
	require.Equal(t, creado.ID, encontrado.ID)
	require.Equal(t, 70.0, encontrado.Peso)

	actualizado := models.SeguimientoFisico{
		DeportistaID: 1,
		EntrenadorID: 1,
		Peso:         75,
		Altura:       1.75,
		IMC:          24.49,
		EstadoFisico: "Normal",
	}

	guardado, ok := repo.ActualizarSeguimiento(creado.ID, actualizado)

	require.True(t, ok)
	require.Equal(t, 75.0, guardado.Peso)

	eliminado := repo.BorrarSeguimiento(creado.ID)

	require.True(t, eliminado)

	_, ok = repo.BuscarSeguimientoPorID(creado.ID)

	require.False(t, ok)
}
