package services

import (
	"errors"
	"math"
	"time"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

func calcularIMC(peso float64, altura float64) float64 {
	imc := peso / (altura * altura)
	return math.Round(imc*100) / 100
}

func clasificarEstadoFisico(imc float64) string {
	if imc < 18.5 {
		return "Bajo peso"
	}

	if imc >= 18.5 && imc <= 24.9 {
		return "Normal"
	}

	if imc >= 25 && imc <= 29.9 {
		return "Sobrepeso"
	}

	return "Obesidad"
}

func requiereEvaluacionNutricional(estado string) bool {
	return estado == "Bajo peso" || estado == "Obesidad"
}

func CrearSeguimiento(seguimiento models.SeguimientoFisico) (models.SeguimientoFisico, error) {
	if seguimiento.DeportistaID <= 0 {
		return seguimiento, errors.New("el deportista_id es obligatorio")
	}

	if seguimiento.EntrenadorID <= 0 {
		return seguimiento, errors.New("el entrenador_id es obligatorio")
	}

	if seguimiento.Peso <= 0 {
		return seguimiento, errors.New("el peso debe ser mayor a 0")
	}

	if seguimiento.Altura <= 0 {
		return seguimiento, errors.New("la altura debe ser mayor a 0")
	}

	if seguimiento.Altura < 0.5 || seguimiento.Altura > 2.5 {
		return seguimiento, errors.New("la altura debe estar entre 0.5 y 2.5 metros")
	}

	seguimiento.ID = storage.SeguimientoIDCounter
	storage.SeguimientoIDCounter++

	seguimiento.IMC = calcularIMC(seguimiento.Peso, seguimiento.Altura)
	seguimiento.EstadoFisico = clasificarEstadoFisico(seguimiento.IMC)
	seguimiento.RequiereEvaluacionNutricional = requiereEvaluacionNutricional(seguimiento.EstadoFisico)

	if seguimiento.FechaRegistro == "" {
		seguimiento.FechaRegistro = time.Now().Format("2006-01-02")
	}

	seguimiento.CreatedAt = time.Now()
	seguimiento.UpdatedAt = time.Now()

	storage.Seguimientos = append(storage.Seguimientos, seguimiento)

	return seguimiento, nil
}

func ObtenerSeguimientos() []models.SeguimientoFisico {
	return storage.Seguimientos
}

func ObtenerSeguimientoPorID(id int) (models.SeguimientoFisico, bool) {
	for _, seguimiento := range storage.Seguimientos {
		if seguimiento.ID == id {
			return seguimiento, true
		}
	}

	return models.SeguimientoFisico{}, false
}

func ObtenerSeguimientosPorDeportista(deportistaID int) []models.SeguimientoFisico {
	var resultado []models.SeguimientoFisico

	for _, seguimiento := range storage.Seguimientos {
		if seguimiento.DeportistaID == deportistaID {
			resultado = append(resultado, seguimiento)
		}
	}

	return resultado
}

func ActualizarSeguimiento(id int, datos models.SeguimientoFisico) (models.SeguimientoFisico, error) {
	for i, seguimiento := range storage.Seguimientos {
		if seguimiento.ID == id {
			if datos.DeportistaID <= 0 {
				return seguimiento, errors.New("el deportista_id es obligatorio")
			}

			if datos.EntrenadorID <= 0 {
				return seguimiento, errors.New("el entrenador_id es obligatorio")
			}

			if datos.Peso <= 0 {
				return seguimiento, errors.New("el peso debe ser mayor a 0")
			}

			if datos.Altura <= 0 {
				return seguimiento, errors.New("la altura debe ser mayor a 0")
			}

			if datos.Altura < 0.5 || datos.Altura > 2.5 {
				return seguimiento, errors.New("la altura debe estar entre 0.5 y 2.5 metros")
			}

			datos.ID = id
			datos.IMC = calcularIMC(datos.Peso, datos.Altura)
			datos.EstadoFisico = clasificarEstadoFisico(datos.IMC)
			datos.RequiereEvaluacionNutricional = requiereEvaluacionNutricional(datos.EstadoFisico)
			datos.CreatedAt = seguimiento.CreatedAt
			datos.UpdatedAt = time.Now()

			if datos.FechaRegistro == "" {
				datos.FechaRegistro = seguimiento.FechaRegistro
			}

			storage.Seguimientos[i] = datos
			return datos, nil
		}
	}

	return models.SeguimientoFisico{}, errors.New("seguimiento no encontrado")
}

func EliminarSeguimiento(id int) bool {
	for i, seguimiento := range storage.Seguimientos {
		if seguimiento.ID == id {
			storage.Seguimientos = append(storage.Seguimientos[:i], storage.Seguimientos[i+1:]...)
			return true
		}
	}

	return false
}
