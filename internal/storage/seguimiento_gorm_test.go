package storage

import (
	"testing"

	"LigaFit-AWII2026/internal/models"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestSeguimientoGORM_CrearYBuscarSeguimiento(t *testing.T) {
	// Preparar: base SQLite en memoria
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

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

	// Ejecutar: crear
	creado := repo.CrearSeguimiento(seguimiento)

	// Verificar: debe tener ID
	require.NotZero(t, creado.ID)

	// Ejecutar: buscar
	encontrado, ok := repo.BuscarSeguimientoPorID(creado.ID)

	// Verificar: sí existe y conserva datos
	require.True(t, ok)
	require.Equal(t, creado.ID, encontrado.ID)
	require.Equal(t, 1, encontrado.DeportistaID)
	require.Equal(t, 1, encontrado.EntrenadorID)
	require.Equal(t, 70.0, encontrado.Peso)
}
