package handlers

import (
	"encoding/json"
	"net/http"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/services"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "datos invalidos", http.StatusBadRequest)
		return
	}

	nuevoUsuario, err := services.RegistrarUsuario(usuario)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoUsuario)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "datos invalidos", http.StatusBadRequest)
		return
	}

	token, err := services.Login(login.Email, login.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
