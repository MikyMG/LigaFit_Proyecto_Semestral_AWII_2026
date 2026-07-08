package storage

import (
	"testing"

	"LigaFit-AWII2026/internal/models"
)

func setupMemoryTest() {
	Seguimientos = nil
	SeguimientoIDCounter = 1
}

func TestListarSeguimientos_Vacio(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	got := repo.ListarSeguimientos()
	if len(got) != 0 {
		t.Errorf("ListarSeguimientos() devolvio %d elementos, esperaba 0", len(got))
	}
}

func TestCrearSeguimiento_AsignaID(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	s := models.SeguimientoFisico{DeportistaID: 1}
	result := repo.CrearSeguimiento(s)
	if result.ID != 1 {
		t.Errorf("CrearSeguimiento() ID = %d, esperaba 1", result.ID)
	}
}

func TestCrearSeguimiento_IncrementaID(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	repo.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1})
	repo.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 2})
	if len(Seguimientos) != 2 {
		t.Errorf("Seguimientos tiene %d elementos, esperaba 2", len(Seguimientos))
	}
}

func TestListarSeguimientos_ConDatos(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	repo.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1})
	repo.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 2})
	got := repo.ListarSeguimientos()
	if len(got) != 2 {
		t.Errorf("ListarSeguimientos() devolvio %d elementos, esperaba 2", len(got))
	}
}

func TestBuscarSeguimientoPorID_Encontrado(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	creado := repo.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 10})
	encontrado, ok := repo.BuscarSeguimientoPorID(creado.ID)
	if !ok {
		t.Fatal("BuscarSeguimientoPorID() ok = false, esperaba true")
	}
	if encontrado.DeportistaID != 10 {
		t.Errorf("BuscarSeguimientoPorID() DeportistaID = %d, esperaba 10", encontrado.DeportistaID)
	}
}

func TestBuscarSeguimientoPorID_NoEncontrado(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	_, ok := repo.BuscarSeguimientoPorID(999)
	if ok {
		t.Error("BuscarSeguimientoPorID(999) ok = true, esperaba false")
	}
}

func TestActualizarSeguimiento_Existente(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	creado := repo.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1, Peso: 70})
	actualizado, ok := repo.ActualizarSeguimiento(creado.ID, models.SeguimientoFisico{DeportistaID: 1, Peso: 80})
	if !ok {
		t.Fatal("ActualizarSeguimiento() ok = false, esperaba true")
	}
	if actualizado.Peso != 80 {
		t.Errorf("ActualizarSeguimiento() Peso = %f, esperaba 80", actualizado.Peso)
	}
}

func TestActualizarSeguimiento_NoExistente(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	_, ok := repo.ActualizarSeguimiento(999, models.SeguimientoFisico{DeportistaID: 1})
	if ok {
		t.Error("ActualizarSeguimiento(999) ok = true, esperaba false")
	}
}

func TestBorrarSeguimiento_Existente(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	creado := repo.CrearSeguimiento(models.SeguimientoFisico{DeportistaID: 1})
	if !repo.BorrarSeguimiento(creado.ID) {
		t.Fatal("BorrarSeguimiento() devolvio false, esperaba true")
	}
	if len(Seguimientos) != 0 {
		t.Errorf("despues de borrar, Seguimientos tiene %d elementos, esperaba 0", len(Seguimientos))
	}
}

func TestBorrarSeguimiento_NoExistente(t *testing.T) {
	setupMemoryTest()
	repo := NewSeguimientoMemoryRepository()
	if repo.BorrarSeguimiento(999) {
		t.Error("BorrarSeguimiento(999) devolvio true, esperaba false")
	}
}
