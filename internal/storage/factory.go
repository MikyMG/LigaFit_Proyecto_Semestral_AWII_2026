package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"LigaFit-AWII2026/internal/models"
)

type Recursos struct {
	Seguimientos SeguimientoRepository
	BackendUsado string
	Cerrar       func() error
}

func Inicializar(driver, dsn, rutaDB string) (*Recursos, error) {
	gdb, err := abrirGorm(driver, dsn, rutaDB)
	if err != nil {
		return nil, err
	}

	if err := gdb.AutoMigrate(
		&models.SeguimientoFisico{},
		&models.Usuario{},
	); err != nil {
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}

	seguimientos := NewSeguimientoGORM(gdb)

	cerrar := func() error {
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
		Seguimientos: seguimientos,
		BackendUsado: "gorm",
		Cerrar:       cerrar,
	}, nil
}

func abrirGorm(driver, dsn, rutaDB string) (*gorm.DB, error) {
	switch driver {
	case "postgres":
		var gdb *gorm.DB
		var err error

		for intento := 1; intento <= 10; intento++ {
			gdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				return gdb, nil
			}

			log.Printf("PostgreSQL no esta listo (intento %d/10): %v", intento, err)
			time.Sleep(2 * time.Second)
		}

		return nil, fmt.Errorf("conectar a PostgreSQL tras reintentos: %w", err)

	default:
		gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("abrir SQLite: %w", err)
		}
		return gdb, nil
	}
}
