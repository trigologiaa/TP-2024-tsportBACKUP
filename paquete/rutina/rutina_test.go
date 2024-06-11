package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAgregarRutina(t *testing.T) {
	gestorEj := ejercicio.NuevoGestorEjercicios()
	gestorRutinas := NuevoGestorRutinas(gestorEj)
	ejercicio1 := &ejercicio.Ejercicio{Nombre: "Ejercicio1", TiempoEnSegundos: 600, Calorias: 200, Tipo: "Cardio", Puntos: 10, Dificultad: "Media"}
	gestorEj.AgregarEjercicio(ejercicio1)
	rutina1 := &Rutina{Nombre: "Rutina1", CaracteristicasIndividuales: []ejercicio.Ejercicio{*ejercicio1}}
	err := gestorRutinas.AgregarRutina(rutina1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(gestorRutinas.ListarRutinas()))
}

func TestEliminarRutina(t *testing.T) {
	gestorEj := ejercicio.NuevoGestorEjercicios()
	gestorRutinas := NuevoGestorRutinas(gestorEj)
	ejercicio1 := &ejercicio.Ejercicio{Nombre: "Ejercicio1", TiempoEnSegundos: 600, Calorias: 200, Tipo: "Cardio", Puntos: 10, Dificultad: "Media"}
	gestorEj.AgregarEjercicio(ejercicio1)
	rutina1 := &Rutina{Nombre: "Rutina1", CaracteristicasIndividuales: []ejercicio.Ejercicio{*ejercicio1}}
	gestorRutinas.AgregarRutina(rutina1)
	err := gestorRutinas.EliminarRutina("Rutina1")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(gestorRutinas.ListarRutinas()))
}

func TestConsultarRutina(t *testing.T) {
	gestorEj := ejercicio.NuevoGestorEjercicios()
	gestorRutinas := NuevoGestorRutinas(gestorEj)
	ejercicio1 := &ejercicio.Ejercicio{Nombre: "Ejercicio1", TiempoEnSegundos: 600, Calorias: 200, Tipo: "Cardio", Puntos: 10, Dificultad: "Media"}
	gestorEj.AgregarEjercicio(ejercicio1)
	rutina1 := &Rutina{Nombre: "Rutina1", CaracteristicasIndividuales: []ejercicio.Ejercicio{*ejercicio1}}
	gestorRutinas.AgregarRutina(rutina1)
	rutinaConsultada, err := gestorRutinas.ConsultarRutina("Rutina1")
	assert.Nil(t, err)
	assert.Equal(t, "Rutina1", rutinaConsultada.Nombre)
}

func TestModificarRutina(t *testing.T) {
	gestorEj := ejercicio.NuevoGestorEjercicios()
	gestorRutinas := NuevoGestorRutinas(gestorEj)
	ejercicio1 := &ejercicio.Ejercicio{Nombre: "Ejercicio1", TiempoEnSegundos: 600, Calorias: 200, Tipo: "Cardio", Puntos: 10, Dificultad: "Media"}
	ejercicio2 := &ejercicio.Ejercicio{Nombre: "Ejercicio2", TiempoEnSegundos: 300, Calorias: 100, Tipo: "Fuerza", Puntos: 5, Dificultad: "Alta"}
	gestorEj.AgregarEjercicio(ejercicio1)
	gestorEj.AgregarEjercicio(ejercicio2)
	rutina1 := &Rutina{Nombre: "Rutina1", CaracteristicasIndividuales: []ejercicio.Ejercicio{*ejercicio1}}
	gestorRutinas.AgregarRutina(rutina1)
	nuevaRutina := &Rutina{Nombre: "Rutina1", CaracteristicasIndividuales: []ejercicio.Ejercicio{*ejercicio2}}
	err := gestorRutinas.ModificarRutina("Rutina1", nuevaRutina)
	assert.Nil(t, err)
	rutinaConsultada, _ := gestorRutinas.ConsultarRutina("Rutina1")
	assert.Equal(t, 1, len(rutinaConsultada.CaracteristicasIndividuales))
	assert.Equal(t, "Ejercicio2", rutinaConsultada.CaracteristicasIndividuales[0].Nombre)
}

func TestListarRutinas(t *testing.T) {
	gestorEj := ejercicio.NuevoGestorEjercicios()
	gestorRutinas := NuevoGestorRutinas(gestorEj)
	ejercicio1 := &ejercicio.Ejercicio{Nombre: "Ejercicio1", TiempoEnSegundos: 600, Calorias: 200, Tipo: "Cardio", Puntos: 10, Dificultad: "Media"}
	ejercicio2 := &ejercicio.Ejercicio{Nombre: "Ejercicio2", TiempoEnSegundos: 300, Calorias: 100, Tipo: "Fuerza", Puntos: 5, Dificultad: "Alta"}
	gestorEj.AgregarEjercicio(ejercicio1)
	gestorEj.AgregarEjercicio(ejercicio2)
	rutina1 := &Rutina{Nombre: "Rutina1", CaracteristicasIndividuales: []ejercicio.Ejercicio{*ejercicio1}}
	rutina2 := &Rutina{Nombre: "Rutina2", CaracteristicasIndividuales: []ejercicio.Ejercicio{*ejercicio2}}
	gestorRutinas.AgregarRutina(rutina1)
	gestorRutinas.AgregarRutina(rutina2)
	rutinas := gestorRutinas.ListarRutinas()
	assert.Equal(t, 2, len(rutinas))
	assert.Equal(t, "Rutina1", rutinas[0].Nombre)
	assert.Equal(t, "Rutina2", rutinas[1].Nombre)
}

func TestListarRutinasPorDificultad(t *testing.T) {
	gestorEj := ejercicio.NuevoGestorEjercicios()
	gestorRutinas := NuevoGestorRutinas(gestorEj)
	ejercicio1 := &ejercicio.Ejercicio{Nombre: "Ejercicio1", TiempoEnSegundos: 600, Calorias: 200, Tipo: "Cardio", Puntos: 10, Dificultad: "Media"}
	ejercicio2 := &ejercicio.Ejercicio{Nombre: "Ejercicio2", TiempoEnSegundos: 300, Calorias: 100, Tipo: "Fuerza", Puntos: 5, Dificultad: "Alta"}
	gestorEj.AgregarEjercicio(ejercicio1)
	gestorEj.AgregarEjercicio(ejercicio2)
	rutina1 := &Rutina{Nombre: "Rutina1", Dificultad: "Media", CaracteristicasIndividuales: []ejercicio.Ejercicio{*ejercicio1}}
	rutina2 := &Rutina{Nombre: "Rutina2", Dificultad: "Alta", CaracteristicasIndividuales: []ejercicio.Ejercicio{*ejercicio2}}
	gestorRutinas.AgregarRutina(rutina1)
	gestorRutinas.AgregarRutina(rutina2)
	rutinasMedia := gestorRutinas.ListarRutinasPorDificultad("Media")
	rutinasAlta := gestorRutinas.ListarRutinasPorDificultad("Alta")
	assert.Equal(t, 1, len(rutinasMedia))
	assert.Equal(t, "Rutina1", rutinasMedia[0].Nombre)
	assert.Equal(t, 1, len(rutinasAlta))
	assert.Equal(t, "Rutina2", rutinasAlta[0].Nombre)
}