package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/config"
	"LigaFit-AWII2026/internal/routes"
	"LigaFit-AWII2026/internal/services"
	"LigaFit-AWII2026/internal/storage"
)

func main() {
	cfg := config.Cargar()

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()

	// Aquí conectamos el repositorio GORM creado por el factory con el Service.
	// Así el módulo 2 ya no usa memoria, sino el repositorio que viene del factory.
	services.SetSeguimientoRepository(recursos.Seguimientos)

	log.Printf("Motor de base de datos: %s | Backend usado: %s", cfg.DBDriver, recursos.BackendUsado)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("LigaFit-AWII2026 funcionando"))
	})

	routes.RegisterRoutes(r)

	srv := &http.Server{
		Addr:         cfg.Puerto,
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errServidor := make(chan error, 1)

	go func() {
		log.Printf("Servidor escuchando en http://localhost%s", cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errServidor <- err
		}
	}()

	select {
	case err := <-errServidor:
		return err
	case <-ctx.Done():
		log.Println("Senal de apagado recibida, cerrando ordenadamente...")
	}

	ctxApagado, cancelar := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelar()

	if err := srv.Shutdown(ctxApagado); err != nil {
		return err
	}

	log.Println("Servidor detenido limpiamente.")
	return nil
}
