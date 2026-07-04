package database

import (
	"fmt"
	"log"

	"LigaFit-AWII2026/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQLite(rutaDB string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con SQLite: %w", err)
	}

	if err := migrarModelos(db); err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con PostgreSQL: %w", err)
	}

	if err := migrarModelos(db); err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectSQLiteMemory() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("error al conectar con SQLite en memoria:", err)
	}

	if err := migrarModelos(db); err != nil {
		log.Fatal("error al migrar modelos en memoria:", err)
	}

	return db
}

func migrarModelos(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.SeguimientoFisico{},
		&models.Usuario{},
	); err != nil {
		return fmt.Errorf("error al migrar modelos: %w", err)
	}

	return nil
}
