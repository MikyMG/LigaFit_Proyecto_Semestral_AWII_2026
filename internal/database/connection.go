package database

import (
	"log"

	"LigaFit-AWII2026/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ligafit.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("error al conectar con SQLite:", err)
	}

	err = db.AutoMigrate(&models.SeguimientoFisico{})
	if err != nil {
		log.Fatal("error al migrar modelos:", err)
	}

	return db
}

func ConnectSQLiteMemory() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("error al conectar con SQLite en memoria:", err)
	}

	err = db.AutoMigrate(&models.SeguimientoFisico{})
	if err != nil {
		log.Fatal("error al migrar modelos en memoria:", err)
	}

	return db
}
