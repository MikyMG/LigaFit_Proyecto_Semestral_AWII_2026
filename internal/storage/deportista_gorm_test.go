package storage

import (
	"testing"

	"LigaFit-AWII2026/internal/models"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestDeportistaGORM_CRUDCompleto(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	sqlDB, err := db.DB()
	require.NoError(t, err)
	sqlDB.SetMaxOpenConns(1)

	err = db.AutoMigrate(&models.Deportista{})
	require.NoError(t, err)

	repo := NewDeportistaGORM(db)

	deportista := models.Deportista{
		Nombre:    "Juan Perez",
		Edad:      15,
		Genero:    "Masculino",
		Telefono:  "0987654321",
		DeporteID: 1,
		Categoria: "Juvenil",
		GrupoID:   1,
	}

	creado := repo.CrearDeportista(deportista)

	require.NotZero(t, creado.ID)

	encontrado, ok := repo.BuscarDeportistaPorID(creado.ID)

	require.True(t, ok)
	require.Equal(t, creado.ID, encontrado.ID)
	require.Equal(t, "Juan Perez", encontrado.Nombre)

	actualizado := models.Deportista{
		Nombre:    "Juan Actualizado",
		Edad:      16,
		Genero:    "Masculino",
		Telefono:  "0987654321",
		DeporteID: 1,
		Categoria: "Juvenil",
		GrupoID:   1,
	}

	guardado, ok := repo.ActualizarDeportista(creado.ID, actualizado)

	require.True(t, ok)
	require.Equal(t, "Juan Actualizado", guardado.Nombre)

	eliminado := repo.BorrarDeportista(creado.ID)

	require.True(t, eliminado)

	_, ok = repo.BuscarDeportistaPorID(creado.ID)

	require.False(t, ok)
}
