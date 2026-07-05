package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Puerto       string
	DBDriver     string
	DBDsn        string
	RutaDB       string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Cargar() Config {
	_ = godotenv.Load()

	return Config{
		Puerto:       conTexto("PUERTO", ":8080"),
		DBDriver:     conTexto("DB_DRIVER", "sqlite"),
		DBDsn:        conTexto("DB_DSN", ""),
		RutaDB:       conTexto("RUTA_DB", "ligafit.db"),
		ReadTimeout:  conDuracion("HTTP_READ_TIMEOUT", 10*time.Second),
		WriteTimeout: conDuracion("HTTP_WRITE_TIMEOUT", 10*time.Second),
	}
}

func conTexto(clave, porDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}

func conDuracion(clave string, porDefecto time.Duration) time.Duration {
	v := os.Getenv(clave)
	if v == "" {
		return porDefecto
	}

	d, err := time.ParseDuration(v)
	if err != nil {
		return porDefecto
	}

	return d
}
