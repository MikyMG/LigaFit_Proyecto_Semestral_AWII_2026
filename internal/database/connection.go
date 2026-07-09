package database

import (
	"log"
	"os"

	"LigaFit-AWII2026/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	driver := os.Getenv("DB_DRIVER")

	if driver == "postgres" {
		db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
		if err != nil {
			log.Fatal("error al conectar con PostgreSQL:", err)
		}

		err = db.AutoMigrate(&models.Deportista{})
		if err != nil {
			log.Fatal("error al migrar modelos en PostgreSQL:", err)
		}

		log.Println("Base de datos PostgreSQL conectada")
		return db
	}

	db, err := gorm.Open(sqlite.Open("ligafit.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("error al conectar con SQLite:", err)
	}

	err = db.AutoMigrate(&models.Deportista{})
	if err != nil {
		log.Fatal("error al migrar modelos en SQLite:", err)
	}

	log.Println("Base de datos SQLite conectada")
	return db
}
