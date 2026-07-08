package services

import (
	"testing"

	"LigaFit-AWII2026/internal/models"
	"LigaFit-AWII2026/internal/storage"
)

func resetServiceTest() {
	storage.Seguimientos = nil
	storage.SeguimientoIDCounter = 1
	SetSeguimientoRepository(storage.NewSeguimientoMemoryRepository())
}

func TestCalcularIMC(t *testing.T) {
	tests := []struct {
		nombre string
		peso   float64
		altura float64
		want   float64
	}{
		{"normal", 70, 1.75, 22.86},
		{"bajo peso", 50, 1.75, 16.33},
		{"sobrepeso", 85, 1.75, 27.76},
		{"obesidad", 100, 1.60, 39.06},
	}
	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			got := calcularIMC(tt.peso, tt.altura)
			if got != tt.want {
				t.Errorf("calcularIMC(%v, %v) = %v; want %v", tt.peso, tt.altura, got, tt.want)
			}
		})
	}
}

func TestClasificarEstadoFisico(t *testing.T) {
	tests := []struct {
		nombre string
		imc    float64
		want   string
	}{
		{"bajo peso", 18.4, "Bajo peso"},
		{"normal minimo", 18.5, "Normal"},
		{"normal medio", 22.0, "Normal"},
		{"normal maximo", 24.9, "Normal"},
		{"sobrepeso minimo", 25.0, "Sobrepeso"},
		{"sobrepeso medio", 27.5, "Sobrepeso"},
		{"sobrepeso maximo", 29.9, "Sobrepeso"},
		{"obesidad minimo", 30.0, "Obesidad"},
	}
	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			got := clasificarEstadoFisico(tt.imc)
			if got != tt.want {
				t.Errorf("clasificarEstadoFisico(%v) = %v; want %v", tt.imc, got, tt.want)
			}
		})
	}
}

func TestRequiereEvaluacionNutricional(t *testing.T) {
	tests := []struct {
		estado string
		want   bool
	}{
		{"Bajo peso", true},
		{"Normal", false},
		{"Sobrepeso", false},
		{"Obesidad", true},
	}
	for _, tt := range tests {
		t.Run(tt.estado, func(t *testing.T) {
			got := requiereEvaluacionNutricional(tt.estado)
			if got != tt.want {
				t.Errorf("requiereEvaluacionNutricional(%v) = %v; want %v", tt.estado, got, tt.want)
			}
		})
	}
}

func TestCrearSeguimiento_DeportistaIDObligatorio(t *testing.T) {
	resetServiceTest()
	s := models.SeguimientoFisico{EntrenadorID: 1, Peso: 70, Altura: 1.75}
	_, err := CrearSeguimiento(s)
	if err == nil {
		t.Error("CrearSeguimiento() deberia fallar cuando DeportistaID es 0")
	}
}

func TestCrearSeguimiento_EntrenadorIDObligatorio(t *testing.T) {
	resetServiceTest()
	s := models.SeguimientoFisico{DeportistaID: 1, Peso: 70, Altura: 1.75}
	_, err := CrearSeguimiento(s)
	if err == nil {
		t.Error("CrearSeguimiento() deberia fallar cuando EntrenadorID es 0")
	}
}

func TestCrearSeguimiento_PesoMayorACero(t *testing.T) {
	resetServiceTest()
	s := models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 0, Altura: 1.75}
	_, err := CrearSeguimiento(s)
	if err == nil {
		t.Error("CrearSeguimiento() deberia fallar cuando Peso es 0")
	}
}

func TestCrearSeguimiento_AlturaMayorACero(t *testing.T) {
	resetServiceTest()
	s := models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 0}
	_, err := CrearSeguimiento(s)
	if err == nil {
		t.Error("CrearSeguimiento() deberia fallar cuando Altura es 0")
	}
}

func TestCrearSeguimiento_AlturaRangoInvalido(t *testing.T) {
	resetServiceTest()
	s := models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 3.0}
	_, err := CrearSeguimiento(s)
	if err == nil {
		t.Error("CrearSeguimiento() deberia fallar cuando Altura > 2.5")
	}
}

func TestCrearSeguimiento_Exitoso(t *testing.T) {
	resetServiceTest()
	s := models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75}
	result, err := CrearSeguimiento(s)
	if err != nil {
		t.Fatalf("CrearSeguimiento() error inesperado: %v", err)
	}
	if result.ID == 0 {
		t.Error("CrearSeguimiento() result.ID no deberia ser 0")
	}
	if result.IMC != 22.86 {
		t.Errorf("CrearSeguimiento() IMC = %v; want 22.86", result.IMC)
	}
	if result.EstadoFisico != "Normal" {
		t.Errorf("CrearSeguimiento() EstadoFisico = %v; want Normal", result.EstadoFisico)
	}
}

func TestCrearSeguimiento_RequiereEvaluacion(t *testing.T) {
	resetServiceTest()
	s := models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 100, Altura: 1.60}
	result, err := CrearSeguimiento(s)
	if err != nil {
		t.Fatalf("CrearSeguimiento() error inesperado: %v", err)
	}
	if !result.RequiereEvaluacionNutricional {
		t.Error("CrearSeguimiento() RequiereEvaluacionNutricional deberia ser true para obesidad")
	}
}

func TestObtenerSeguimientos_Vacio(t *testing.T) {
	resetServiceTest()
	got := ObtenerSeguimientos()
	if len(got) != 0 {
		t.Errorf("ObtenerSeguimientos() devolvio %d elementos, esperaba 0", len(got))
	}
}

func TestObtenerSeguimientos_ConDatos(t *testing.T) {
	resetServiceTest()
	CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})
	CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 2, EntrenadorID: 1, Peso: 80, Altura: 1.80})
	got := ObtenerSeguimientos()
	if len(got) != 2 {
		t.Errorf("ObtenerSeguimientos() devolvio %d elementos, esperaba 2", len(got))
	}
}

func TestObtenerSeguimientoPorID_Encontrado(t *testing.T) {
	resetServiceTest()
	creado, _ := CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})
	encontrado, ok := ObtenerSeguimientoPorID(creado.ID)
	if !ok {
		t.Fatal("ObtenerSeguimientoPorID() ok = false, esperaba true")
	}
	if encontrado.ID != creado.ID {
		t.Errorf("ObtenerSeguimientoPorID() ID = %d, esperaba %d", encontrado.ID, creado.ID)
	}
}

func TestObtenerSeguimientoPorID_NoEncontrado(t *testing.T) {
	resetServiceTest()
	_, ok := ObtenerSeguimientoPorID(999)
	if ok {
		t.Error("ObtenerSeguimientoPorID(999) ok = true, esperaba false")
	}
}

func TestObtenerSeguimientosPorDeportista_Filtra(t *testing.T) {
	resetServiceTest()
	CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})
	CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 2, EntrenadorID: 1, Peso: 80, Altura: 1.80})
	CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 2, Peso: 65, Altura: 1.65})
	got := ObtenerSeguimientosPorDeportista(1)
	if len(got) != 2 {
		t.Errorf("ObtenerSeguimientosPorDeportista(1) devolvio %d elementos, esperaba 2", len(got))
	}
}

func TestActualizarSeguimiento_Exitoso(t *testing.T) {
	resetServiceTest()
	creado, _ := CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})
	actualizado, err := ActualizarSeguimiento(creado.ID, models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 80, Altura: 1.75})
	if err != nil {
		t.Fatalf("ActualizarSeguimiento() error inesperado: %v", err)
	}
	if actualizado.Peso != 80 {
		t.Errorf("ActualizarSeguimiento() Peso = %v; want 80", actualizado.Peso)
	}
}

func TestActualizarSeguimiento_NoEncontrado(t *testing.T) {
	resetServiceTest()
	_, err := ActualizarSeguimiento(999, models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})
	if err == nil {
		t.Error("ActualizarSeguimiento(999) deberia fallar, seguimiento no encontrado")
	}
}

func TestEliminarSeguimiento_Exitoso(t *testing.T) {
	resetServiceTest()
	creado, _ := CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, EntrenadorID: 1, Peso: 70, Altura: 1.75})
	if !EliminarSeguimiento(creado.ID) {
		t.Fatal("EliminarSeguimiento() devolvio false, esperaba true")
	}
	if len(ObtenerSeguimientos()) != 0 {
		t.Error("despues de eliminar, la lista deberia estar vacia")
	}
}

func TestEliminarSeguimiento_NoEncontrado(t *testing.T) {
	resetServiceTest()
	if EliminarSeguimiento(999) {
		t.Error("EliminarSeguimiento(999) devolvio true, esperaba false")
	}
}
