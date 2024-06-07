package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerarRutinaAutomagica2(t *testing.T) {
	gestorRutinas := &GestorRutinas{}
	gestorEjercicios := ejercicio.NuevoGestorEjercicios()
	ejercicio1 := &ejercicio.Ejercicio{
		Nombre:     "Sentadillas",
		Tiempo:     60,
		Calorias:   100,
		Tipo:       "Fuerza",
		Puntos:     10,
		Dificultad: "Media",
	}
	ejercicio2 := &ejercicio.Ejercicio{
		Nombre:     "Flexiones",
		Tiempo:     45,
		Calorias:   80,
		Tipo:       "Fuerza",
		Puntos:     8,
		Dificultad: "Alta",
	}
	gestorEjercicios.AgregarEjercicio(ejercicio1)
	gestorEjercicios.AgregarEjercicio(ejercicio2)
	ejercicios := gestorEjercicios.ObtenerTodosLosEjercicios()
	nombre := "MiRutina"
	caloriasTotales := 150
	rutina, err := gestorRutinas.GenerarRutinaAutomagica2(ejercicios, nombre, caloriasTotales)
	assert.NoError(t, err, "Se esperaba nil, pero se obtuvo error")
	assert.NotNil(t, rutina, "La rutina no debería ser nil")
	assert.Equal(t, nombre, rutina.Nombre, "Nombre de la rutina incorrecto")
	assert.Equal(t, caloriasTotales, rutina.Calorias, "Calorías totales de la rutina incorrectas")
}
