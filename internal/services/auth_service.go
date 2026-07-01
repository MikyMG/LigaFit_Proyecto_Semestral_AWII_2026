package services

import (
	"errors"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("ligafit-secret-key")

type Claims struct {
	Email string `json:"email"`
	Rol   string `json:"rol"`
	jwt.RegisteredClaims
}

func RegistrarUsuario(usuario models.Usuario) (models.Usuario, error) {
	for _, u := range storage.Usuarios {
		if u.Email == usuario.Email {
			return usuario, errors.New("el email ya existe")
		}
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)

	usuario.ID = storage.UsuarioIDCounter
	storage.UsuarioIDCounter++
	usuario.Password = string(hash)
	usuario.CreatedAt = time.Now()
	usuario.UpdatedAt = time.Now()

	storage.Usuarios = append(storage.Usuarios, usuario)

	return usuario, nil
}

func Login(email, password string) (string, error) {
	for _, u := range storage.Usuarios {

		if u.Email == email {

			err := bcrypt.CompareHashAndPassword(
				[]byte(u.Password),
				[]byte(password),
			)

			if err != nil {
				return "", errors.New("credenciales invalidas")
			}

			expiration := time.Now().Add(24 * time.Hour)

			claims := &Claims{
				Email: u.Email,
				Rol:   u.Rol,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(expiration),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			tokenString, err := token.SignedString(jwtKey)

			if err != nil {
				return "", err
			}

			return tokenString, nil
		}
	}

	return "", errors.New("usuario no encontrado")
}
