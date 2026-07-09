package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"LigaFit-AWII2026/internal/database"
	"LigaFit-AWII2026/internal/routes"
)

func main() {
	database.ConnectDatabase()

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("LigaFit-AWII2026 funcionando"))
	})

	routes.RegisterRoutes(r)

	fmt.Println("Servidor corriendo en puerto 8080")
	http.ListenAndServe(":8080", r)
}
