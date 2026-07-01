package storage

import (
	"testing"

	"LigaFit-AWII2026/internal/models"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestDeportistaGORM_CrearYBuscarDeportista(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

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
	require.Equal(t, "Juvenil", encontrado.Categoria)
}
